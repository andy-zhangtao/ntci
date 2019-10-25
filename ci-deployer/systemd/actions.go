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

	unit := c.Units[name]
	if unit == "" {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Not find correspond unit service"))
		return
	}

	if err := restartUnit(unit); err == nil {
		return
	} else {
		logrus.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Restart unit service error"))
		return
	}
}
