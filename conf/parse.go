package conf

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

/*ParseJSONConfigFile Read and parse config file in JSON format from path
 */
func ParseJSONConfigFile(path string) error {
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
