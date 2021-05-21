package coolq

import (
	"github.com/sirodeneko/QQBOT/util"
	"github.com/tidwall/gjson"
	"runtime/debug"
	"sync"
)

type QQBoT struct {
	EventFunc     map[Event][]func(eventData interface{})
	EventFuncLock sync.RWMutex
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
	bot.EventFuncLock.Lock()
	defer bot.EventFuncLock.Unlock()

	if bot.EventFunc[eventName] == nil {
		bot.EventFunc[eventName] = make([]func(interface{}), 0)
	}
	bot.EventFunc[eventName] = append(bot.EventFunc[eventName], fn)
}

func (bot *QQBoT) OnEvent(eventName Event, eventData interface{}) {
	bot.EventFuncLock.RLock()
	defer bot.EventFuncLock.RUnlock()

	if efunc := bot.EventFunc[eventName]; efunc != nil {
		for _, itemFunc := range efunc {
			if itemFunc != nil {
				go func(fn func(eventData interface{})) {
					defer func() {
						if err := recover(); err != nil {
							util.Logger.Printf("插件运行失败：%v\n", err)
						}
					}()
					// 运行事件函数
					fn(eventData)

				}(itemFunc)

			}
		}
	}
}
