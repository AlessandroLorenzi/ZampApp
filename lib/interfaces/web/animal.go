package web

import (
	"encoding/json"
	"net/http"
	"strconv"
	"zampapp/lib/entity/model"

	"github.com/gorilla/mux"
)

func (s Service) getAnimal(w http.ResponseWriter, r *http.Request) {
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

func (s Service) getAnimals(w http.ResponseWriter, r *http.Request) {
	aa, err := s.repoService.GetAnimals()
	if err != nil {
		s.logger.WithField("error", err).Errorf("Unexpected error")
		s.webReturn(w, 500, "unexpected error")

		return
	}
	s.webReturn(w, 200, "ok", responseContent{"animals": aa})
}

func (s Service) newAnimal(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		s.webReturn(w, 400, "impossible to get parse form")
	}
	if len(r.Form) != 1 {
		s.webReturn(w, 400, "empty document")
		return
	}

	type newAnimalPost struct {
		Name          string         `json:"name"`
		Breed         string         `json:"breed"`
		Size          int            `json:"size"`
		Sex           bool           `json:"sex"`
		OwnerID       int            `json:"owner_id"`
		Picture       string         `json:"picture"`
		Wormed        bool           `json:"wormed"`
		ChildFriendly bool           `json:"child_friendly"`
		Position      model.Location `json:"position"`
		PositionDesc  string         `json:"position_desc"`
		Description   string         `json:"description"`
	}

	var ap newAnimalPost

	for key := range r.Form {
		err := json.Unmarshal([]byte(key), &ap)
		if err != nil {
			s.webReturn(w, 400, "document not valid")
			return
		}
		continue
	}

	a, err := model.NewAnimal(
		ap.Name,
		ap.Breed,
		ap.Size,
		ap.Sex,
		ap.OwnerID,
		ap.Picture,
		ap.Wormed,
		ap.ChildFriendly,
		ap.Position,
		ap.PositionDesc,
		ap.Description,
	)
	if err != nil {
		s.webReturn(w, 400, "document error", responseContent{"error": err.Error()})
		return
	}

	s.webReturn(w, 200, "token generated",
		responseContent{
			"animal": a,
		},
	)

}
