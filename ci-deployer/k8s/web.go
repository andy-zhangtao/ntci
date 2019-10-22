package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func start(port int) {
	addr := fmt.Sprintf("0.0.0.0:%d", port)
	logrus.Infof("Kubernetes-Deployer Listen on: %s", addr)

	r := mux.NewRouter()
	r.HandleFunc("/deploy/{name}", deploy)

	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	logrus.Fatal(srv.ListenAndServe())
}
