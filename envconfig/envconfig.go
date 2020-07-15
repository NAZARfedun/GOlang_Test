package envconfig

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
)

type (
	// Configuration is struct for holding service's configuration info
	Configuration struct {
		ListenPort string       `json:"ListenPort" validate:"required"`
		Log        LoggerConfig `json:"Log validate:"required"`
	}

	// LoggerConfig is a struct for holding logger configuration
	LoggerConfig struct {
		FileName string `json:"FileName" validate:"required"`
		Level    uint32 `json:"Level" required:"required"`
	}
)

func Load(configFilePath string) (err error, config Configuration) {
	if err, config = readConfigJSON(configFilePath); err != nil {
		return
	}

	return
}

// readConfigJSON reads config info from JSON file
func readConfigJSON(filePath string) (error, Configuration) {
	log.Printf("Searching JSON config file (%s)", filePath)
	var config Configuration

	contents, err := ioutil.ReadFile(filePath)
	if err == nil {
		reader := bytes.NewBuffer(contents)
		err = json.NewDecoder(reader).Decode(&config)
	}
	if err != nil {
		log.Printf("Error while reading configuration from JSON (%s) error: %s\n", filePath, err.Error())
	} else {
		log.Printf("Configuration from JSON (%s) provided\n", filePath)
	}

	return err, config
}
