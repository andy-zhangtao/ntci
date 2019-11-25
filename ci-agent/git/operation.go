package git

import (
	"errors"

	"github.com/sirupsen/logrus"
)

/*
ParseProject

Parse Project via parse .ntci.yml file.
*/
func ParseAndExecuteBuild(g GitOperation) (n Ntci, err error) {
	n, err = g.FetchNtCI()
	if err != nil {
		return
	}

	logrus.Debug("ntci:")
	logrus.Debugf("  Language: %s", n.Language)
	logrus.Debugf("  Env: %v", n.Env)
	logrus.Debugf("  Build: %v", n.Build)
	logrus.Debugf("  AfterBuild: %v", n.AfterBuild)
	logrus.Debugf("  BeforeBuild: %v", n.BeforeBuild)
	logrus.Debugf("  Deploy: %v", n.Deployer)

	if !g.VerifyNtci(n) {
		err = errors.New("Invalid ntci configure ")
		return
	}

	err = g.InvokeBuildService(n)

	return
}
