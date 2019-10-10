package dataBus

import (
	"io/ioutil"

	"github.com/BurntSushi/toml"
)

var bus *dataBus

/*
DataBus

The Global Configure Data
Other package can get configure from this object.

[access.gitlab]
	token="xxx" // The gitlab access token. Generate by user.
*/
type dataBus struct {
	Access struct {
		Gitlab struct {
			Token string `toml:"token"`
		} `toml:"gitlab"`
	} `toml:"access"`
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

	if _, err := toml.Decode(string(data), bus); err != nil {
		return
	}

	return
}

func GetBus() *dataBus {
	return bus
}
