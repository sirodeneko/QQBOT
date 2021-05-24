package plugin

import (
	"github.com/sirodeneko/QQBOT/coolq"
	"github.com/sirodeneko/QQBOT/server"
	"github.com/sirodeneko/QQBOT/util"
	"strings"
)

func Imitator1(eventData interface{}) {
	util.Logger.Debugf("收到消息：%v", eventData.(*coolq.PrivateMessage).Message)
}

// 复读机
func Imitator(c *server.Client) func(interface{}) {

	return func(eventData interface{}) {
		util.Logger.Debugf("收到消息，进行复读")
		msg := eventData.(*coolq.PrivateMessage)

		api := &coolq.SendMSGParams{
			MessageType: msg.MessageType,
			UserId:      msg.UserId,
			Message:     msg.Message,
			AutoEscape:  false,
		}
		c.CallWithWs(coolq.SendMSG, api, "xxx")
	}
}

// 群复读
// 复读机
func GroupImitator(c *server.Client) func(interface{}) {

	return func(eventData interface{}) {
		util.Logger.Debugf("收到消息，进行复读")
		msg := eventData.(*coolq.GroupMessage)

		// 对于特定的消息不进行复读
		if strings.Contains(msg.Message, "邹") {
			api := &coolq.SendMSGParams{
				MessageType: msg.MessageType,
				GroupId:     msg.GroupId,
				Message:     coolq.At(msg.UserId, "") + "哼 ╭(╯^╰)╮ 又欺负小舞",
				AutoEscape:  false,
			}
			c.CallWithWs(coolq.SendMSG, api, "xxx")
			return
		}

		api := &coolq.SendMSGParams{
			MessageType: msg.MessageType,
			GroupId:     msg.GroupId,
			Message:     coolq.At(msg.UserId, "") + msg.Message,
			AutoEscape:  false,
		}
		c.CallWithWs(coolq.SendMSG, api, "xxx")
	}
}

//
