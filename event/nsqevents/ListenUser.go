package nsqevents

import (
	"bigeChat/events"
	"bigeChat/msgmodels"
	"bigeChat/services"

	"github.com/buguang01/bige/messages"
)

type NsqdListenEvent struct {
	messages.NsqdMessage
	messages.LogicMessage
	Data events.LogicListen
}

//Nsq的回调
func (msg *NsqdListenEvent) NsqDirectCall() {
	msg.UserID = msg.Data.MemberID
	services.LogicEx.AddMsg(msg)
}

//调用方法
func (msg *NsqdListenEvent) MessageHandle() {
	msgmodels.NsqdTryRun(msg, msg.Data.Hander)
}
