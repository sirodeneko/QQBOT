package plugin

import (
	"github.com/sirodeneko/QQBOT/coolq"
	"github.com/sirodeneko/QQBOT/server"
	"github.com/sirodeneko/QQBOT/util"
	"strings"
)

// 群公告
func Notice(c *server.Client) func(interface{}) {

	return func(eventData interface{}) {
		msg := eventData.(*coolq.GroupMessage)

		if strings.HasPrefix(msg.Message, "公告") && msg.Sender.Role != "member" {
			api := &coolq.SendGroupNoticeParams{
				GroupId: msg.GroupId,
				Content: strings.Trim(msg.Message[2:], " :"),
			}
			resq, _ := c.CallWithHttp(coolq.SendGroupNotice, api)
			util.Logger.Debugf("收到返回：%v", string(resq))
		}
	}
}
