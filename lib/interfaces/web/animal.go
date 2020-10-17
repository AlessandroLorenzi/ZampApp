package web

import (
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
		s.webReturn(w, 400, "animal code is not valid")
		return
	}

	a, err := s.repoService.GetAnimal(idAnimal)
	if err != nil {
		if err.Error() == "not found" {
			s.logger.Warningf("animal code is not valid")
			s.webReturn(w, 400, "animal code is not valid")
			return
		}
		s.logger.WithField("error", err).Errorf("Unexpected error")
		s.webReturn(w, 500, "unexpected error")
		return
	}

	s.webReturn(w, 200, "ok", responseContent{"animal": a})
}

func (s Service) GetAnimals(w http.ResponseWriter, r *http.Request) {
	aa, err := s.repoService.GetAnimals()
	if err != nil {
		s.logger.WithField("error", err).Errorf("Unexpected error")
		s.webReturn(w, 500, "unexpected error")

		return
	}
	s.webReturn(w, 200, "ok", responseContent{"animals": aa})
}
