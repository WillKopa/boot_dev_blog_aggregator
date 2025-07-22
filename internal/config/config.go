package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	DB_url 				string
	Current_user_name 	string
}

func get_config_path() (string, error) {
	const configFileName = "/.gatorconfig.json"

	path, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error getting home directory %v", err)
	}

	path += configFileName

	return path, nil
}

func Read() (*Config, error) {
	path, err := get_config_path()
	if err != nil {
		return &Config{}, err
	}

	file_content, err := os.ReadFile(path)
	if err != nil {
		return &Config{}, fmt.Errorf("error reading %v. received error %v", path, err)
	}

	var config Config
	err = json.Unmarshal(file_content, &config)

	if err != nil {
		return &Config{}, fmt.Errorf("error converting json to Config struct: %v", err)
	}

	return &config, nil
}

func (c *Config) SetUser(user string) error {
	c.Current_user_name = user
	path, err := get_config_path()
	if err != nil {
		return fmt.Errorf("error getting path: %v", err)
	}

	data, err := json.Marshal(c)
	if err != nil {
		return fmt.Errorf("error marshaling struct: %v", err)
	}
	
	err = os.WriteFile(path, data, os.ModePerm)

	if err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}
	return nil
}
