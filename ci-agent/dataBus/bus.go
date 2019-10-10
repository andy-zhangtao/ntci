package dataBus

import (
	"io/ioutil"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
)

var bus *dataBus

/*
DataBus

The Global Configure Data
Other package can get configure from this object.

[access.gitlab]
	token="xxx" // The gitlab access token. Generate by user.

[language]
	# The language runtime image name
	go=[
		"name:tag"
		]

# Choose build mode
# User can use default agent.(single/k8s)
# If user wants custom agent, it should implement the flowing apis:
# 	- GET /_ping health check
# 	- POST /trigger execute
build-mode="single"
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

	Language map[string][]string `toml:"language"`

	// LanguageRuntime
	// Format:
	// map[language name] = map[tag]name
	LanguageRuntime map[string]map[string]string

	// BuildMode
	// This mode must exist in Build Sections
	BuildMode string `toml:"build-mode"`
	// Build
	// Build Sections contains all valid build service
	Build map[string]buildService `toml:"build"`
}

type buildService struct {
	Addr string `toml:"addr"`
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

	bus.LanguageRuntime = drawOffImg(bus.Language)
	logrus.Debugf("bus: %v", bus)
	return
}

func GetBus() *dataBus {
	return bus
}

// drawOffImg
// Convert Language string to struct.
func drawOffImg(lan map[string][]string) map[string]map[string]string {
	runtime := make(map[string]map[string]string)

	for key, value := range lan {
		image := make(map[string]string)

		for _, v := range value {
			if strings.Contains(v, ":") {
				_v := strings.Split(v, ":")
				image[_v[0]] = _v[1]
			} else {
				image["latest"] = v
			}
		}

		runtime[key] = image
	}

	return runtime
}
