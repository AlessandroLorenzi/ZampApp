package usecases

import (
	"fmt"
	"testing"
	"zampapp/lib/entity/model"
	"zampapp/lib/interfaces/mysqlrepo"

	"github.com/stretchr/testify/assert"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Test_Login(t *testing.T) {
	connString := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s",
		"zampapp",
		"zampapp",
		"127.0.0.1",
		"zampapp",
	)
	gormDB, _ := gorm.Open(mysql.Open(connString), &gorm.Config{})

	repoService := mysqlrepo.New(
		gormDB,
		logrus.NewEntry(logrus.New()),
	)
	u, _ := model.NewUser(
		`https://images.pexels.com/photos/2745151/pexels-photo-2745151.jpeg?auto=compress&cs=tinysrgb&dpr=1&w=500`,
		`antonio@test.it`,
		`antonio`,
		`Amo i cani`,
		`pippo`,
	)
	_ = repoService.CreateUser(u)

	useCasesService := New(repoService)
	theUser, err := useCasesService.Login(`antonio`, `pippo`)

	assert.Nil(t, err)
	assert.Equal(t, u.NickName, theUser.NickName)
	assert.Equal(t, u.Email, theUser.Email)

	theUser2, err := useCasesService.Login(`antonio@test.it`, `pippo`)

	assert.Nil(t, err)
	assert.Equal(t, u.NickName, theUser2.NickName)
	assert.Equal(t, u.Email, theUser2.Email)

	_, err = useCasesService.Login(`antonio@test.it`, `errore`)
	assert.NotNil(t, err)
	assert.Equal(t, `user and password don't match`, err.Error())

	_, err = useCasesService.Login(`errore`, `errore`)
	assert.NotNil(t, err)
	assert.Equal(t, `record not found`, err.Error())

	_ = repoService.DeleteUser(u.ID)
}
