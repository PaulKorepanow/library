package server

import (
	"encoding/json"
	"net/http"
)

func (s *Server) HandleMain() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := json.NewEncoder(w).Encode("hallo world!"); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
