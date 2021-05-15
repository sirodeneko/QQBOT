package coolq

import (
	"github.com/sirodeneko/QQBOT/util"
	"github.com/tidwall/gjson"
	"runtime/debug"
)

type QQBoT struct {
	EventFunc map[Event][]func(eventData interface{})
}

func NewQQBot() *QQBoT {
	bot := &QQBoT{
		EventFunc: make(map[Event][]func(interface{})),
	}

	return bot
}

func (bot *QQBoT) HandleRequest(payload []byte) {
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
