package model

import (
	"fmt"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/stretchr/testify/assert"
)

func Test_Location(t *testing.T) {
	connString := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s",
		"zampapp",
		"zampapp",
		"127.0.0.1",
		"zampapp",
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
