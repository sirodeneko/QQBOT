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
	client, err := server.NewClient(config.QQBotConfig, qqBot)
	if err != nil {
		util.Logger.Warnf("websocket 连接失败：%v", err)
		return
	}

	//通过Ues函数进行插件的注册，参数1：事件类型，参数2：运行插件的方法
	//参数2 接受的类型为func(eventData interface{})，对于复杂的逻辑，可通过闭包进行
	//变量的传递
	//qqBot.Ues(coolq.PrivateMessageEvent, plugin.Imitator(client))
	//qqBot.Ues(coolq.GroupMessageEvent, plugin.GroupImitator(client))
	//qqBot.Ues(coolq.GroupMessageEvent, plugin.Welcome(client))
	qqBot.Ues(coolq.GroupMessageEvent, plugin.Notice(client))

	client.WsListen()
}
