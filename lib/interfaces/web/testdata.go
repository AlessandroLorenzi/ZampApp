package web

import (
	"net/http"
	"zampapp/lib/entity/model"
)

// TODO REMOVE
func (s Service) TestData(w http.ResponseWriter, r *http.Request) {
	uu := make([]model.User, 5)

	uu[0] = model.NewUser(
		`https://images.pexels.com/photos/2745151/pexels-photo-2745151.jpeg?auto=compress&cs=tinysrgb&dpr=1&w=500`,
		`Antonio`,
		`antonio@test.it`,
		`Amo i cani`,
		`pippo`,
	)
	uu[1] = model.NewUser(
		`https://images.pexels.com/photos/3294248/pexels-photo-3294248.jpeg?auto=compress&cs=tinysrgb&h=750&w=1260`,
		`Anna`,
		`anna@casa.it`,
		`Amo gli cani tutti`,
		`pluto`,
	)
	uu[2] = model.NewUser(
		`https://images.pexels.com/photos/2666154/pexels-photo-2666154.jpeg?auto=compress&cs=tinysrgb&h=750&w=1260`,
		`Giovanna`,
		`giovanna@gmail.com`,
		`Ho un bastardino che mi piace tanto`,
		`paperino`,
	)
	uu[3] = model.NewUser(
		`https://images.pexels.com/photos/4064423/pexels-photo-4064423.jpeg?auto=compress&cs=tinysrgb&h=750&w=1260`,
		`Silvana`,
		`silvana@coseacaso.it`,
		`Sono una gattara`,
		`paperino`,
	)
	uu[4] = model.NewUser(
		`https://images.pexels.com/photos/3889895/pexels-photo-3889895.jpeg?auto=compress&cs=tinysrgb&h=750&w=1260`,
		`Eliana`,
		`eliana@blabla.it`,
		`Viaggio spesso e vado in giro`,
		`paperino`,
	)

	aa := []model.Animal{
		model.Animal{
			ID:            "1",
			Name:          "Fufi",
			Breed:         "Terrier",
			OwnerID:       1,
			Picture:       "https://images.pexels.com/photos/733416/pexels-photo-733416.jpeg?auto=compress&cs=tinysrgb&h=750&w=1260",
			Wormed:        true,
			ChildFriendly: true,
			Sex:           true,
			Position: model.Location{
				X: 1.02332049,
				Y: 2.32490,
			},
			PositionDesc: "Allevamento tal de tali",
			Description:  "Cane terrier molto simpatico",
		},
		model.Animal{
			ID:            "2",
			Name:          "Fido",
			Breed:         "Dalmata",
			OwnerID:       1,
			Picture:       "https://images.pexels.com/photos/933498/pexels-photo-933498.jpeg?auto=compress&cs=tinysrgb&dpr=1&w=500",
			Wormed:        true,
			ChildFriendly: true,
			Sex:           true,
			Position: model.Location{
				X: 1.02332049,
				Y: 2.32490,
			},
			PositionDesc: "Allevamento tal de tali",
			Description:  "Dalmata super ammaestrato. bravo",
		},
		model.Animal{
			ID:            "3",
			Name:          "Witch",
			Breed:         "Boxer",
			OwnerID:       2,
			Picture:       "https://images.pexels.com/photos/5422769/pexels-photo-5422769.jpeg?auto=compress&cs=tinysrgb&dpr=1&w=500",
			Wormed:        true,
			ChildFriendly: true,
			Sex:           false,
			Position: model.Location{
				X: 1.01231239,
				Y: 2.335490,
			},
			PositionDesc: "Allevamento talaltro",
			Description:  "Cagnetta molto simpatica a cui piacciono i cappelli",
		},
	}

	ret := s.gormDB.Create(uu)
	if ret.Error != nil {
		s.logger.Errorf("Errore utenti %v", ret.Error)
	}
	ret = s.gormDB.Create(aa)
	if ret.Error != nil {
		s.logger.Errorf("Errore animali %v", ret.Error)
	}
}
