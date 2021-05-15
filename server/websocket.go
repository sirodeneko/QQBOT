package server

import (
	"bytes"
	"github.com/gorilla/websocket"
	"github.com/sirodeneko/QQBOT/coolq"
	"github.com/sirodeneko/QQBOT/util"
	"net/http"
	"sync"
	"time"
)

type WsClient struct {
	bot *coolq.QQBoT

	wsUrl string
	token string

	*websocket.Conn
	sync.Mutex
}

func NewWebSocketClient(urlStr string, token string, bot *coolq.QQBoT) (*WsClient, error) {
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

	wsClient := &WsClient{
		bot:   bot,
		wsUrl: urlStr,
		token: token,
		Conn:  conn,
		Mutex: sync.Mutex{},
	}

	return wsClient, nil
}

func (ws *WsClient) Listen() {
	defer func() { _ = ws.Close() }()

	for {
		buffer := util.NewBuffer()
		t, reader, err := ws.NextReader()
		if err != nil {
			util.Logger.Warnf("websocket出现错误: %v", err)
			break
		}
		_, err = buffer.ReadFrom(reader)
		if err != nil {
			util.Logger.Warnf("websocket出现错误: %v", err)
			break
		}
		if t == websocket.TextMessage {
			go func(buffer *bytes.Buffer) {
				defer util.PutBuffer(buffer)
				ws.bot.HandleRequest(buffer.Bytes())
			}(buffer)
		} else {
			util.PutBuffer(buffer)
		}
	}

	// 重连
	util.Logger.Info("5s后开始重连")
	time.Sleep(5 * time.Second)
	go ws.connect()
}

func (ws *WsClient) connect() {
	util.Logger.Infof("开始尝试连接WebSocket服务器")
	header := http.Header{
		"User-Agent": []string{"QQBOT"},
	}
	if ws.token != "" {
		header["Authorization"] = []string{"Token " + ws.token}
	}
	conn, _, err := websocket.DefaultDialer.Dial(ws.wsUrl, header) // nolint
	if err != nil {
		util.Logger.Warnf("连接到WebSocket服务器 %v 时出现错误: %v", ws.wsUrl, err)
		// 重连
		util.Logger.Info("5s后开始重连")
		time.Sleep(5 * time.Second)
		return
	}
	util.Logger.Infof("已连接到反向WebSocket API服务器 %v", ws.wsUrl)
	ws.Conn = conn
	go ws.Listen()
}
