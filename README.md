# QQBOT


## 📖 简介

`QQBOT`是基于 [cqhttp ](https://github.com/Mrs4s/go-cqhttp) 的QQ机器人,cqhttp负责与qq服务器进行交互，QQBOT通过接受cqhttp的推送，提供便捷的api进行事件的处理.

## 🛠 使用

并项目的使用要求您会编写简单的Golang和了解cqhttp消息的规则。

- 运行cqhttp,详细方法见[cqhttp文档](https://docs.go-cqhttp.org/)

- 将.env.examples文件重命名为.env 并修改参数

- 编写你的代码

  示例1：复读机

  plugin/Imitator.go

  ```
  
  func Imitator(ws *server.WsClient) func(interface{}) {
  
  	return func(eventData interface{}) {
  		// eventData为对应消息的结构体，经过断言后可进行使用
  		util.Logger.Debugf("收到消息，进行复读")
  		msg := eventData.(*coolq.PrivateMessage)
  		
  		// 构建消息的结构体
  		api := &coolq.SendMSGParams{
  			MessageType: msg.MessageType,
  			UserId:      msg.UserId,
  			Message:     msg.Message,
  			AutoEscape:  false,
  		}
  		// 调用提供的api
  		// 参数1方法类型 2参数结构图 3echo
  		ws.CallWithWs(coolq.SendMSG, api, "xxx")
  	}
  }
  ```
- 注册你的插件
  main.go

  ```
  //通过Ues函数进行插件的注册，参数1：时间类型，参数2：运行插件的方法
  //参数2 接受的类型为func(eventData interface{})，对于复杂的逻辑，可通过闭包进行
  //变量的传递，如本例，通过传入wsClient进行了消息的发送
  qqBot.Ues(coolq.PrivateMessageEvent, plugin.Imitator(wsClient))
  ```

- 运行main.go

## 😥 警告

 - 项目正在开发中，项目结构可能会有大变化
 - 欢迎pr,以提供更多的api和event