package config

import (
	"github.com/joho/godotenv"
	"github.com/sirodeneko/QQBOT/util"
	"os"
)

type Config struct {
	WsUrl   string
	Token   string
	HttpUrl string
}

var QQBotConfig *Config

func LoadConfig() error {
	err := godotenv.Load()
	if err != nil {
		//return err
		util.Logger.Infof("配置文件错误,请确保环境变量无误：%v", err)
	}
	config := &Config{
		WsUrl:   os.Getenv("WS_URL"),
		Token:   os.Getenv("TOKEN"),
		HttpUrl: os.Getenv("HTTP_URL"),
	}
	QQBotConfig = config
	return nil
}
