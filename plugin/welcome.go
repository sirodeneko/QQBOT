package plugin

import (
	"github.com/sirodeneko/QQBOT/coolq"
	"github.com/sirodeneko/QQBOT/server"
	"github.com/sirodeneko/QQBOT/util"
)

// 群欢迎
func Welcome(c *server.Client) func(interface{}) {

	return func(eventData interface{}) {
		util.Logger.Debugf("收到消息，进行复读")
		msg := eventData.(*coolq.GroupIncrease)

		api := &coolq.SendMSGParams{
			MessageType: coolq.Group,
			GroupId:     msg.GroupId,
			Message:     coolq.At(msg.UserId, "") + "喵帕斯~",
			AutoEscape:  false,
		}
		c.CallWithWs(coolq.SendMSG, api, "xxx")
	}
}
