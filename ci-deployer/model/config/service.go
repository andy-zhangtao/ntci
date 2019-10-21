package config

type Configure struct {
	Port  int               `toml:port`
	Units map[string]string `toml:"units"`
}
