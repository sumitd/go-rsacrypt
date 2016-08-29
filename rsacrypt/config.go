package rsacrypt

import (
	"encoding/json"
	"flag"
	"log"
	"os"
)

// Info from config file
type Configuration struct {
	PublicKeyURL string `json:"publickey_api_url"`
	Mailserver   string `json:"mail_server"`
	ApiKey       string `json:"git_api_key"`
	FromEmail    string `json:"from_email"`
}

// config file location from flag
var config = flag.String("config", "/etc/go-rsacrypt/config.json", "Path to config file")

// Reads info from config file
func ReadConfig() Configuration {
	if *config == "" {
		flag.PrintDefaults()
	}

	configFile, err := os.Open(*config)
	if err != nil {
		log.Fatalf("Unable to open config file. %s", err)
	}
	decoder := json.NewDecoder(configFile)

	var configuration Configuration
	if err := decoder.Decode(&configuration); err != nil {
		log.Fatalf("Unable to read config. %s", err)
	}

	return configuration
}
