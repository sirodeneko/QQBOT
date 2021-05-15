package server

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/sirodeneko/QQBOT/coolq"
	"github.com/sirodeneko/QQBOT/util"
	"time"
)

func (ws *WsClient) CallWithHttp(api coolq.Api, params interface{}) {
	// TODO
}

func (ws *WsClient) CallWithWs(api coolq.Api, params interface{}, echo string) {
	apiJson := &coolq.ApiBase{
		Action: api,
		Params: params,
		Echo:   echo,
	}

	buf := util.NewBuffer()
	defer util.PutBuffer(buf)
	_ = json.NewEncoder(buf).Encode(apiJson)

	conn := ws
	conn.Lock()
	defer conn.Unlock()
	_ = conn.SetWriteDeadline(time.Now().Add(time.Second * 15))
	if err := conn.WriteMessage(websocket.TextMessage, buf.Bytes()); err != nil {
		util.Logger.Warnf("向WS服务器 %v 调用时出现错误: %v", ws.wsUrl, err)
		_ = ws.Close()
		util.Logger.Warnf("连接到WebSocket服务器 %v 时出现错误: %v", ws.wsUrl, err)
		// 重连
		util.Logger.Info("5s后开始重连")
		time.Sleep(5 * time.Second)
	}
}
