package main

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"ntci/ci-build/k8s/dataBus"
	"ntci/ci-build/k8s/deploy"
	"ntci/ci-build/k8s/store"
)

var bus *dataBus.DataBus

func init() {

	switch strings.ToLower(os.Getenv("CI_K8S_LOG_LEVEL")) {
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	default:
		logrus.SetLevel(logrus.DebugLevel)
	}

	b, err := dataBus.InitBus(os.Getenv("CI_K8S_CONF"))
	if err != nil {
		logrus.Fatalf("Parse Configure Error: %s", err.Error())
	}

	bus = b
	store.PGInit(bus)

	if err = deploy.InitK8sClient(bus); err != nil {
		logrus.Fatalf("Kubernetes Init Error: %s", err.Error())
	}
}
