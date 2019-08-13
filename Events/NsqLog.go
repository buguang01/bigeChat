package Events

import (
	"bigeChat/Code/ActionCode"
	"bigeChat/Models"
	"bigeChat/Service"

	"github.com/buguang01/bige/event"
)

//sendNsqdLogMD写入日志
func sendNsqdLogMD(logmd *Models.LogInfoMD) {
	msg := event.NewNsqdMessage(logmd.MID, ActionCode.Nsqd_Log, Service.Sconf.LogServerID, logmd)
	Service.NsqdEx.AddMsg(msg)
}

// func Nlog_Register(user *userobj.UserModel) {
// 	dt := util.GetCurrTime()
// 	logmd := Dal.NewLogInfoByNum("PassPort", "Register", "",
// 		user.MemberID(), Service.Sconf.GameConf.ServiceID, dt,
// 		0, user.Member.CreateTime.Unix())
// 	sendNsqdLogMD(logmd)
// }
