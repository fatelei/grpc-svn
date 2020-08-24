package config

import (
	"github.com/BurntSushi/toml"
	"io/ioutil"
)

type Config struct {
	DestPath string `toml:"dest_path`
}


var DefaultConfig Config


func ParseConfig(path *string) {
	var data []byte
	var err error
	if data, err = ioutil.ReadFile(*path); err != nil {
		panic(err)
	}

	if _, err = toml.Decode(string(data), &DefaultConfig); err != nil {
		panic(err)
	}
}
