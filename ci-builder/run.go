package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"text/template"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

/*
Ntci is model of .ntci.yml struct
The classic .ntci.yml content:

language: go[:tag]
env:
- key=value

before_build:
- shell command

build:
- shell command

after_build:
- shell_command

*/
type ntci struct {
	Language    string   `yaml:"language"`
	Env         []string `yaml:"env"`
	BeforeBuild []string `yaml:"before_build"`
	Build       []string `yaml:"build"`
	AfterBuild  []string `yaml:"after_build"`
}

//clone
//This function clone git project into local. And try to
//parse .ntci.yaml.
//
//root 	Build Root Path
//url 	Repository URL
//branch Remote Branch
//
//Execute two command:
// git clone jid url
// git checkout -b branch origin/branch
func git() (err error) {
	t := template.Must(template.New("git").Parse(cloneTpl))

	f, err := os.Create("/build.sh")
	if err != nil {
		return err
	}

	err = t.Execute(f, gm)
	if err != nil {
		return err
	}

	f.Close()

	return exec.Command("sh", "/build.sh").Run()
}

//parse
//If clone() return success, then use this function to parse
//.ntci.yaml.
//If parse success, then return a filling object. Otherwise return a error
func parse(file string) (nt ntci, err error) {
	if _, err = os.Stat(file); os.IsNotExist(err) {
		err = errors.New(file + " not exist.")
		return
	}

	data, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(data, &nt)
	if err != nil {
		return
	}

	return
}

func run() (err error) {
	if err := git(); err != nil {
		return errors.New(fmt.Sprintf("Execute Git Script Error: %s", err.Error()))
	}

	ntciConfig := fmt.Sprintf("%s/%s/.ntci.yml", gm.Root, gm.Name)

	nt, err := parse(ntciConfig)
	if err != nil {
		return errors.New(fmt.Sprintf("Parse .ntci.yml Error: %s", err.Error()))
	}

	logrus.Debug(".ntci.yml")
	logrus.Debugf("  language: %s", nt.Language)
	logrus.Debugf("  env: %s", nt.Env)
	logrus.Debugf("  build: %s", nt.Build)
	logrus.Debugf("  before build: %s", nt.BeforeBuild)
	logrus.Debugf("  after build: %s", nt.AfterBuild)
	logrus.Debug(" ")
	return
}
