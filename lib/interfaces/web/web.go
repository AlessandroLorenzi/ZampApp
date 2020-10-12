package web

import (
	"net/http"
	"time"
	"zampapp/lib/entity/model"

	"gorm.io/gorm"

	"github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

type Service struct {
	repoService repo
	gormDB      *gorm.DB // mi serve per test data
	server      *http.Server
	logger      *logrus.Entry
}

type repo interface {
	GetAnimal(idAnimal int) (model.Animal, error)
	GetAnimals() ([]model.Animal, error)
	GetUser(idUser int) (model.User, error)
}

func New(
	logger *logrus.Entry,
	gormDB *gorm.DB,
	repoService repo,
) Service {
	s := Service{
		gormDB:      gormDB,
		logger:      logger,
		repoService: repoService,
	}

	router := mux.NewRouter()
	router.Use(s.loggingMiddleware)

	router.HandleFunc("/api/health", s.healthCheck)

	router.HandleFunc("/api/animal/{id_animal}", s.GetAnimal)
	router.HandleFunc("/api/animals", s.GetAnimals)

	router.HandleFunc("/api/user/{id_user}", s.GetUser)

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
		w.Header().Add("Content-Type", "application/json")

		s.logger = s.logger.WithField("requestId", 123)
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
