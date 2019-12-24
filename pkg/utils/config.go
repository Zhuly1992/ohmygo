package utils

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Logger Logger `yaml:"logger"`
}

type Logger struct {
	SavePath string `json:"savePath" yaml:"savePath"`
	FileName string `json:"fileName" yaml:"fileName"`
	StdOut   string `json:"stdOut" yaml:"stdOut"`
	Debug    string `json:"debug" yaml:"debug"`
}

func NewConfigService() *Config {
	return &Config{}
}

func (c *Config) ReadYaml(file string) error {
	var (
		err  error
		data []byte
	)

	data, err = ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, c)
	if err != nil {
		return err
	}
	return nil
}
