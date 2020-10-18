package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Animal(t *testing.T) {
	a, err := NewAnimal(
		"Fufi",
		"Terrier",
		1,
		true,
		"test",
		"https://images.pexels.com/photos/733416/pexels-photo-733416.jpeg?auto=compress&cs=tinysrgb&h=750&w=1260",
		true,
		true,
		Location{
			X: 1.02332049,
			Y: 2.32490,
		},
		"Allevamento tal de tali",
		"Cane terrier molto simpatico",
	)
	assert.Nil(t, err)
	assert.NotEqual(t, "", a.ID)

	a2, _ := NewAnimal(
		"Fufi2",
		"Terrier",
		1,
		true,
		"test",
		"https://images.pexels.com/photos/733416/pexels-photo-733416.jpeg?auto=compress&cs=tinysrgb&h=750&w=1260",
		true,
		true,
		Location{
			X: 1.02332049,
			Y: 2.32490,
		},
		"Allevamento tal de tali",
		"Cane terrier molto simpatico",
	)

	// Two animals, two IDs
	assert.NotEqual(t, a2.ID, a.ID)

}
