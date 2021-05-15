package coolq

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/sirodeneko/QQBOT/util"
	"time"
)

type Api string

type apiBase struct {
	Action Api         `json:"action"`
	Params interface{} `json:"params"`
	Echo   string      `json:"echo"`
}

func (bot *QQBoT) CallWithHttp(api Api, params interface{}) {
	// TODO
}

func (bot *QQBoT) CallWithWs(api Api, params interface{}, echo string) {
	apiJson := &apiBase{
		Action: api,
		Params: params,
		Echo:   echo,
	}

	buf := util.NewBuffer()
	defer util.PutBuffer(buf)
	_ = json.NewEncoder(buf).Encode(apiJson)

	conn := bot.WsClient
	conn.Lock()
	defer conn.Unlock()
	_ = conn.SetWriteDeadline(time.Now().Add(time.Second * 15))
	if err := conn.WriteMessage(websocket.TextMessage, buf.Bytes()); err != nil {
		log.Warnf("向WS服务器 %v 推送Event时出现错误: %v", c.eventConn.RemoteAddr().String(), err)
		_ = c.eventConn.Close()
		if c.conf.ReconnectInterval != 0 {
			time.Sleep(time.Millisecond * time.Duration(c.conf.ReconnectInterval))
			c.connectEvent()
		}
	}
}

func (bot *QQBoT) sendMSG() {
	api := apiBase{
		Action: "",
		Params: nil,
		Echo:   "",
	}
}
