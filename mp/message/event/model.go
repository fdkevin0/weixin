package event

import (
	"WeChat-golang/mp/core"
)

type Message struct {
	*core.Message        //MsgType event
	Event         string `xml:"Event"` //事件类型，subscribe(订阅)、unsubscribe(取消订阅)
}

type SubscribeEvent struct {
	*core.Message //Event 事件类型，subscribe(订阅)、unsubscribe(取消订阅)
}

type QrCodeEventMessage struct {
	*core.Message //
}
