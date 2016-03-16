package conf

import (
	"bufio"
	"encoding/json"
	"io"
	"os"
)

/*ParseJSONConfigFile Read and parse config file in JSON format from path
 */
func ParseJSONConfigFile(path string) error {
	configFile, err := os.Open(path)
	if err != nil {
		return err
	}

	decoder := json.NewDecoder(bufio.NewReader(configFile))

	for {
		if err = decoder.Decode(&Conf); err == io.EOF {
			break
		} else if err != nil {
			return err
		}
	}

	return nil
}
