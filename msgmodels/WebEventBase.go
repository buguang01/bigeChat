package msgmodels

import (
	"bigeChat/constantcode"
	"bigeChat/services"
	"net/http"

	"github.com/buguang01/bige/messages"
	"github.com/buguang01/util/threads"
)

type WebResultJSON struct {
	ActionID  uint32 //消息号
	ActionCom int
	//数据
	Data interface{} `json:"omitempty"`
	//收到的消息数据，如果有需要返回的话一般建议包装在Data中
	// MsgData interface{}
}

func NewWebResultJSON(msg messages.IHttpMessageHandle, actioncom int, jsuser H) *WebResultJSON {
	result := new(WebResultJSON)
	result.ActionID = msg.GetAction()
	result.ActionCom = actioncom
	result.Data = jsuser
	return result
}

//nsqd消息的处理包装
func WebTryRun(msg messages.IHttpMessageHandle, f func() (jsuser H, result int), w http.ResponseWriter, req *http.Request) {
	var resultmsg *WebResultJSON
	threads.Try(
		func() {
			jsuser, result := f()
			resultmsg = NewWebResultJSON(msg, result, jsuser)
		},
		func(err interface{}) {
			result := constantcode.LOGIC_Unknown_Error
			resultmsg = NewWebResultJSON(msg, result, nil)
		},
		func() {
			if resultmsg != nil {
				buff, err := services.WebEx.RouteHandle.Marshal(msg.GetAction(), resultmsg)
				if err == nil {
					w.Write(buff)
				}
			}
		},
	)

}
