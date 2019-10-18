package main

import "github.com/sirupsen/logrus"

//This is code builder. This builder will execute in other build container, like go, java, nodejs etc.
//
//Builder will try to clone repository from git, and parse .ntci.yml in root path. Then execute all command by order.
//
//When startup success, it will update build record status.
//
//Every job has a uniq id. Update status via this id.
//
//Builder receive key args via the flowing environment.
// NTCI_BUILDER_JID  	job id
// NTCI_BUILDER_GIT 	git repository url
// NTCI_BUILDER_BRANCH 	git branch
// NTCI_BUILDER_ADDR	build server addr
func main() {
	//Whatever the result fo build, builder should return sucess.
	if err := run(); err != nil {
		logrus.Errorf("Build Execut Failed. %s", err.Error())
	}

	logrus.Debug("")
	logrus.Debugf("===========Build Finish===========")
}
