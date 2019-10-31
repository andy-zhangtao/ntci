package git

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

deploy:
- branch:
	xxxx
*/
type Ntci struct {
	Language    string                            `yaml:"language"`
	Env         []string                          `yaml:"env"`
	BeforeBuild []string                          `yaml:"before_build"`
	Build       []string                          `yaml:"build"`
	AfterBuild  []string                          `yaml:"after_build"`
	Deployer    map[string]map[string]interface{} `yaml:"deploy"`
}
