package conf

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var Conf = map[string]interface{}{}

/*ParseConfigFile Read and parse config file from path
*
* The config file should conform json format
 */
func ParseConfigFile(path string) error {
	configFile, err := os.Open(path)
	if err != nil {
		return err
	}

	config, err := ioutil.ReadAll(configFile)
	if err != nil {
		return err
	}

	err = json.Unmarshal(config, &Conf)
	if err != nil {
		return err
	}

	return nil
}
