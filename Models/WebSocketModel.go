package Models

import (
	"bigeChat/Code/ConstantCode"
	"bigeChat/Service"

	"github.com/buguang01/Logger"
	"github.com/buguang01/bige/event"
	"github.com/buguang01/bige/threads"
	"github.com/buguang01/util"
)

type WebSocketModel struct {
	MemberID int                   //用户ID
	JsMsg    event.JsonMap         //收到的数据
	Wsmd     *event.WebSocketModel //连接信息
	Runobj   *threads.ThreadGo     //连接对应的协程管理器
}

//所在协程的KEY 因为要放到Logic上运行
func (this *WebSocketModel) KeyID() string {
	if Service.Sconf.LogicConf.InitNum == 0 {
		return util.NewStringInt(this.MemberID).ToString()
	} else if this.MemberID == 0 {
		util.NewStringInt(Service.Sconf.LogicConf.InitNum).ToString()
	}
	return util.NewStringInt(this.MemberID % Service.Sconf.LogicConf.InitNum).ToString()
}

func (this *WebSocketModel) LogicRun(f func(event.JsonMap) int) {
	jsdata := make(event.JsonMap)
	result := ConstantCode.NotLogic
	threads.Try(
		func() {
			result = f(jsdata)
			event.WebSocketReplyMsg(this.Wsmd, this.JsMsg, result, jsdata)
		}, func(err interface{}) {
			Logger.PFatal(err)
			result = ConstantCode.LOGIC_ERROR
			event.WebSocketReplyMsg(this.Wsmd, this.JsMsg, result, nil)
		}, nil,
	)
}

func NewWebSocketModel(msg event.JsonMap, wsmd *event.WebSocketModel, runobj *threads.ThreadGo) *WebSocketModel {
	result := new(WebSocketModel)
	result.JsMsg = msg
	result.Wsmd = wsmd
	result.Runobj = runobj
	result.MemberID = msg.GetMemberID()
	return result
}

//自动任务的逻辑模型,这个是在运行结束后，还要用Websocket推送给用户的自动任务逻辑
type WebSocketLogicAuto struct {
	*WebSocketModel
}

func (this *WebSocketLogicAuto) LogicRun(f func(event.JsonMap) int) {
	jsdata := make(event.JsonMap)
	result := ConstantCode.NotLogic
	threads.Try(
		func() {
			result = f(jsdata)
			if result == ConstantCode.Success {
				// this.user.PushWebSocket(jsdata)
				/*
					如果与用户的通信是建立在websocket模式下的，那可以使用这个方法给用户回复
				*/
			}
		}, func(err interface{}) {
			Logger.PFatal(err)
			result = ConstantCode.LOGIC_ERROR
		}, nil,
	)
}

//不回复用户的逻辑处理模型
type WebSocketNotResult struct {
	*WebSocketModel
}

func (this *WebSocketNotResult) LogicRun(f func(event.JsonMap) int) {
	jsdata := make(event.JsonMap)
	result := ConstantCode.NotLogic
	threads.Try(
		func() {
			result = f(jsdata)
			if result != ConstantCode.Success {
				event.WebSocketReplyMsg(this.Wsmd, this.JsMsg, result, jsdata)
			}
		}, func(err interface{}) {
			Logger.PFatal(err)
			result = ConstantCode.LOGIC_ERROR
			event.WebSocketReplyMsg(this.Wsmd, this.JsMsg, result, nil)
		}, nil,
	)
}
func NewWebSocketNotResult(msg event.JsonMap, wsmd *event.WebSocketModel, runobj *threads.ThreadGo) *WebSocketNotResult {
	result := new(WebSocketNotResult)
	result.JsMsg = msg
	result.Wsmd = wsmd
	result.Runobj = runobj
	result.MemberID = msg.GetMemberID()
	return result
}
