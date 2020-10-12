package web

import (
	"net/http"
	"zampapp/lib/entity/model"
)

// TODO REMOVE
func (s Service) TestData(w http.ResponseWriter, r *http.Request) {
	uu := []model.User{
		model.User{
			ID:          1,
			Picture:     `https://images.pexels.com/photos/2745151/pexels-photo-2745151.jpeg?auto=compress&cs=tinysrgb&dpr=1&w=500`,
			NickName:    `Antonio`,
			Email:       `antonio@test.it`,
			Description: `Amo i cani`,
		},
		model.User{
			ID:          2,
			Picture:     `https://images.pexels.com/photos/3294248/pexels-photo-3294248.jpeg?auto=compress&cs=tinysrgb&h=750&w=1260`,
			NickName:    `Anna`,
			Email:       `anna@casa.it`,
			Description: `Amo gli cani tutti`,
		},
		model.User{
			ID:          3,
			Picture:     `https://images.pexels.com/photos/2666154/pexels-photo-2666154.jpeg?auto=compress&cs=tinysrgb&h=750&w=1260`,
			NickName:    `Giovanna`,
			Email:       `giovanna@gmail.com`,
			Description: `Ho un bastardino che mi piace tanto`,
		},
		model.User{
			ID:          4,
			Picture:     `https://images.pexels.com/photos/4064423/pexels-photo-4064423.jpeg?auto=compress&cs=tinysrgb&h=750&w=1260`,
			NickName:    `Silvana`,
			Email:       `silvana@coseacaso.it`,
			Description: `Sono una gattara`,
		},
		model.User{
			ID:          5,
			Picture:     `https://images.pexels.com/photos/3889895/pexels-photo-3889895.jpeg?auto=compress&cs=tinysrgb&h=750&w=1260`,
			NickName:    `Eliana`,
			Email:       `eliana@blabla.it`,
			Description: `Viaggio spesso e vado in giro`,
		},
	}

	aa := []model.Animal{
		model.Animal{
			ID:            1,
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
			ID:            2,
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
			ID:            3,
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
