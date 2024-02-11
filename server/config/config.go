package config

import (
	"bytes"
	"fmt"
	"os"

	"github.com/ghodss/yaml"
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

func ReadConfYAML(filename string) (*Config, error) {
	buffer, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = yaml.Unmarshal(buffer, &config)
	if err != nil {
		fmt.Printf("err: %v\n", err)

	}
	return config, nil
}

func ReadConfJSON(filename string) (*Config, error) {
	buffer, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = yaml.Unmarshal(buffer, &config)
	if err != nil {
		fmt.Printf("err: %v\n", err)

	}
	return config, nil
}

func ReadConfigFile(file string) *Config {
	config, err := ReadConfYAML(file)
	if err != nil {
		return nil
	}
	if config.Log == nil {
		config.Log = &Log{
			Level: `info`,
			Path:  `./logs`,
			Days:  7,
		}
	}

	if len(config.Salt) > 24 {
		return nil
	}
	config.SaltBytes = []byte(config.Salt)
	config.SaltBytes = append(config.SaltBytes, bytes.Repeat([]byte{25}, 24)...)
	config.SaltBytes = config.SaltBytes[:24]

	return config
}
