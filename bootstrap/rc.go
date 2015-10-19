package bootstrap

import (
	"io/ioutil"
	"os"

	"gopkg.in/leyra/toml.v1"
)

type RcConfig struct {
	DatabaseEnable string
}

func NewRcConfig() *RcConfig {
	return new(RcConfig)
}

func (rc *RcConfig) Apply() *RcConfig {
	f, err := os.Open("./etc/rc.conf")
	if err != nil {
		panic(err)
	}

	defer f.Close()
	buf, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	if err := toml.Unmarshal(buf, rc); err != nil {
		panic(err)
	}

	return rc
}
