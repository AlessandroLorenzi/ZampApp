package web

import (
	"encoding/json"
	"net/http"
	"zampapp/lib/platform/jwt"
)

func (s *Service) login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		s.webReturn(w, 400, "impossible to get parse form")
	}
	type loginPost struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	var lp loginPost
	if len(r.Form) != 1 {
		s.webReturn(w, 400, "empty document")
		return
	}
	for key := range r.Form {
		err := json.Unmarshal([]byte(key), &lp)
		if err != nil {
			s.webReturn(w, 400, "document not valid")
			return
		}
		continue
	}

	u, err := s.useCasesService.Login(lp.Login, lp.Password)
	if err != nil {
		s.webReturn(w, 400, "login not valid", responseContent{
			"error": err.Error(),
		})
		return
	}

	s.webReturn(w, 200, "token generated",
		responseContent{
			"user": u,
			"jwt":  jwt.GenerateUserJWT(u),
		},
	)

}
