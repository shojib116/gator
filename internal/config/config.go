package config

import (
	"encoding/json"
	"log"
	"os"
	"path"
)

type Config struct {
	DBUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

const configFile = ".gatorconfig.json"

func (cfg *Config) SetUser(username string) error {
	cfg.CurrentUserName = username
	return write(*cfg)
}

func Read() (Config, error) {
	filePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	byteData, err := os.ReadFile(filePath)
	if err != nil {
		return Config{}, err
	}

	var config Config
	err = json.Unmarshal(byteData, &config)
	if err != nil {
		log.Fatal(err)
	}

	return config, nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return path.Join(homeDir, configFile), nil
}

func write(cfg Config) error {
	filePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	byteData, err := json.Marshal(cfg)
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = os.WriteFile(filePath, byteData, 0666)

	return nil
}
