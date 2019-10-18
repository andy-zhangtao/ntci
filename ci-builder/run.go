package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
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

	f, err := os.Create("/git.sh")
	if err != nil {
		return err
	}

	err = t.Execute(f, gm)
	if err != nil {
		return err
	}

	f.Close()

	return exec.Command("sh", "/git.sh").Run()
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

//build
//Execute build script by order.
func build(nt ntci) (err error) {
	t := template.Must(template.New("build").Parse(buildTpl))

	f, err := os.Create("/build.sh")
	if err != nil {
		return err
	}

	err = t.Execute(f, nt)
	if err != nil {
		return err
	}

	f.Close()

	cmd := exec.Command("sh", "/build.sh")
	cmd.Env = append(os.Environ(), nt.Env...)
	cmd.Dir = fmt.Sprintf("%s/%s", gm.Root, gm.Name)
	stdout, err := cmd.StdoutPipe()
	if err = cmd.Start(); err != nil {
		logrus.Error(err)
		return
	}

	//out, err := cmd.CombinedOutput()

	logrus.Info("===========Build Log===========")
	logrus.Info("")
	scanner := bufio.NewScanner(stdout)
	//scanner := bufio.NewScanner(strings.NewReader(string(out)))
	for scanner.Scan() {
		logrus.Info(scanner.Text())
	}

	return err
}

func run() (err error) {
	if err := git(); err != nil {
		updateJobStatus(GitFailed)
		return errors.New(fmt.Sprintf("Execute Git Script Error: %s", err.Error()))
	}

	updateJobStatus(GitSuccess)
	ntciConfig := fmt.Sprintf("%s/%s/.ntci.yml", gm.Root, gm.Name)

	nt, err := parse(ntciConfig)
	if err != nil {
		updateJobStatus(NtciParseFailed)
		return errors.New(fmt.Sprintf("Parse .ntci.yml Error: %s", err.Error()))
	}

	updateJobStatus(NtciParseSuccess)
	logrus.Info(".ntci.yml")
	logrus.Infof("  language: %s", nt.Language)
	logrus.Info("  env:")
	for _, e := range nt.Env {
		logrus.Infof("    %s", e)
	}
	logrus.Info("  before build:")
	for _, b := range nt.BeforeBuild {
		logrus.Infof("    %s", b)
	}
	logrus.Info("  build:")
	for _, b := range nt.Build {
		logrus.Infof("    %s", b)
	}

	logrus.Info("  after build:")
	for _, a := range nt.AfterBuild {
		logrus.Infof("    %s", a)
	}
	logrus.Infof(" ")

	updateJobStatus(Building)
	err = build(nt)
	if err != nil {
		updateJobStatus(BuildFailed)
		return errors.New(fmt.Sprintf("Execute Build Error: %s", err.Error()))
	}

	updateJobStatus(BuildSuccess)
	return
}

func output() {
	logrus.Info("")
	f, err := os.Open("/build.log")
	if err != nil {
		logrus.Error(err.Error())
		return
	}

	defer f.Close()

	logrus.Info("===========Build Log===========")
	br := bufio.NewReader(f)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		logrus.Info(string(a))
	}

}
