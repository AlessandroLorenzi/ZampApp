package web

import (
	"encoding/json"
	"net/http"
	"strconv"
	"zampapp/lib/entity/model"

	"github.com/gorilla/mux"
)

func (s Service) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idUserStr := vars["id_user"]

	// Sanificazione input
	if idUserStr == "" {
		s.logger.Warningf("Valore non valido")
		w.WriteHeader(400)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{"msg": "Codice utente non valido"})
		return
	}

	idUser, err := strconv.Atoi(idUserStr)
	if err != nil {
		s.logger.Warningf("Valore non valido")
		w.WriteHeader(400)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{"msg": "User code not valid"})
		return
	}

	u := model.User{}
	// Ricerca utente
	res := s.gormDB.Find(&u, idUser)
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
		s.logger.Warningf("Utente non valido")
		w.WriteHeader(400)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{"msg": "User code not valid"})
		return
	}

	_ = json.NewEncoder(w).Encode(u)
}
