package dataBus

import (
	"errors"
	"io/ioutil"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
)

type DataBus struct {
	// Service Listen Port
	Port int `toml:"port"`

	K8S k8sConf `toml:"k8s"`
	// Support build language
	Language map[string][]string `toml:"language"`
	// LanguageRuntime
	// Format:
	// map[language name] = map[tag]name
	LanguageRuntime map[string]map[string]string
	// Postgres metadata
	Postgres string `toml:"postgres"`
}

/*
k8sConf

Kubernetes metadata.
*/
type k8sConf struct {
	// K8s API Endpoint
	Endpoint string `toml:"endpoint"`
	// Namespace
	Namespace string `toml:"namespace"`
	// Config file path, if use token, this property can empty
	Config string `toml:"config"`
}

func InitBus(file string) (b *DataBus, err error) {

	if file == "" {
		file = "k8s.toml"
	}

	b = new(DataBus)
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}

	_, err = toml.Decode(string(data), b)
	if err != nil {
		return
	}

	b.LanguageRuntime = drawOffImg(b.Language)

	if err = isValid(b); err != nil {
		return
	}

	debug(b)

	return
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

func isValid(bus *DataBus) error {

	if bus.K8S.Config == "" {
		return errors.New("No Valid Kubernetes config file! ")
	}

	if len(bus.LanguageRuntime) == 0 {
		return errors.New("No Valid Language! ")
	}

	if bus.Port == 0 {
		bus.Port = 80
	}

	if bus.Postgres == "" {
		return errors.New("No Valid Postgres Addr! ")
	}

	return nil
}

func debug(bus *DataBus) {
	logrus.Debug("DATA-BUS")
	logrus.Debug("*************************************")
	logrus.Debugf("Listen on: %d", bus.Port)

	logrus.Debug("Kubernetes: ")
	logrus.Debugf("  Endpoint: %s", bus.K8S.Endpoint)
	logrus.Debugf("  Namespace: %s", bus.K8S.Namespace)
	logrus.Debugf("  Config: %s", bus.K8S.Config)

	logrus.Debug("")

	for l, v := range bus.LanguageRuntime {
		logrus.Debugf("Language: %s", l)
		for tag, image := range v {
			logrus.Debugf("  %s:%s", image, tag)
		}
	}

	logrus.Debug("Postgres")
	logrus.Debugf("  Endpoint: %s", bus.Postgres)
	logrus.Debug("")

	logrus.Debug("*************************************")
}
