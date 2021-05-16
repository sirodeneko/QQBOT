package plugin

import (
	"github.com/sirodeneko/QQBOT/coolq"
	"github.com/sirodeneko/QQBOT/server"
	"github.com/sirodeneko/QQBOT/util"
)

func Imitator1(eventData interface{}) {
	util.Logger.Debugf("收到消息：%v", eventData.(*coolq.PrivateMessage).Message)
}

// 复读机
func Imitator(ws *server.WsClient) func(interface{}) {
	util.Logger.Debugf("收到消息，进行复读")
	return func(eventData interface{}) {
		msg := eventData.(*coolq.PrivateMessage)

		api := &coolq.SendMSGParams{
			MessageType: msg.MessageType,
			UserId:      msg.UserId,
			Message:     msg.Message,
			AutoEscape:  false,
		}
		ws.CallWithWs(coolq.SendMSG, api, "xxx")
	}
}
