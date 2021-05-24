package server

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/sirodeneko/QQBOT/coolq"
	"github.com/sirodeneko/QQBOT/util"
	"time"
)

func (c *Client) CallWithHttp(api coolq.Api, params interface{}) ([]byte, error) {
	buf := util.NewBuffer()
	defer util.PutBuffer(buf)
	_ = json.NewEncoder(buf).Encode(params)

	req := c.HttpClient.newRequest()

	resp, err := req.PostJson(fmt.Sprintf("%s/%s", c.HttpClient.httpUrl, api), buf.String())

	if err != nil {
		return nil, err
	}

	return resp.Content(), nil
}

func (c *Client) CallWithWs(api coolq.Api, params interface{}, echo string) {
	apiJson := &coolq.ApiBase{
		Action: api,
		Params: params,
		Echo:   echo,
	}

	buf := util.NewBuffer()
	defer util.PutBuffer(buf)
	_ = json.NewEncoder(buf).Encode(apiJson)

	conn := c.WsClient
	conn.Lock()
	defer conn.Unlock()
	_ = conn.SetWriteDeadline(time.Now().Add(time.Second * 15))
	if err := conn.WriteMessage(websocket.TextMessage, buf.Bytes()); err != nil {
		util.Logger.Warnf("向WS服务器 %v 调用时出现错误: %v", c.WsClient.wsUrl, err)
		_ = c.WsClient.Close()
		util.Logger.Warnf("连接到WebSocket服务器 %v 时出现错误: %v", c.WsClient.wsUrl, err)
		// 重连
		util.Logger.Info("5s后开始重连")
		time.Sleep(5 * time.Second)
	}
}
