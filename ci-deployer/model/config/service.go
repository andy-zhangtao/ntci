package config

type Configure struct {
	Port  int               `toml:port`
	Units map[string]string `toml:"units"`
}

type K8sConfigure struct {
	Port       int               `toml:port`
	Service    map[string]string `toml:"units"`
	Namespaces map[string]string `toml:"namespace"`
	K8sConf    string            `toml:"k8s"`
}
