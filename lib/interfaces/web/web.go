package web

import (
	"zampapp/lib/interfaces/mysqlrepo"
	"zampapp/lib/usecases"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Service struct {
	repoService     mysqlrepo.Service
	server          *gin.Engine
	logger          *logrus.Entry
	useCasesService usecases.Service
}

func New(
	logger *logrus.Entry,
	repoService mysqlrepo.Service,
	useCasesService usecases.Service,
) Service {
	s := Service{
		logger:          logger,
		repoService:     repoService,
		useCasesService: useCasesService,
	}

	router := gin.Default()

	router.GET("/api/health", s.healthCheck)
	router.POST("/api/login", s.login)

	router.GET("/api/animal/:id_animal", s.getAnimal)
	router.POST("/api/animal", s.newAnimal)
	router.GET("/api/animals", s.getAnimals)

	router.GET("/api/user/:id_user", s.GetUser)

	router.GET("/api/testdata", s.TestData) // TODO REMOVE

	s.server = router
	return s
}

func (s Service) ListenAndServe() error {
	return s.server.Run()
}
