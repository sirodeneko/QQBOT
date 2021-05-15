package coolq

import (
	"bytes"
	"github.com/gorilla/websocket"
	"github.com/sirodeneko/QQBOT/server"
	"github.com/sirodeneko/QQBOT/util"
	"github.com/tidwall/gjson"
	"runtime/debug"
	"strconv"
	"time"
)

type QQBoT struct {
	EventFunc map[Event][]func(eventData interface{})
}

func NewQQBot(wsClient *server.WsClient) *QQBoT {
	bot := &QQBoT{
		WsClient:  wsClient,
		EventFunc: make(map[Event][]func(interface{})),
	}

	return bot
}

func (bot *QQBoT) Listen() {
	defer func() { _ = bot.WsClient.Close() }()

	for {
		buffer := util.NewBuffer()
		t, reader, err := bot.WsClient.NextReader()
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
				bot.handleRequest(buffer.Bytes())
			}(buffer)
		} else {
			util.PutBuffer(buffer)
		}
	}

	// 重连
	util.Logger.Info("5s后开始重连")
	time.Sleep(5 * time.Second)
	go bot.connect()
}

func (bot *QQBoT) connect() {
	util.Logger.Infof("开始尝试连接WebSocket服务器")
	header := http.Header{
		"X-Client-Role": []string{"API"},
		"X-Self-ID":     []string{strconv.FormatInt(c.bot.Client.Uin, 10)},
		"User-Agent":    []string{"CQHttp/4.15.0"},
	}
	if c.token != "" {
		header["Authorization"] = []string{"Token " + c.token}
	}
	conn, _, err := websocket.DefaultDialer.Dial(c.conf.API, header) // nolint
	if err != nil {
		log.Warnf("连接到反向WebSocket API服务器 %v 时出现错误: %v", c.conf.API, err)
		if c.conf.ReconnectInterval != 0 {
			time.Sleep(time.Millisecond * time.Duration(c.conf.ReconnectInterval))
			c.connectAPI()
		}
		return
	}
	log.Infof("已连接到反向WebSocket API服务器 %v", c.conf.API)
	wrappedConn := &webSocketConn{Conn: conn, apiCaller: newAPICaller(c.bot)}
	if c.conf.RateLimit.Enabled {
		wrappedConn.apiCaller.use(rateLimit(c.conf.RateLimit.Frequency, c.conf.RateLimit.Bucket))
	}
	go c.listenAPI(wrappedConn, false)
}

func (bot *QQBoT) handleRequest(payload []byte) {
	defer func() {
		if err := recover(); err != nil {
			util.Logger.Printf("处置WS命令时发生无法恢复的异常：%v\n%s", err, debug.Stack())
		}
	}()

	if !gjson.ValidBytes(payload) {
		util.Logger.Debugf("ws收到的数据非json格式：%v", string(payload))
		return
	}

	bot.CallEvent(payload)

	return
}

func (bot *QQBoT) Ues(eventName Event, fn func(eventData interface{})) {
	if bot.EventFunc[eventName] == nil {
		bot.EventFunc[eventName] = make([]func(interface{}), 0)
	}
	bot.EventFunc[eventName] = append(bot.EventFunc[eventName], fn)
}

func (bot *QQBoT) OnEvent(eventName Event, eventData interface{}) {
	if efunc := bot.EventFunc[eventName]; efunc != nil {
		for _, itemFunc := range efunc {
			itemFunc(eventData)
		}
	}
}
