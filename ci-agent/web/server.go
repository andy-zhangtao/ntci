package web

import (
	"fmt"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"ntci/ci-agent/gitlab"
)

/**
Run is web server entry function.

It will start up a web server , and listen on specify port.

This server has a ping api(/_ping), this api use for health check and return all valid APIs.
*/
func Run(port int) {

	ap := new(api)
	ap.init()
	ap.gatherAPI()
	ap.registerAPI()

	ap.run(port)
}

/**
gatherAPI

Generate api path and handler function map.
*/
func (a *api) gatherAPI() {
	a.apiMap[wrapAPI("/version")] = version
	a.apiMap[wrapAPI("/gitlab/push")] = new(gitlab.Service).GitCallBack
}

/**
register API

Iterator path and handler from API Map, then register it into mux.router
*/
func (a *api) registerAPI() {
	for path, hander := range a.apiMap {
		a.router.HandleFunc(path, hander)
	}

	a.router.HandleFunc("/_ping", a.ping)
}

/**
wrapAPI

Wrap API Path use api version. Api should invoke this function before register.
*/
func wrapAPI(path string) string {
	return fmt.Sprintf("%s%s", APIVersion, path)
}

func (a *api) run(port int) {
	addr := fmt.Sprintf("%s:%d", "0.0.0.0", port)
	srv := &http.Server{
		Handler: a.router,
		Addr:    addr,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logrus.Fatal(srv.ListenAndServe())
}
