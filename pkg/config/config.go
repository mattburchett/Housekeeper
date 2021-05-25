package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// Config - This struct will hold configuration components.
type Config struct {
	PlexPyURL      string `json:"plexPyURL"`
	PlexPyAPIKey   string `json:"plexPyAPIKey"`
	PlexToken      string `json:"plexToken"`
	PlexHost       string `json:"plexHost"`
	PlexPort       int    `json:"plexPort"`
	TelegramToken  string `json:"telegramToken"`
	TelegramChatID string `json:"telegramChatID"`
	ServerName     string `json:"serverName"`
	SonarrURL      string `json:"sonarrURL"`
	SonarrAPIKey   string `json:"sonarrAPIKey"`
	ExcludeList    string `json:"excludeList"` // ExcludeList will be checked against any section.
}

//GetConfig gets the configuration values for the api using the file in the supplied configPath.
func GetConfig(configPath string) (Config, error) {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return Config{}, fmt.Errorf("could not find the config file at path %s", configPath)
	}
	log.Println("Loading Configuration File: " + configPath)
	return loadConfigFromFile(configPath)
}

//if the config loaded from the file errors, no defaults will be loaded and the app will exit.
func loadConfigFromFile(configPath string) (conf Config, err error) {
	file, err := os.Open(configPath)
	if err != nil {
		log.Printf("Error opening config file: %v", err)
	} else {
		defer file.Close()

		err = json.NewDecoder(file).Decode(&conf)
		if err != nil {
			log.Printf("Error decoding config file: %v", err)
		}
	}

	return conf, err
}
