package routes

import (
	"bigeChat/actioncode"
	"bigeChat/events/websocketevents"

	"github.com/buguang01/bige/messages"
)

var (
	WebSocketRoute = messages.JsonMessageHandleNew()
)

func init() {
	WebSocketRoute.SetRoute(actioncode.Nsqd_ListenUser, &websocketevents.WsocketListenEvent{})
}

func WebScoketOnline(conn *messages.WebSocketModel) {
	req := conn.Request()
	// Logger.PInfo("%v", req.Header)
	//这个方法是用来拿IP的，因为会被https代理，所以RemoteAddr不一定拿到客户的IP；
	//所以与你自己的运营沟通一下看看在哪里可以拿到IP；
	if ips, ok := req.Header["X-Forwarded-For"]; ok {
		if len(ips) > 0 {
			conn.ConInfo = ips[0]
			return
		}
	}
	conn.ConInfo = req.RemoteAddr
}
