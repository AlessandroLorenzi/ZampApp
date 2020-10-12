package mysqlrepo

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Service struct {
	gormDB *gorm.DB
	logger *logrus.Entry
}

func New(
	gormDB *gorm.DB,
	logger *logrus.Entry,
) Service {
	return Service{
		gormDB,
		logger,
	}
}
