package NsqEvents

import (
	"bigeChat/Code/ActionCode"
	"bigeChat/Events"
	"bigeChat/Models"
	"bigeChat/Service"

	"github.com/buguang01/Logger"

	"github.com/buguang01/util"
)

func Nsqd_SystemEvent(lg *Models.NsqEventBase) {
	nsqlg := new(Nsq_System)
	nsqlg.NsqEventBase = lg
	switch lg.GetActionID() {
	case ActionCode.Nsqd_SystemMsg:
		nsqlg.ParmMD.ChatMsg = util.ToString(lg.Data["ChatMsg"])
		nsqlg.ParmMD.ChatName = util.ToString(lg.Data["ChatName"])
	}
	Service.LogicEx.AddMsg(nsqlg)
}

type Nsq_System struct {
	Models.NsqLogicModel
	ParmMD Events.LogicSystem
}

func (this *Nsq_System) Run() {
	switch this.GetActionID() {
	case ActionCode.Nsqd_SystemMsg:
		this.LogicRun(this.ParmMD.Nsqd_SystemMsg)
	default:
		Logger.PError(nil, "Not Action:%d", this.ActionID)
	}
}
