package Service

import (
	"github.com/buguang01/bige/model"
	"github.com/buguang01/bige/module"
	"github.com/buguang01/bige/runserver"
	"github.com/buguang01/bige/threads"
)

var (
	MysqlEx *model.MysqlAccess         //mysql管理器
	RedisEx *model.RedisAccess         //redis管理器
	GameEx  *runserver.GameServiceBase //系统模块管理器

	// DBEx     *module.SqlDataModule //DB操作模块
	LogicEx  *module.LogicModule  //逻辑操作模块
	MemoryEx *module.MemoryModule //内存管理模块

	// HTTPEx      *module.HTTPModule      //HTTP通信模块
	WebSocketEx *module.WebSocketModule //ws通信模块
	NsqdEx      *module.NsqdModule      //nsq消息队列通信模块

	// WorkIDEx   *util.SnowFlakeIdGenerator //自动ID生成器
	GoTreandEx *threads.ThreadGo //自动任务协程管理器
)

func ServiceStop() {
	//因为服务器在关闭时，会在不收新消息的情况下，
	//把所有没有处理完成的逻辑都处理完；
	//因为NSQ可以用来处理跨服逻辑，所以需要把自己处理完的部分推送到nsq上，以便其他服务器继续处理；
	//下面这个方法是用来关掉nsq的收消息，但nsq还是可以在关服的时候，继续把消息发出去；
	NsqdEx.StopConsumer()
	//如果你使用了这个来管理你在服务内打开的其他协程；那么可以用他来关闭这些协程
	GoTreandEx.CloseWait()
}
