package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	DB_HOST     string `json:"db_host"`
	DB_PORT     string `json:"db_port"`
	DB_USERNAME string `json:"db_username"`
	DB_PASSWORD string `json:"db_password"`
	DB_NAME     string `json:"db_name"`
}

func GetConfig() (Config, error) {
	PATH := os.Getenv("MOVIE_CONF_PATH")
	dat, err := os.ReadFile(PATH)
	if err != nil {
		return Config{}, err
	}
	var result Config
	err = json.Unmarshal(dat, &result)
	return result, err
}
