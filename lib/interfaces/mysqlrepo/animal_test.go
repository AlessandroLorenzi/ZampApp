package mysqlrepo

import (
	"testing"
	"zampapp/lib/entity/model"

	"github.com/stretchr/testify/assert"
)

func Test_Animal(t *testing.T) {
	s := generateTestService(t)

	u, _ := model.NewUser(
		`https://images.pexels.com/photos/2745151/pexels-photo-2745151.jpeg?auto=compress&cs=tinysrgb&dpr=1&w=500`,
		`Antonio`,
		`antonio@test.it`,
		`Amo i cani`,
		`pippo`,
	)
	_ = s.CreateUser(u)

	a, _ := model.NewAnimal(
		"Fufi",
		"Terrier",
		1,
		true,
		u.ID,
		"https://images.pexels.com/photos/733416/pexels-photo-733416.jpeg?auto=compress&cs=tinysrgb&h=750&w=1260",
		true,
		true,
		model.Location{
			X: 1.02332049,
			Y: 2.32490,
		},
		"Allevamento tal de tali",
		"Cane terrier molto simpatico",
	)

	var err error

	err = s.CreateAnimal(a)
	assert.Nil(t, err)

	err = s.CreateAnimal(a)
	assert.Contains(t, err.Error(), "Duplicate entry")

	err = s.DeleteAnimal(a.ID)
	assert.Nil(t, err)
}
