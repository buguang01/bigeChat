package WebSocketEvents

import (
	"bigeChat/ChatModels"
	"bigeChat/Code/ConstantCode"
	"bigeChat/Dal"
	"bigeChat/Service"
	"fmt"

	"github.com/buguang01/bige/json"

	"github.com/buguang01/Logger"
	"github.com/buguang01/bige/event"
	"github.com/buguang01/bige/threads"
)

func WsEventRegister(et event.JsonMap, wsmd *event.WebSocketModel, runobj *threads.ThreadGo) {
	jsdata := make(event.JsonMap)
	result := ConstantCode.Timeout
	threads.Try(
		func() {
			redis := Service.RedisEx.GetConn()
			defer redis.Close()
			mid := et.GetMemberID()
			hash := et.GetHash()

			//没找到用户的时候，从redis里读
			rinfo, err := redis.DictGet("MemberIDList", fmt.Sprintf("%d", mid))
			if err != nil {
				result = ConstantCode.Player_Not_Exist

				return
			}
			// fmt.Println(string(rinfo.([]byte)))
			Logger.PDebug(string(rinfo.([]byte)))
			mbmd := new(Dal.MemberMD)
			err = json.Unmarshal(rinfo.([]byte), &mbmd)
			if err != nil {
				result = ConstantCode.Player_Not_Exist
				Logger.PError(err, "")
				return
			}

			//读到后，再次检查
			if hash != mbmd.HashKey {
				result = ConstantCode.Player_Hash_Error
				return
			}
			mbbuf, _ := json.Marshal(mbmd.ToJson())
			user := Dal.NewUserModel()
			user.UserInfo = string(mbbuf)
			wsmd.ConInfo = user
			wsmd.KeyID = mid
			wsmd.CloseFun = WebSocketClose
			ChatModels.ChatEx.RegisterPlayer(mid, wsmd)
			result = ConstantCode.Success
		},
		nil,
		func() {
			event.WebSocketReplyMsg(wsmd, et, result, jsdata)
		},
	)
}

func WebSocketClose(wsmd *event.WebSocketModel) {
	user, wsok := wsmd.ConInfo.(Dal.UserModel)
	if !wsok {
		return
	}
	for name, _ := range user.ChatLi {
		chatmd := ChatModels.ChatEx.GetChat(name)
		chatmd.PusDel(wsmd)
	}
	ChatModels.ChatEx.OfflinePlayer(wsmd.KeyID, wsmd)
}
