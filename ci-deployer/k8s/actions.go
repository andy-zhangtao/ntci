package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func deploy(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]

	logrus.Debugf("Start Deploy Service. %s", name)

	service := c.Service[name]
	if service == "" {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Not find correspond unit service"))
		return
	}

	namespace := c.Namespaces[name]
	if namespace == "" {
		namespace = "default"
	}

	if err := restartService(service, namespace); err == nil {
		return
	} else {
		logrus.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Restart unit service error"))
		return
	}
}
