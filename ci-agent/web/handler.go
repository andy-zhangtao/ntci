package web

import (
	"encoding/json"
	"net/http"
)

/**
ALL restful api handler functions.
*/

/**
version return currently api version
*/
func version(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(APIVersion))
}

/**
ping
Health Check Function
*/
func (a *api) ping(w http.ResponseWriter, r *http.Request) {
	var ps []string
	for path := range a.apiMap {
		ps = append(ps, path)
	}

	json.NewEncoder(w).Encode(ps)
}
