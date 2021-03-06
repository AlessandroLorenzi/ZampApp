package web

import (
	"zampapp/lib/entity/model"

	"github.com/gin-gonic/gin"
)

// TODO REMOVE
func (s Service) TestData(_ *gin.Context) {
	uu := make([]model.User, 6)
	uu[0], _ = model.NewUser(
		`https://images.pexels.com/photos/2745151/pexels-photo-2745151.jpeg?auto=compress&cs=tinysrgb&dpr=1&w=500`,
		`antonio@test.it`,
		`antonio`,
		`Amo i cani`,
		`pippo`,
	)
	uu[1], _ = model.NewUser(
		`https://images.pexels.com/photos/3294248/pexels-photo-3294248.jpeg?auto=compress&cs=tinysrgb&h=750&w=1260`,
		`anna@casa.it`,
		`Anna`,
		`Amo gli cani tutti`,
		`pluto`,
	)
	uu[2], _ = model.NewUser(
		`https://images.pexels.com/photos/2666154/pexels-photo-2666154.jpeg?auto=compress&cs=tinysrgb&h=750&w=1260`,
		`giovanna@gmail.com`,
		`Giovanna`,
		`Ho un bastardino che mi piace tanto`,
		`paperino`,
	)
	uu[3], _ = model.NewUser(
		`https://images.pexels.com/photos/4064423/pexels-photo-4064423.jpeg?auto=compress&cs=tinysrgb&h=750&w=1260`,
		`silvana@coseacaso.it`,
		`Silvana`,
		`Mi piacciono i gatti`,
		`paperino`,
	)
	uu[4], _ = model.NewUser(
		`https://images.pexels.com/photos/3889895/pexels-photo-3889895.jpeg?auto=compress&cs=tinysrgb&h=750&w=1260`,
		`eliana@blabla.it`,
		`Eliana`,
		`Viaggio spesso e vado in giro`,
		`paperino`,
	)
	uu[5], _ = model.NewUser(
		`https://avatars0.githubusercontent.com/u/150980?s=460&u=4e1a54c01546f218f8911185cd565d33fe6c571e&v=4`,
		`alorenzi@alorenzi.eu`,
		`alorenzi`,
		`test`,
		`test`,
	)

	_ = s.repoService.CreateUser(uu[0])
	_ = s.repoService.CreateUser(uu[1])
	_ = s.repoService.CreateUser(uu[2])
	_ = s.repoService.CreateUser(uu[3])
	_ = s.repoService.CreateUser(uu[4])
	_ = s.repoService.CreateUser(uu[5])

	aa := make([]model.Animal, 4)
	aa[0], _ = model.NewAnimal(
		"Fufi",
		"Terrier",
		1,
		true,
		uu[0].ID,
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

	aa[1], _ = model.NewAnimal(
		"Fufi",
		"Terrier",
		1,
		true,
		uu[0].ID,
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
	aa[2], _ = model.NewAnimal(
		"Fido",
		"Dalmata",
		1,
		false,
		uu[1].ID,
		"https://images.pexels.com/photos/933498/pexels-photo-933498.jpeg?auto=compress&cs=tinysrgb&dpr=1&w=500",
		true,
		true,
		model.Location{
			X: 1.02332049,
			Y: 2.32490,
		},
		"Allevamento tal de tali",
		"Dalmata super ammaestrato. bravo",
	)
	aa[3], _ = model.NewAnimal(
		"Witch",
		"Boxer",
		2,
		false,
		uu[1].ID,
		"https://images.pexels.com/photos/5422769/pexels-photo-5422769.jpeg?auto=compress&cs=tinysrgb&dpr=1&w=500",
		true,
		true,
		model.Location{
			X: 1.01231239,
			Y: 2.335490,
		},
		"Allevamento talaltro",
		"Cagnetta molto simpatica a cui piacciono i cappelli",
	)

	_ = s.repoService.CreateAnimal(aa[0])
	_ = s.repoService.CreateAnimal(aa[1])
	_ = s.repoService.CreateAnimal(aa[2])
	_ = s.repoService.CreateAnimal(aa[3])
}
