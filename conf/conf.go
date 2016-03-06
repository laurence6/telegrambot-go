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

/*ParseConfigFile Read and parse config file from path
*
* The config file should conform json format
 */
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
