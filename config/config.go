package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type RedisConfig struct {
	IsCache   bool   `yaml:"IsCache"`
	MaxIdle   int    `yaml:"MaxIdle"`
	MaxActive int    `yaml:"MaxActive"`
	Port      int    `yaml:"Port"`
	User      string `yaml:"Username"`
	URL       string `yaml:"URL"`
	Password  string
}

type AppConfig struct {
	BaseUrl    string      `yaml:"base_url"`
	HashLength int         `yaml:"hash_length"`
	Redis      RedisConfig `yaml:"redis"`
}

func ReadConfigFile(filepath string) ([]byte, error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func ParseYamlConfig(yamlConfig []byte) (*AppConfig, error) {
	config := &AppConfig{}
	if err := yaml.Unmarshal(yamlConfig, config); err != nil {
		return &AppConfig{}, err
	}
	return config, nil
}
