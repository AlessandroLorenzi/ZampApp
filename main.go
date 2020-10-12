package main

import (
	"fmt"
	"log"
	"zampapp/lib/entity/model"
	"zampapp/lib/interfaces/mysqlrepo"
	"zampapp/lib/interfaces/web"

	"github.com/sirupsen/logrus"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	loggerEntity := logrus.NewEntry(logger)

	connString := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s",
		"zampapp",
		"zampapp",
		"127.0.0.1",
		"zampapp",
	)
	gormDB, err := gorm.Open(mysql.Open(connString), &gorm.Config{})
	if err != nil {
		loggerEntity.WithField("err", err).Fatal("Impossible to open the connection!")
	}

	err = gormDB.AutoMigrate(&model.User{}, &model.Animal{})
	if err != nil {
		loggerEntity.WithField("err", err).Fatal("Error during  migration")
	}

	err = gormDB.AutoMigrate(&model.Animal{})
	if err != nil {
		loggerEntity.WithField("err", err).Fatal("Error animal migration")
	}

	repoService := mysqlrepo.New(
		gormDB,
		loggerEntity,
	)

	webservice := web.New(
		loggerEntity,
		gormDB,
		repoService,
	)

	loggerEntity.Info("Here we GO!")

	log.Fatal(webservice.ListenAndServe())

}
