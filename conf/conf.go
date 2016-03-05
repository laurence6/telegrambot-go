package conf

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var Conf map[string]interface{}

func init() {
	Conf = map[string]interface{}{}
}

func ParseConfigFile(path string) (err error) {
	configFile, err := os.Open(path)
	if err != nil {
		return
	}

	config, err := ioutil.ReadAll(configFile)
	if err != nil {
		return
	}

	json.Unmarshal(config, &Conf)

	return
}
