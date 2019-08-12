package config

import (
	"io/ioutil"
	"os"

	"github.com/go-yaml/yaml"
)

//Param ..
var Param Configuration

// Configuration stores global configuration loaded from json file
type Configuration struct {
	ListenPort string `yaml:"listenPort"`
	DBUrl      string `yaml:"dbUrl"`
	RedisURL   string `yaml:"redisURL"`
	RedisKEY   string `yaml:"redisKEY"`
	RedisEXP   int    `yaml:"redisEXP"`
	DBType     string `yaml:"dbType"`
	JwtKEY     string `yaml:"jwtKEY"`
	Log        struct {
		FileName string `yaml:"filename"`
		Level    string `yaml:"level"`
	} `yaml:"log"`
}

// LoadConfigFromFile use to load global configuration
func LoadConfigFromFile(fn *string) {
	if err := LoadYAML(fn, &Param); err != nil {
		os.Exit(1)
	}
}

// LoadYAML load yaml format configuration
func LoadYAML(filename *string, v interface{}) error {
	raw, err := ioutil.ReadFile(*filename)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(raw, v)
	if err != nil {
		return err
	}
	return nil
}
