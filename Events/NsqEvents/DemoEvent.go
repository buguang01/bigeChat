package NsqEvents

import (
	"bigeChat/Events"
	"bigeChat/Models"
	"bigeChat/Service"

	"github.com/buguang01/util"
)

//这个方法跑在nsq的新协程上
//监听用户所有日志；设置后，可以把指定用户的日志都输出到一个独立的文件
func Nsqd_ListenUser(lg *Models.NsqEventBase) {
	nsqlg := new(Nsq_ListenUser)
	nsqlg.NsqEventBase = lg
	nsqlg.ParmMD.MemberID = lg.MemberID
	nsqlg.ParmMD.Listen = util.NewStringAny(lg.Data["Listen"]).ToBoolV()
	Service.LogicEx.AddMsg(nsqlg)
}

//通过这个类，把逻辑包起来，放到Logic的协程上运行，然后回复到NSQ上
type Nsq_ListenUser struct {
	Models.NsqLogicModel
	ParmMD Events.LogicListenUser
}

func (this *Nsq_ListenUser) Run() {
	this.LogicRun(this.ParmMD.Hander)
}
