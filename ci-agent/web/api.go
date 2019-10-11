package web

import (
	"net/http"

	"github.com/gorilla/mux"
)

/**
api is restful api struct.

In this struct, there has the flowing property:

+ apiMap : A map contains path and handler function.
+ router : The Mux router point, use for register handler functions.

*/
type api struct {
	apiMap map[string]func(http.ResponseWriter, *http.Request)
	router *mux.Router
}

func (a *api) init() {
	a.apiMap = make(map[string]func(http.ResponseWriter, *http.Request))
	a.router = mux.NewRouter()
}

const (
	APIVersion = "/v1"
)
