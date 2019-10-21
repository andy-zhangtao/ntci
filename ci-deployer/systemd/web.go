package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func start(port int) {
	r := mux.NewRouter()
	r.HandleFunc("/deploy/{name}", deploy)

	addr := fmt.Sprintf("0.0.0.0:%d", port)
	logrus.Infof("Systemd-Deployer Listen on: %s", addr)

	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	logrus.Fatal(srv.ListenAndServe())
}
