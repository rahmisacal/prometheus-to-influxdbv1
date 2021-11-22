package configuration

import (
	"../domain"
	"io/ioutil"

	"github.com/go-yaml/yaml"
)

func Configuration() *domain.Configuration {

	confContent, err := ioutil.ReadFile("config.yml")
	if err != nil {
		panic(err)
	}
	conf := &domain.Configuration{}
	if err := yaml.Unmarshal(confContent, conf); err != nil {
		panic(err)
	}
	return conf
}
