package configurations

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Configuration struct {
	Hostname string `json:"hostname"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetConfiguration() *Configuration {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()

	var config Configuration

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalln(err)
	}

	json.Unmarshal(byteValue, &config)

	return &config
}
