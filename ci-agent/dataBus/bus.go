package dataBus

import (
	"io/ioutil"

	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
	"ntci/ci-agent/store"
)

var bus *dataBus

/*
DataBus

The Global Configure Data
Other package can get configure from this object.

# Choose build mode
build-mode="single"

language=[
		"go"
	]
[access.gitlab]
	token="xxx" // The gitlab access token. Generate by user.


	# The language runtime image name


# User can use default agent.(single/k8s)
# If user wants custom agent, it should implement ci-grpc/build/v1.proto
[build]
	[build.single]
	addr=""
*/
type dataBus struct {
	Access struct {
		Gitlab struct {
			Token string `toml:"token"`
		} `toml:"gitlab"`
	} `toml:"access"`

	//Language map[string][]string `toml:"language"`
	//Language []string `toml:"language"`

	// LanguageRuntime
	// Format:
	// map[language name] = map[tag]name
	//LanguageRuntime map[string]map[string]string

	// BuildMode
	// This mode must exist in Build Sections
	BuildMode string `toml:"build-mode"`
	// Build
	// Build Sections contains all valid build service
	Build map[string]buildService `toml:"build"`
	// WebPort Default 8000
	WebPort int `toml:"web_port"`
	// GateWayPort Default 8001
	GateWayPort int `toml:"gateway_port"`
	// Postgres metadata
	Postgres string `toml:"postgres"`

	// Deployer Valid Deployer Addr
	// e.g.
	// k8s="xxxx"
	Deployer map[string]string `toml:"deployer"`

	Pb *store.PGBus

	JobStatus chan *Status
}

type buildService struct {
	Addr string `toml:"addr"`
}

// Status
// Ntci Job Status
type Status struct {
	User   string
	Name   string
	Branch string
	Id     int
	Stauts int
}

/*
InitDataBus

Read and parse ntci.toml. If parse success, will filling data into bus point.

Other module can invoke GetBus get this var.
*/
func InitDataBus(file string) (err error) {

	bus = new(dataBus)
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}

	_, err = toml.Decode(string(data), bus)
	if err != nil {
		return
	}

	bus.Pb = store.PG(bus.Postgres)
	bus.JobStatus = make(chan *Status, 100)

	debug()
	return
}

func GetBus() *dataBus {
	return bus
}

func debug() {
	logrus.Info("DATA-BUS")
	logrus.Info("*************************************")
	logrus.Infof("Web Server Listen: %d", bus.WebPort)
	logrus.Infof("Gateway Server Listen: %d", bus.GateWayPort)
	logrus.Infof("Postgresql Addr: %s", bus.Postgres)

	if bus.Access.Gitlab.Token != "" {
		logrus.Infof("GitLab Token: %s", bus.Access.Gitlab.Token)
	}

	logrus.Infof("Build Mode: %s", bus.BuildMode)
	for m, svc := range bus.Build {
		logrus.Infof("  %s:[%s]", m, svc.Addr)
	}

	logrus.Info("Deployer:")
	for n, addr := range bus.Deployer {
		logrus.Infof("  %s:[%s]", n, addr)
	}

	logrus.Info("*************************************")
}
