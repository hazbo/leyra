package bootstrap

import (
	"gopkg.in/leyra/toml.v1"

	bootutil "leyra/bootstrap/util"
)

type ServerConfig struct {
	Port string
}

func NewServerConfig() *ServerConfig {
	return new(ServerConfig)
}

func (sc *ServerConfig) Apply() *ServerConfig {
	buf := bootutil.ConfigBuffer("./etc/server.conf")

	if err := toml.Unmarshal(buf, sc); err != nil {
		panic(err)
	}

	return sc
}
