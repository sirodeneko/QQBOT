package coolq

import (
	"encoding/json"
	"github.com/sirodeneko/QQBOT/util"
	"github.com/tidwall/gjson"
)

type Event string

const (
	PrivateMessageEvent Event = "privateMessageEvent"
)

func (bot *QQBoT) CallEvent(bytes []byte) {

	postType := gjson.GetBytes(bytes, "post_type").String()
	messageType := gjson.GetBytes(bytes, "message_type").String()
	subType := gjson.GetBytes(bytes, "sub_type").String()

	if postType == "message" {
		if messageType == "private" {
			if subType == "friend" {
				//"post_type":    "message",
				//"message_type": "private",
				//"sub_type":     "friend",
				bot.privateMessageEvent(bytes)
			}
		}
	}

	return
}

func (bot *QQBoT) privateMessageEvent(bytes []byte) {
	var privateMessage PrivateMessage
	err := json.Unmarshal(bytes, &privateMessage)
	if err != nil {
		util.Logger.Debugf("私聊消息序列化失败：%v", err)
		return
	}

	bot.OnEvent(PrivateMessageEvent, &privateMessage)
}
