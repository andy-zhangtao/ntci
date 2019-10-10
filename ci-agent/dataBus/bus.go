package dataBus

import (
	"io/ioutil"

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
	[language.go]
	name="image name"
	tag="latest/other tag"

*/
type dataBus struct {
	Access struct {
		Gitlab struct {
			Token string `toml:"token"`
		} `toml:"gitlab"`
	} `toml:"access"`
	Language map[string]languageRuntime `toml:"language"`
}

type languageRuntime struct {
	Name string `toml:"name"`
	Tag  string `toml:"tag"`
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

	logrus.Debugf("bus: %v", bus)
	return
}

func GetBus() *dataBus {
	return bus
}
