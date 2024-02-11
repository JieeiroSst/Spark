package config

import (
	"bytes"
)

type Config struct {
	Listen    string            `json:"listen" yaml:"listen"`
	Salt      string            `json:"salt" yaml:"salt"`
	Auth      map[string]string `json:"auth" yaml:"auth"`
	Log       *Log              `json:"log" yaml:"log"`
	SaltBytes []byte            `json:"-" yaml:"-"`
}
type Log struct {
	Level string `json:"level" yaml:"level"`
	Path  string `json:"path" yaml:"path"`
	Days  uint   `json:"days" yaml:"days"`
}

// COMMIT is hash of this commit, for auto upgrade.
var COMMIT = ``
var BuiltPath = `./built/%v_%v`

func ReadConfigFile(file string) *Config {
	config := &Config{
		Listen: ":8000",
		Salt:   "123456abcdef",
		Auth: map[string]string{
			"username": "password",
		},
		Log: &Log{
			Level: "info",
			Path:  "./logs",
			Days:  7,
		},
	}
	config.SaltBytes = []byte(config.Salt)
	config.SaltBytes = append(config.SaltBytes, bytes.Repeat([]byte{25}, 24)...)
	config.SaltBytes = config.SaltBytes[:24]

	return config
}
