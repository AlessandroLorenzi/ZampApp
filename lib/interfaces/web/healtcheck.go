package web

import (
	"net/http"
)

func (s Service) healthCheck(w http.ResponseWriter, _ *http.Request) {
	s.webReturn(w, 200, "ok")
}
