package services

import (
	"time"

	"github.com/buguang01/util/threads"

	"github.com/buguang01/Logger"
	"github.com/buguang01/bige/model"
	"github.com/buguang01/bige/modules"
)

var (
	Sconf *ServiceConf
	// MysqlEx *model.MysqlAccess   //mysql管理器
	RedisEx *model.RedisAccess   //redis管理器
	GameEx  *modules.GameService //系统模块管理器

	// DBEx        *modules.DataBaseModule  //DB操作模块
	LogicEx *modules.LogicModule    //逻辑操作模块
	TaskEx  *modules.AutoTaskModule //内存管理模块
	NsqdEx  *modules.NsqdModule     //nsq消息队列通信模块
	// WebEx       *modules.WebModule       //HTTP通信模块
	WebSocketEx *modules.WebSocketModule //ws通信模块
	ThGo        *threads.ThreadGo
)

type ServiceConf struct {
	ServiceID   int             //游戏服务器ID
	PStatusTime time.Duration   //打印状态的时间（秒）
	LogLv       Logger.LogLevel //写日志的等级
	LogPath     string          //日志写目录
	LogMode     Logger.LogMode  //日志模式
	LogServerID string          //log服务器ID

	DBConf         model.MysqlConfigModel //Mysql管理器
	RedisConf      model.RedisConfigModel //Redis管理器
	NsqdAddr       []string               //nsqd地址组
	NsqLookupdAddr []string               //lookup 地址组
	WebAddr        string                 //Web的地址
	WsAddr         string                 //Websocket 的地址

}
