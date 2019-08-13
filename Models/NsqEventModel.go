package Models

import (
	"bigeChat/Code/ConstantCode"
	"bigeChat/Service"

	"github.com/buguang01/Logger"
	"github.com/buguang01/bige/event"
	"github.com/buguang01/bige/threads"
	"github.com/buguang01/util"
)

type NsqdLogicHander func(lg *NsqEventBase)

type NsqEventBase struct {
	SendID   int           //发信息用户ID
	SendSID  string        //发信息服务器（回复用的信息）
	ActionID int           //消息号
	Topic    string        //处理服务器
	MemberID int           //处理用户ID
	Step     int           //处理步骤
	Data     event.JsonMap //`json:"omitempty"` //消息数据
	// NsqdLogicHander `json:"-"`
}

func NewNsqdMessage(mid, actid int, topic string, getmid, step int, data event.JsonMap) event.INsqdMessage {
	result := new(NsqEventBase)
	result.SendID = mid
	result.ActionID = actid
	result.MemberID = getmid
	result.Topic = topic
	result.Step = step
	result.Data = data
	return result
}
func (this *NsqEventBase) GetSendID() int {
	return this.SendID
}
func (this *NsqEventBase) GetSendSID() string {
	return this.SendSID
}
func (this *NsqEventBase) SetSendSID(sid string) {
	this.SendSID = sid
}
func (this *NsqEventBase) GetActionID() int {
	return this.ActionID
}
func (this *NsqEventBase) GetData() interface{} {
	return this.Data
}
func (this *NsqEventBase) GetTopic() string {
	return this.Topic
}

type NsqLogicModel struct {
	*NsqEventBase
}

//所在协程的KEY 因为要放到Logic上运行
func (this *NsqLogicModel) KeyID() string {
	if Service.Sconf.LogicConf.InitNum == 0 {
		return util.NewStringInt(this.MemberID).ToString()
	} else if this.MemberID == 0 {
		util.NewStringInt(Service.Sconf.LogicConf.InitNum).ToString()
	}
	return util.NewStringInt(this.MemberID % Service.Sconf.LogicConf.InitNum).ToString()
}

// //调用方法
// func (this *NsqLogicBase) Run() {
// 	this.NsqdLogicHander(this)
// }

//会回复消息的
func (this *NsqLogicModel) LogicRun(f func(jsuser event.JsonMap) int) {
	result := ConstantCode.Success
	jsuser := make(event.JsonMap)
	threads.Try(
		func() {
			result = f(jsuser)
		},
		func(err interface{}) {
			Logger.PFatal(err)
			result = ConstantCode.NsqRequestFailure
		},
		func() {
			this.Data["Result"] = result
			this.Data["JsData"] = jsuser
			msg := NewNsqdMessage(this.MemberID, this.ActionID, this.SendSID, this.SendID, 1, this.Data)
			Service.NsqdEx.AddMsg(msg)
		},
	)
}
