package main

import (
	"errors"
	"fmt"
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
func clone() (err error) {
	return nil
}

//parse
//If clone() return success, then use this function to parse
//.ntci.yaml.
//If parse success, then return a filling object. Otherwise return a error
func parse(file string) (nt ntci, err error) {
	return
}

func run() (err error) {
	if err := clone(); err != nil {
		return errors.New(fmt.Sprintf("Clone Error: %s", err.Error()))
	}
}
