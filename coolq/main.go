package coolq

import (
	"encoding/json"
	"github.com/sirodeneko/QQBOT/util"
)

func CallEvent(bytes []byte) (string, error) {

	return "", nil
}

func privateMessageEvent(bytes []byte) {
	var privateMessage PrivateMessage
	err := json.Unmarshal(bytes, &privateMessage)
	if err != nil {
		util.Logger.Debugf("私聊消息序列化失败：%v", err)
		return
	}

}
