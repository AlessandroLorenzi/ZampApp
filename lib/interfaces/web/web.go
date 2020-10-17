package web

import (
	"encoding/json"
	"net/http"
	"time"
	"zampapp/lib/entity/model"

	"gorm.io/gorm"

	"github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

type Service struct {
	repoService     repo
	gormDB          *gorm.DB // mi serve per test data
	server          *http.Server
	logger          *logrus.Entry
	usecasesService useCasesResolver
}

type repo interface {
	GetAnimal(idAnimal int) (model.Animal, error)
	GetAnimals() ([]model.Animal, error)
	GetUser(idUser int) (model.User, error)
}

type useCasesResolver interface {
	Login(login, password string) (model.User, error)
}

func New(
	logger *logrus.Entry,
	gormDB *gorm.DB,
	repoService repo,
	usecasesService useCasesResolver,

) Service {
	s := Service{
		gormDB:          gormDB,
		logger:          logger,
		repoService:     repoService,
		usecasesService: usecasesService,
	}

	router := mux.NewRouter()
	router.Use(s.loggingMiddleware)

	router.HandleFunc("/api/health", s.healthCheck).Methods("GET")
	router.HandleFunc("/api/login", s.login).Methods("POST")

	router.HandleFunc("/api/animal/{id_animal}", s.GetAnimal).Methods("GET")
	router.HandleFunc("/api/animals", s.GetAnimals).Methods("GET")

	router.HandleFunc("/api/user/{id_user}", s.GetUser).Methods("GET")

	router.HandleFunc("/api/testdata", s.TestData) // TODO REMOVE

	s.server = &http.Server{
		Handler: router,
		Addr:    "0.0.0.0:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	return s
}

func (s Service) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		s.logger = s.logger.WithField("requestId", 123) // TODO unique request id
		executionStart := time.Now()

		// Run
		next.ServeHTTP(w, r)

		s.logger.WithFields(logrus.Fields{
			"time_spent":  time.Since(executionStart),
			"request_url": r.RequestURI,
			"method":      r.Method,
		}).Printf("")
	})
}

func (s Service) ListenAndServe() error {
	return s.server.ListenAndServe()
}

type responseContent map[string]interface{}

func (s *Service) webReturn(w http.ResponseWriter, statusCode int, msg string, rcs ...responseContent) {

	rc := make(responseContent)
	for _, tmpRc := range rcs {
		for k, v := range tmpRc {
			rc[k] = v
		}
	}

	rc["msg"] = msg

	logger := s.logger.WithFields(logrus.Fields{
		"code": statusCode,
	})
	if statusCode < 500 {
		logger.Info(msg)
	} else {
		logger.Error(msg)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(rc)
}
