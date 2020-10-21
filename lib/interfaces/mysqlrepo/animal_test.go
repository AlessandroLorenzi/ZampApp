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
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "Duplicate entry")

	// Test GetAnimals
	aa, err := s.GetAnimals()

	assert.Nil(t, err)
	assert.NotEmpty(t, aa)

	// Test UpdateAnimal
	a.Description = "updated description"
	err = s.UpdateAnimal(a)
	assert.Nil(t, err)

	// Test GetAnimal
	a1, err := s.GetAnimal(a.ID)

	assert.Nil(t, err)
	assert.Equal(t, a.ID, a1.ID)
	assert.Equal(t, "updated description", a1.Description)

	// Delete test animal & user
	err = s.DeleteAnimal(a.ID)
	assert.Nil(t, err)

	err = s.DeleteUser(u.ID)
	assert.Nil(t, err)
}

func TestService_GetAnimal_NotFound(t *testing.T) {
	s := generateTestService(t)

	_, err := s.GetAnimal("id-false")

	assert.NotNil(t, err)
	assert.Equal(t, "not found", err.Error())

}
