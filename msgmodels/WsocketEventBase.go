package msgmodels

import (
	"bigeChat/constantcode"
	"bigeChat/services"

	"github.com/buguang01/bige/messages"
	"github.com/buguang01/util/threads"
)

type WsocketResult struct {
	ActionID  uint32 //消息号
	ActionCom int
	//数据
	Data interface{} `json:"omitempty"`
	//收到的消息数据，如果有需要返回的话一般建议包装在Data中
	// MsgData interface{}
}

//自定义结构的返回
//msg:收到的消息
//actioncom：处理完消息的返回码（错误码）
//jsuser:返回的自定义数据
func NewWsocketResult(msg messages.IWebSocketMessageHandle, actioncom int, jsuser H) (result *WsocketResult) {
	md := new(WsocketResult)
	result.ActionID = msg.GetAction()
	result.ActionCom = actioncom
	md.Data = jsuser
	return md
}

//nsqd消息的处理包装
func WsocketTryRun(msg messages.IWebSocketMessageHandle, ws *messages.WebSocketModel, f func() (jsuser H, result int)) {
	var resultmsg *WsocketResult
	threads.Try(
		func() {
			jsuser, result := f()
			resultmsg = NewWsocketResult(msg, result, jsuser)
		},
		func(err interface{}) {
			result := constantcode.LOGIC_Unknown_Error
			resultmsg = NewWsocketResult(msg, result, nil)
		},
		func() {
			if resultmsg != nil {
				buff, err := services.WebSocketEx.RouteHandle.Marshal(msg.GetAction(), resultmsg)
				if err == nil {
					ws.Write(buff)
				}
			}
		},
	)

}
