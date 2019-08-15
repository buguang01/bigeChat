package WebSocketEvents

import (
	"bigeChat/ChatModels"
	"bigeChat/Code/ConstantCode"
	"bigeChat/Conf"
	"bigeChat/Dal"
	"bigeChat/Service"
	"fmt"

	"github.com/buguang01/util"

	"github.com/buguang01/bige/event"
	"github.com/buguang01/bige/threads"
)

//监听聊天频道
func WsEventChatPus(et event.JsonMap, wsmd *event.WebSocketModel, runobj *threads.ThreadGo) {
	jsdata := make(event.JsonMap)
	result := ConstantCode.Timeout
	threads.Try(
		func() {
			arr := et.GetArray("ChatLi")
			user := wsmd.ConInfo.(Dal.UserModel)
			for _, name := range arr {
				chatmd := ChatModels.ChatEx.GetChat(util.ToString(name))
				chatmd.PusAdd(wsmd)
				user.ChatLi[chatmd.GetChatName()] = true

			}
			result = ConstantCode.Success
		},
		nil,
		func() {
			event.WebSocketReplyMsg(wsmd, et, result, jsdata)
		},
	)
}

//取消监听聊天频道
func WsEventChatCancelPus(et event.JsonMap, wsmd *event.WebSocketModel, runobj *threads.ThreadGo) {
	jsdata := make(event.JsonMap)
	result := ConstantCode.Timeout
	threads.Try(
		func() {
			arr := et.GetArray("ChatLi")
			user := wsmd.ConInfo.(Dal.UserModel)
			for _, name := range arr {
				chatmd := ChatModels.ChatEx.GetChat(util.ToString(name))
				chatmd.PusDel(wsmd)
				delete(user.ChatLi, chatmd.GetChatName())
			}
			result = ConstantCode.Success
		},
		nil,
		func() {
			event.WebSocketReplyMsg(wsmd, et, result, jsdata)
		},
	)
}

//心跳
func WsEventChatDiDa(et event.JsonMap, wsmd *event.WebSocketModel, runobj *threads.ThreadGo) {
	event.WebSocketReplyMsg(wsmd, et, ConstantCode.Success, nil)
}

//发聊天信息
func WsEventChatSendMsg(et event.JsonMap, wsmd *event.WebSocketModel, runobj *threads.ThreadGo) {
	jsdata := make(event.JsonMap)
	result := ConstantCode.Timeout
	threads.Try(
		func() {
			//检查是不是被禁言了
			redis := Service.RedisEx.GetConn()
			defer redis.Close()
			if reply, err := redis.Get(fmt.Sprintf("ChatBan%d", wsmd.KeyID)); err == nil && reply != nil {
				result = ConstantCode.Chat_Player_Ban
				return
			}
			msg := new(ChatModels.ChatMessage)
			msg.UserInfo = wsmd.ConInfo.(Dal.UserModel).UserInfo
			msg.MemberID = wsmd.KeyID
			msg.ChatNode = Conf.FilterChack(util.ToString(et["ChatMsg"]))
			msg.CreateTime = util.GetCurrTimeSecond()
			name := util.ToString(et["ChatName"])
			chatmd := ChatModels.ChatEx.GetChat(name)
			chatmd.AddMsg(msg)
			result = ConstantCode.Success
		},
		nil,
		func() {
			event.WebSocketReplyMsg(wsmd, et, result, jsdata)
		},
	)
}
