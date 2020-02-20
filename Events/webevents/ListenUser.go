package webevents

import (
	"bigeChat/events"
	"bigeChat/msgmodels"
	"net/http"

	"github.com/buguang01/bige/messages"
)

type WebListenEvent struct {
	messages.WebMessage
	Data events.LogicListen
}

//HTTP的回调
func (msg *WebListenEvent) HttpDirectCall(w http.ResponseWriter, req *http.Request) {
	msgmodels.WebTryRun(msg, msg.Data.Hander, w, req)
}
