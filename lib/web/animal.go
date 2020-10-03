package web

import (
	"encoding/json"
	"net/http"
	"strconv"
	"zampapp/lib/entity/model"

	"github.com/gorilla/mux"
)

func (s Service) GetAnimal(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idAnimalStr := vars["id_animal"]

	idAnimal, err := strconv.Atoi(idAnimalStr)
	if err != nil {
		s.logger.Warningf("animal code is not valid")
		w.WriteHeader(400)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{"msg": "animal code is not valid"})
		return
	}

	u := model.Animal{}
	// Search animal
	res := s.gormDB.Find(&u, idAnimal)
	err = res.Error
	if err != nil {
		s.logger.WithField("error", err).Errorf("Unexpected error")
		w.WriteHeader(500)
		_ = json.NewEncoder(w).Encode(
			map[string]interface{}{"msg": "Unexpected error"},
		)
		return
	}

	if res.RowsAffected != 1 {
		s.logger.Warningf("animal code is not valid")
		w.WriteHeader(400)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{"msg": "animal code is not valid"})
		return
	}

	_ = json.NewEncoder(w).Encode(u)
}

func (s Service) GetAnimals(w http.ResponseWriter, r *http.Request) {
	var (
		aa []model.Animal
	)

	res := s.gormDB.Find(&aa)
	if res.Error != nil {
		s.logger.WithField("error", res.Error).Errorf("Error retriving animals")
		w.WriteHeader(500)
		_ = json.NewEncoder(w).Encode(
			map[string]interface{}{"msg": "Error retriving animals", "err": res.Error},
		)
		return
	}

	for i := range aa {
		res = s.gormDB.Find(&aa[i].Owner, aa[i].OwnerID)
		if res.Error != nil {
			s.logger.WithField("error", res.Error).Errorf("Error retriving animals")
			w.WriteHeader(500)
			_ = json.NewEncoder(w).Encode(
				map[string]interface{}{"msg": "Error retriving user", "err": res.Error},
			)
			return
		}
	}

	_ = json.NewEncoder(w).Encode(aa)
}
