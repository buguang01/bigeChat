package Routes

import (
	"bigeChat/Code/ActionCode"
	"bigeChat/Events/NsqEvents"
	"bigeChat/Models"
	"bigeChat/Service"
	"context"

	"github.com/buguang01/Logger"
	"github.com/buguang01/bige/event"
)

func init() {
	NsqRouteList = make(map[int]Models.NsqdLogicHander)

	NsqRouteList[ActionCode.Nsqd_SystemMsg] = NsqEvents.Nsqd_SystemEvent
}

var (
	NsqRouteList map[int]Models.NsqdLogicHander
)

func NsqdInit() {
	Service.NsqdEx.RouteFun = NsqdHander
	Service.NsqdEx.GetNewMsg = func() event.INsqdMessage {
		return new(Models.NsqEventBase)
	}
}

func NsqdHander(msg event.INsqdMessage) {
	//这里就要写消息的确认处理方法
	//和消息处理方法的运行
	//fmt.Println(msg,msg.GetActionID(),msg.(*GameLogic.LogicRoute).GetActionID())
	hander, ok := NsqRouteList[msg.GetActionID()]
	if ok {
		//把运行逻辑放到按服务器来源分的协程中
		// logicmd := NsqLogic.NewLogicRoute(hander, msg)
		logicmd := msg.(*Models.NsqEventBase)
		// logicmd.MemberID = logicmd.Data.GetMemberID()
		// logicmd.NsqdLogicHander = hander
		// Service.LogicEx.AddMsg(logicmd)
		Service.GoTreandEx.Try(func(ctx context.Context) {
			hander(logicmd)
		}, nil, nil)
		/*
			正如你看到的那样，虽然我们在nsq协程上收到了消息，但是会丢到对应的logic协程上运行
			因为这个方法，会因为收到一个新消息然后就会开一个协程进行处理
		*/

		// hander(msg)
	} else {
		Logger.PError(nil, "Nsq Action:%d,not hander.", msg.GetActionID())
	}
}
