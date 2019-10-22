package main

import (
	"os"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
	"ntci/ci-deployer/model/config"
)

var c config.K8sConfigure

func init() {

	switch strings.ToLower(os.Getenv("NTCI_DEPLOY_LOG_LEVEL")) {
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	default:
		logrus.SetLevel(logrus.DebugLevel)
	}

	if os.Getenv("NTCI_DEPLOY_CONFIG") != "" {
		if _, err := toml.DecodeFile(os.Getenv("NTCI_DEPLOY_CONFIG"), &c); err != nil {
			logrus.Fatal("Configure Init Failed! ")
		}
	}

	if err := initK8sClient(c.K8sConf); err != nil {
		logrus.Fatalf("K8s Client Init Failed!  %s ", err.Error())
	}

	if c.Port == 0 {
		c.Port = 80
	}

	logrus.Debug("-------Configure-------")
	logrus.Debugf("Listen On: %d", c.Port)
	logrus.Debugf("K8s Conf Path: %s", c.K8sConf)
	for key, value := range c.Service {
		logrus.Debugf("Name: %s Deployment: %s Namespace: %s", key, value, c.Namespaces[key])
	}
	logrus.Debug("-----------------------")

}
