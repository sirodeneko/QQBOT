package coolq

import (
	"encoding/json"
	"fmt"
	"github.com/sirodeneko/QQBOT/util"
	"github.com/tidwall/gjson"
)

func (bot *QQBoT) CallEvent(bytes []byte) {

	postType := gjson.GetBytes(bytes, "post_type").String()
	messageType := gjson.GetBytes(bytes, "message_type").String()
	noticeType := gjson.GetBytes(bytes, "notice_type").String()
	subType := gjson.GetBytes(bytes, "sub_type").String()

	if gjson.GetBytes(bytes, "meta_event_type").String() != "heartbeat" {
		fmt.Println("收到消息", gjson.ParseBytes(bytes).String())
	}

	if postType == "message" {
		if messageType == "private" {
			if subType == "friend" {
				//"post_type":    "message",
				//"message_type": "private",
				//"sub_type":     "friend",
				bot.privateMessageEvent(bytes)
			}
		} else if messageType == "group" {
			//"post_type":    "message",
			//"message_type": "group",
			bot.groupMessageEvent(bytes)
		}
	} else if postType == "notice" {
		if noticeType == string(GroupIncreaseNotice) {
			//"post_type":    "notice",
			//"notice_type": "group_increase",
			bot.groupIncreaseEvent(bytes)
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

func (bot *QQBoT) groupMessageEvent(bytes []byte) {
	var groupMessage GroupMessage
	err := json.Unmarshal(bytes, &groupMessage)
	if err != nil {
		util.Logger.Debugf("私聊消息序列化失败：%v", err)
		return
	}

	bot.OnEvent(GroupMessageEvent, &groupMessage)
}

func (bot *QQBoT) groupIncreaseEvent(bytes []byte) {
	var groupIncrease GroupIncrease
	err := json.Unmarshal(bytes, &groupIncrease)
	if err != nil {
		util.Logger.Debugf("私聊消息序列化失败：%v", err)
		return
	}

	bot.OnEvent(GroupIncreaseEvent, &groupIncrease)
}
