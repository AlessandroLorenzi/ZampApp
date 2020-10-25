package web

import (
	"log"
	"zampapp/lib/interfaces/mysqlrepo"
	"zampapp/lib/usecases"

	jwt "github.com/appleboy/gin-jwt/v2"
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

	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/api/health", s.healthCheck)

	r.GET("/api/animal/:id_animal", s.getAnimal)
	r.POST("/api/animal", s.newAnimal)
	r.GET("/api/animals", s.getAnimals)

	r.GET("/api/user_loves/:id_animal", s.newLove)
	r.GET("/api/user_loves/", s.listAnimalsLoved)

	r.GET("/api/user/:id_user", s.GetUser)

	r.GET("/api/testdata", s.TestData) // TODO REMOVE

	authMiddleware, _ := jwt.New(s.jwtMiddleware())
	errInit := authMiddleware.MiddlewareInit()
	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

	r.POST("/api/login", authMiddleware.LoginHandler)

	s.server = r
	return s
}

func (s Service) ListenAndServe() error {
	return s.server.Run()
}
