package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	UrlStr string
	Token  string
}

var QQBotConfig *Config

func LoadConfig() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	config := &Config{
		UrlStr: os.Getenv("WS_URL"),
		Token:  os.Getenv("TOKEN"),
	}
	QQBotConfig = config
	return nil
}
