package plugin

import (
	"github.com/sirodeneko/QQBOT/coolq"
	"github.com/sirodeneko/QQBOT/util"
)

func Imitator(eventData interface{}) {
	util.Logger.Debugf("收到消息：%v", eventData.(*coolq.PrivateMessage).Message)
}
