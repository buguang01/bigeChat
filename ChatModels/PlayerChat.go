package ChatModels

import "github.com/buguang01/bige/event"

//玩家私聊管理器
type PlayerChatMD struct {
	MemberID int                   //用户ID
	Chatli   map[string]IChatMD    //私聊频道
	Wsmd     *event.WebSocketModel //连接信息
}

func NewPlayerChatMD(mid int) *PlayerChatMD {
	result := new(PlayerChatMD)
	result.MemberID = mid
	result.Chatli = make(map[string]IChatMD)
	return result
}

//频道添加
func (this *PlayerChatMD) ChatAdd(ctmd IChatMD) {
	this.Chatli[ctmd.GetChatName()] = ctmd
}

//频道删除
func (this *PlayerChatMD) ChatDel(ctmd IChatMD) {
	delete(this.Chatli, ctmd.GetChatName())
}
