package server

import (
	"github.com/gorilla/websocket"
	"github.com/sirodeneko/QQBOT/coolq"
	"github.com/sirodeneko/QQBOT/util"
	"net/http"
	"sync"
)

type WsClient struct {
	bot *coolq.QQBoT

	*websocket.Conn
	sync.Mutex
}

func NewWebSocketClient(urlStr string, token string) (*WsClient, error) {

	header := http.Header{
		"User-Agent": []string{"QQBOT"},
	}
	if token != "" {
		header["Authorization"] = []string{"Token " + token}
	}

	util.Logger.Info("开始连接ws服务器：" + urlStr)
	conn, _, err := websocket.DefaultDialer.Dial(urlStr, header)
	if err != nil {
		util.Logger.Warnf("连接到WebSocket服务器 %v 时出现错误: %v", urlStr, err)
		return nil, err
	}

	err = conn.WriteMessage(websocket.TextMessage, []byte("handshake"))
	if err != nil {
		util.Logger.Warnf("WebSocket 握手时出现错误: %v", err)
	}
	util.Logger.Infof("已连接到WebSocket服务器 %v", urlStr)

	return &WsClient{Conn: conn, Mutex: sync.Mutex{}}, nil
}
