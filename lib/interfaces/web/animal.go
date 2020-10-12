package web

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (s Service) GetAnimal(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idAnimalStr := vars["id_animal"]

	idAnimal, err := strconv.Atoi(idAnimalStr)
	if err != nil {
		s.logger.Warningf("animal code is not valid")
		w.WriteHeader(400)
		_ = json.NewEncoder(w).Encode(
			map[string]interface{}{
				"msg": "animal code is not valid",
			},
		)
		return
	}

	a, err := s.repoService.GetAnimal(idAnimal)
	if err != nil {
		if err.Error() == "not found" {
			s.logger.Warningf("animal code is not valid")
			w.WriteHeader(400)
			_ = json.NewEncoder(w).Encode(map[string]interface{}{
				"msg": "animal code is not valid",
			})
			return
		}
		s.logger.WithField("error", err).Errorf("Unexpected error")
		w.WriteHeader(500)
		_ = json.NewEncoder(w).Encode(
			map[string]interface{}{
				"msg": "unexpected error",
			},
		)
		return
	}

	_ = json.NewEncoder(w).Encode(a)
}

func (s Service) GetAnimals(w http.ResponseWriter, r *http.Request) {
	aa, err := s.repoService.GetAnimals()
	if err != nil {
		s.logger.WithField("error", err).Errorf("Unexpected error")
		w.WriteHeader(500)
		_ = json.NewEncoder(w).Encode(
			map[string]interface{}{
				"msg": "Unexpected error",
			},
		)
		return
	}

	_ = json.NewEncoder(w).Encode(aa)
}
