package coolq

import "github.com/sirodeneko/QQBOT/websocket"

type QQBoT struct {
	WsClient *websocket.WsClient
}

func NewQQBot(wsClient *websocket.WsClient) *QQBoT {
	bot := &QQBoT{
		WsClient: wsClient,
	}

	return bot
}
