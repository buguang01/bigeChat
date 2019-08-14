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

			//聊天的时候可能需要附加一些用户信息上去
			//这些用户信息可以是别的服务器写在redis里的信息
			//也可以是在注册的时候，附带过来的
			//这里的例子是从redis里读
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

			if hash != mbmd.HashKey {
				result = ConstantCode.Player_Hash_Error
				return
			}
			mbbuf, _ := json.Marshal(mbmd.ToJson())
			wsmd.ConInfo = string(mbbuf)
			wsmd.KeyID = mid
			wsmd.CloseFun = WebSocketClose
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
		chatmd.PusDal(wsmd)
	}
}
