package config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"launchpad.net/goyaml"
	"strings"
)

var Configs map[interface{}]interface{}

func ReadConfigBytes(data []byte) error {
	return goyaml.Unmarshal(data, &Configs)
}

func ReadConfigFile(filePath string) error {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	return goyaml.Unmarshal(data, &Configs)
}

func Get(key string) (interface{}, error) {
	keys := strings.Split(key, ":")
	conf, ok := Configs[keys[0]]
	if !ok {
		return nil, errors.New(fmt.Sprintf("key %s not found", key))
	}
	for _, k := range keys[1:] {
		conf, ok = conf.(map[interface{}]interface{})[k]
		if !ok {
			return nil, errors.New(fmt.Sprintf("key %s not found", key))
		}
	}
	return conf, nil
}

func GetString(key string) string {
	value, _ := Get(key)
	return value.(string)
}
