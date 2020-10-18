package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_User(t *testing.T) {
	var err error
	uu := make([]User, 2)
	uu[0], err = NewUser(
		`https://images.pexels.com/photos/2745151/pexels-photo-2745151.jpeg?auto=compress&cs=tinysrgb&dpr=1&w=500`,
		`Antonio`,
		`antonio@test.it`,
		`Amo i cani`,
		`pippo`,
	)
	assert.Nil(t, err)

	uu[1], err = NewUser(
		`https://images.pexels.com/photos/3294248/pexels-photo-3294248.jpeg?auto=compress&cs=tinysrgb&h=750&w=1260`,
		`Anna`,
		`anna@casa.it`,
		`Amo gli cani tutti`,
		`pluto`,
	)
	assert.Nil(t, err)
	assert.NotEqual(t, "", uu[0].ID)
	assert.NotEqual(t, "", uu[1].ID)
	assert.NotEqual(t, uu[0].ID, uu[1].ID)
	//t.Logf("user 0 ID: %s", uu[0].ID)

	validated := uu[0].ValidatePassword("pippo")
	assert.True(t, validated)

	validated = uu[0].ValidatePassword("wrong pass")
	assert.False(t, validated)

}
