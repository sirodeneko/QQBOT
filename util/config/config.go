package config

import (
	"github.com/joho/godotenv"
	"github.com/sirodeneko/QQBOT/util"
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
		//return err
		util.Logger.Infof("配置文件错误,请确保环境变量无误：%v", err)
	}
	config := &Config{
		UrlStr: os.Getenv("WS_URL"),
		Token:  os.Getenv("TOKEN"),
	}
	QQBotConfig = config
	return nil
}
