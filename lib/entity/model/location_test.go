package model

import (
	"fmt"
	"os"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/stretchr/testify/assert"
)

func Test_Location(t *testing.T) {
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
	var a Animal

	tx := gormDB.First(&a)

	assert.Nil(t, tx.Error)
	assert.NotZero(t, a.Position.X)
	assert.NotZero(t, a.Position.Y)
}
