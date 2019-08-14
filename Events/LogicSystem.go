package Events

import (
	"bigeChat/ChatModels"
	"bigeChat/Code/ConstantCode"

	"github.com/buguang01/util"

	"github.com/buguang01/bige/event"
)

type LogicSystem struct {
	ChatMsg  string
	ChatName string
}

//其他服务器发的消息统一用户ID都是0
func (this *LogicSystem) Nsqd_SystemMsg(jsuser event.JsonMap) (result int) {
	result = ConstantCode.Success
	msg := new(ChatModels.ChatMessage)
	msg.UserInfo = ""
	msg.MemberID = 0
	msg.ChatNode = this.ChatMsg
	msg.CreateTime = util.GetCurrTimeSecond()
	chatmd := ChatModels.ChatEx.GetChat(this.ChatName)
	chatmd.AddMsg(msg)
	return result
}
