package web

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Service) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idUser := vars["id_user"]

	u, err := s.repoService.GetUser(idUser)
	if err != nil {
		if err.Error() == "not found" {
			s.webReturn(w, 400, "user does not exists")
			return
		}
		s.logger.WithField("error", err).Errorf("Unexpected error")
		s.webReturn(w, 500, "unexpected error")

		return
	}

	_ = json.NewEncoder(w).Encode(u)
}
