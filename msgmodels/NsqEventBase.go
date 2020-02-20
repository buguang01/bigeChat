package msgmodels

import (
	"bigeChat/constantcode"
	"bigeChat/services"

	"github.com/buguang01/bige/messages"
	"github.com/buguang01/util/threads"
)

type NsqResult struct {
	messages.NsqdMessage
	ActionCom int
	//数据
	Data interface{} `json:"omitempty"`
	//收到的消息数据，如果有需要返回的话一般建议包装在Data中
	// MsgData interface{}
}

//自定义结构的返回
//msg:收到的消息
//com：处理完消息的返回码（错误码）
//v:返回的自定义数据
func NewNsqResult(msg messages.INsqMessageHandle, com int, v interface{}) (result messages.INsqdResultMessage) {
	md := new(NsqResult)
	md.ActionID = msg.GetAction()
	md.SendSID = msg.GetTopic()
	md.SendUserID = msg.GetSendUserID()
	md.Data = v
	md.ActionCom = com
	result = md
	return
}

//json数据的返回
//msg:收到的消息
//com：处理完消息的返回码（错误码）
//v:返回的自定义数据
func NewNsqResultJSON(msg messages.INsqMessageHandle, com int, v H) (result messages.INsqdResultMessage) {
	md := new(NsqResult)
	md.ActionID = msg.GetAction()
	md.SendSID = msg.GetTopic()
	md.SendUserID = msg.GetSendUserID()
	md.Data = v
	md.ActionCom = com
	return md
}

//nsqd消息的处理包装
func NsqdTryRun(msg messages.INsqMessageHandle, f func() (jsuser H, result int)) {
	var resultmsg messages.INsqdResultMessage
	threads.Try(
		func() {
			jsuser, result := f()
			resultmsg = NewNsqResultJSON(msg, result, jsuser)
		},
		func(err interface{}) {
			result := constantcode.LOGIC_Unknown_Error
			resultmsg = NewNsqResultJSON(msg, result, nil)
		},
		func() {
			if resultmsg != nil {
				services.NsqdEx.AddMsg(resultmsg)
			}
		},
	)

}
