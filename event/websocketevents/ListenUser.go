package websocketevents

import (
	"bigeChat/events"
	"bigeChat/msgmodels"
	"bigeChat/services"

	"github.com/buguang01/bige/messages"
)

type WsocketListenEvent struct {
	messages.WebScoketMessage
	messages.LogicMessage
	Data    events.LogicListen
	wsocket *messages.WebSocketModel
}

func (msg *WsocketListenEvent) WebSocketDirectCall(ws *messages.WebSocketModel) {
	msg.wsocket = ws
	msg.UserID = msg.MemberID
	services.LogicEx.AddMsg(msg)

	// panic(errors.New("not virtual func."))

}

//调用方法
func (msg *WsocketListenEvent) MessageHandle() {
	msgmodels.WsocketTryRun(msg, msg.wsocket, msg.Data.Hander)
}
