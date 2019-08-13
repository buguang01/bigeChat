package WebSocketEvents

import (
	"bigeChat/Events"
	"bigeChat/Models"

	"github.com/buguang01/bige/event"
	"github.com/buguang01/bige/threads"
	"github.com/buguang01/util"
)

//WsDomeEvent
// et 是收到消息的Json对象
// wsmd 是websocket的连接对象
// runobj 为这个连接建的协程管理器，会在连接关闭时关闭
func WsDomeEvent(et event.JsonMap, wsmd *event.WebSocketModel, runobj *threads.ThreadGo) {
	//WebSocket的消息处理
	lg := new(Ws_ListenUser)
	lg.WebSocketModel = Models.NewWebSocketModel(et, wsmd, runobj)
	lg.ParmMD.MemberID = lg.MemberID
	lg.ParmMD.Listen = util.NewStringAny(et["Listen"]).ToBoolV()
}

//通过这个类，把逻辑包起来，放到Logic的协程上运行，然后回复到NSQ上
type Ws_ListenUser struct {
	*Models.WebSocketModel
	ParmMD Events.LogicListenUser
}

func (this *Ws_ListenUser) Run() {
	this.LogicRun(this.ParmMD.Hander)
}
