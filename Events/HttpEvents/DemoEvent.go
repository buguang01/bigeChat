package HttpEvents

import (
	"bigeChat/Code/ConstantCode"
	"bigeChat/Events"
	"bigeChat/Service"
	"context"
	"net/http"

	"github.com/buguang01/Logger"
	"github.com/buguang01/bige/event"
	"github.com/buguang01/util"
)

func DemoEvent(et event.JsonMap, w http.ResponseWriter) {
	//这里写收到http消息的处理逻辑
	jsuser := make(event.JsonMap)
	result := ConstantCode.Success
	Service.GoTreandEx.Try(func(ctx context.Context) {
		lg := new(Events.LogicListenUser)
		lg.MemberID = et.GetMemberID()
		lg.Listen = util.NewStringAny(et["Listen"]).ToBoolV()
		result = lg.Hander(jsuser)
	}, func(err interface{}) {
		Logger.PFatal(err)
		result = ConstantCode.LOGIC_ERROR
	}, func() {
		event.HTTPReplyMsg(w, et, result, jsuser)
	})
}
