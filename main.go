package main

import (
	"github.com/sirodeneko/QQBOT/coolq"
	"github.com/sirodeneko/QQBOT/plugin"
	"github.com/sirodeneko/QQBOT/server"
	"github.com/sirodeneko/QQBOT/util"
	"github.com/sirodeneko/QQBOT/util/config"
)

const Versions = "0.0.1 Bate"

func main() {
	util.Logger.Info("QQBOT开始运行")
	util.Logger.Info("版本：" + Versions)
	err := config.LoadConfig()
	if err != nil {
		util.Logger.Warnf("配置文件加载失败：%v", err)
		return
	} else {
		util.Logger.Info("配置文件加载成功")
	}

	qqBot := coolq.NewQQBot()
	wsClient, err := server.NewWebSocketClient(config.QQBotConfig.UrlStr, config.QQBotConfig.Token, qqBot)
	if err != nil {
		util.Logger.Warnf("websocket 连接失败：%v", err)
		return
	}

	qqBot.Ues(coolq.PrivateMessageEvent, plugin.Imitator(wsClient))
	qqBot.Ues(coolq.GroupMessageEvent, plugin.GroupImitator(wsClient))
	wsClient.Listen()
}
