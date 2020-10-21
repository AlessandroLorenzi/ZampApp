package mysqlrepo

import (
	"testing"
	"zampapp/lib/entity/model"

	"github.com/stretchr/testify/assert"
)

func Test_User(t *testing.T) {
	s := generateTestService(t)

	u, _ := model.NewUser(
		`https://images.pexels.com/photos/2745151/pexels-photo-2745151.jpeg?auto=compress&cs=tinysrgb&dpr=1&w=500`,
		`antonio@test.it`,
		`antonio`,
		`Amo i cani`,
		`pippo`,
	)
	err := s.CreateUser(u)
	assert.Nil(t, err)

	u.Description = "test update"
	err = s.UpdateUser(u)
	assert.Nil(t, err)

	u, err = s.GetUser(u.ID)

	assert.Nil(t, err)
	assert.Equal(t, "test update", u.Description)

	err = s.DeleteUser(u.ID)

	assert.Nil(t, err)

	u, err = s.GetUser(u.ID)
	assert.NotNil(t, err)
	assert.Equal(t, "not found", err.Error())

}
