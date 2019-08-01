package config

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

func Element(elem string) string {
	env := os.Getenv("ECHO_ENV")
	file, err := Assets.Open("/settings.yml")
	if err != nil {
		panic(err)
	}
	buf, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	m := make(map[interface{}]interface{})
	err = yaml.Unmarshal(buf, &m)
	if err != nil {
		panic(err)
	}
	return os.ExpandEnv(m[env].(map[interface{}]interface{})[elem].(string))
}
