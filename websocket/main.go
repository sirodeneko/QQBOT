package websocket

import (
	"bytes"
	"github.com/gorilla/websocket"
	"github.com/sirodeneko/QQBOT/coolq"
	"github.com/sirodeneko/QQBOT/util"
	"github.com/tidwall/gjson"
	"net/http"
	"runtime/debug"
	"sync"
)

type WsClient struct {
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
		util.Logger.Warnf("连接到反向WebSocket服务器 %v 时出现错误: %v", urlStr, err)
		return nil, err
	}

	err = conn.WriteMessage(websocket.TextMessage, []byte("handshake"))
	if err != nil {
		util.Logger.Warnf("WebSocket 握手时出现错误: %v", err)
	}
	util.Logger.Infof("已连接到WebSocket服务器 %v", urlStr)

	return &WsClient{Conn: conn, Mutex: sync.Mutex{}}, nil
}

func (c *WsClient) Listen() {
	defer func() { _ = c.Close() }()

	for {
		buffer := util.NewBuffer()
		t, reader, err := c.NextReader()
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
				c.handleRequest(buffer.Bytes())
			}(buffer)
		} else {
			util.PutBuffer(buffer)
		}
	}
}

func (c *WsClient) handleRequest(payload []byte) {
	defer func() {
		if err := recover(); err != nil {
			util.Logger.Printf("处置WS命令时发生无法恢复的异常：%v\n%s", err, debug.Stack())
			_ = c.Close()
		}
	}()

	if !gjson.ValidBytes(payload) {
		util.Logger.Debugf("ws收到的数据非json格式：%v", string(payload))
		return
	}

	ret, err := coolq.CallEvent(payload)
	if err != nil {
		return
	}
	//t := strings.ReplaceAll(j.Get("action").Str, "_async", "")
	//log.Debugf("WS接收到API调用: %v 参数: %v", t, j.Get("params").Raw)
	//ret := c.apiCaller.callAPI(t, j.Get("params"))
	//if j.Get("echo").Exists() {
	//	ret["echo"] = j.Get("echo").Value()
	//}
	c.Lock()
	defer c.Unlock()
	_ = c.WriteJSON(ret)
}
