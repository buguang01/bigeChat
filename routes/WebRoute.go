package routes

import (
	"bigeChat/actioncode"
	"bigeChat/events/webevents"
	"net/http"

	"github.com/buguang01/bige/messages"
)

var (
	WebRoute = messages.HttpJsonMessageHandleNew()
)

func init() {
	WebRoute.SetRoute(actioncode.Nsqd_ListenUser, &webevents.WebListenEvent{})
}

func WebTimeout(webmsg messages.IHttpMessageHandle,
	w http.ResponseWriter, req *http.Request) {
	//超时处理，这可以这里做统一处理
}
