package bootstrap

import (
	"gopkg.in/leyra/toml.v1"
	bootutil "leyra/bootstrap/util"
)

type RcConfig struct {
	DatabaseEnable string
}

func NewRcConfig() *RcConfig {
	return new(RcConfig)
}

func (rc *RcConfig) Apply() *RcConfig {
	buf := bootutil.ConfigBuffer("./etc/rc.conf")

	if err := toml.Unmarshal(buf, rc); err != nil {
		panic(err)
	}

	return rc
}
