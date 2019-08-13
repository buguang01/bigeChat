package Service

import (
	"github.com/buguang01/Logger"
	"github.com/buguang01/bige/model"
	"github.com/buguang01/bige/module"
	"github.com/buguang01/bige/runserver"
)

type ServiceConf struct {
	GameConf  runserver.GameConfigModel
	DBConf    model.MysqlConfigModel
	RedisConf model.RedisConfigModel

	SqlConf    module.SqlDataConfig
	LogicConf  module.LogicConfig
	MemoryConf module.MemoryConfig

	HttpConf    module.HTTPConfig
	NsqdConf    module.NsqdConfig
	WsocketConf module.WebSocketConfig

	WorkerId int64 //自ID的服务器编号

	LogLv       Logger.LogLevel //写日志的等级
	LogPath     string          //日志写目录
	LogMode     Logger.LogMode  //日志模式
	LogServerID string          //log服务器ID
}

var Sconf *ServiceConf
