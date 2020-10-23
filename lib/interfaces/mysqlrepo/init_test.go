package mysqlrepo

import (
	"fmt"
	"os"
	"testing"
	"zampapp/lib/entity/model"

	"github.com/sirupsen/logrus"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func generateTestService(t *testing.T) Service {
	connString := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_DB"),
	)
	gormDB, err := gorm.Open(mysql.Open(connString), &gorm.Config{})
	if err != nil {
		t.Fatal("Impossible to open the connection!")
	}

	err = gormDB.AutoMigrate(&model.User{}, &model.Animal{})
	if err != nil {
		t.Fatal("Error during  migration")
	}

	err = gormDB.AutoMigrate(&model.Animal{})
	if err != nil {
		t.Fatal("Error animal migration")
	}

	repoService := New(
		gormDB,
		logrus.NewEntry(logrus.New()),
	)

	return repoService
}
