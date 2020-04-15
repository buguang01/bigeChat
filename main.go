package main

import (
	"bigeChat/flags"
	"bigeChat/routes"
	"bigeChat/services"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/buguang01/Logger"
	"github.com/buguang01/bige/model"
	"github.com/buguang01/bige/modules"
	"github.com/buguang01/util"
	"github.com/buguang01/util/threads"
)

func main() {
	services.Sconf = new(services.ServiceConf)

	if !flag.Parsed() {
		flag.Parse()
	}

	f, err := os.Open(*flags.Flagc)
	if err != nil {
		panic(err)
	}
	b, _ := ioutil.ReadAll(f)
	f.Close()

	json.Unmarshal(b, services.Sconf)
	Logger.Init(services.Sconf.LogLv, services.Sconf.LogPath, services.Sconf.LogMode)
	defer Logger.LogClose()

	// services.MysqlEx = model.NewMysqlAccess(&services.Sconf.DBConf)
	// defer services.MysqlEx.Close()
	// if err := services.MysqlEx.Ping(); err != nil {
	// 	Logger.PError(err, "")
	// 	return
	// }
	services.RedisEx = model.NewRedisAccess(&services.Sconf.RedisConf)
	defer services.RedisEx.Close()
	c := services.RedisEx.GetConn()
	if err := c.Err(); err != nil {
		Logger.PError(err, "")
		return
	}
	c.Close()

	services.ThGo = threads.NewThreadGo()
	// services.DBEx = modules.NewDataBaseModule(services.MysqlEx.GetDB())
	services.LogicEx = modules.NewLogicModule()
	services.TaskEx = modules.NewAutoTaskModule()
	services.NsqdEx = modules.NewNsqdModule(
		modules.NsqdSetPorts(services.Sconf.NsqdAddr...),
		modules.NsqdSetLookup(services.Sconf.NsqLookupdAddr...),
		modules.NsqdSetMyTopic(util.ToString(services.Sconf.ServiceID)),
		modules.NsqdSetMyChannelName(fmt.Sprintf("chancel_%d", services.Sconf.ServiceID)),
		modules.NsqdSetRoute(routes.NsqdRoute),
	)
	// services.WebEx = modules.NewWebModule(
	// 	modules.WebSetIpPort(services.Sconf.WebAddr),
	// 	modules.WebSetRoute(routes.WebRoute),
	// 	modules.WebSetTimeoutFunc(routes.WebTimeout),
	// )
	services.WebSocketEx = modules.NewWebSocketModule(
		modules.WebSocketSetIpPort(services.Sconf.WsAddr),
		modules.WebSocketSetRoute(routes.WebSocketRoute),
		modules.WebScoketSetOnlineFun(routes.WebScoketOnline),
	)
	services.GameEx.AddModule(
		// services.DBEx,
		services.LogicEx,
		services.TaskEx,
		services.NsqdEx,
		// services.WebEx,
		services.WebSocketEx,
	)
	InitData()
	services.GameEx.Run()

}

//InitData 用来初始化服务器的其他东西
//加载配置表
//加载数据库数据
//其他数据的加载
//初始化一些信息
//需要写入redis的操作等
func InitData() {
	// wg := new(sync.WaitGroup)
	// wg.Add(2)
	// go Manage.UserManageEx.Load(wg)
	// Conf.InitLoad(wg)
	// wg.Wait()
	// Service.GoTreand.Go(Route.AutoTask)

	// a := Models.HorselightModel{
	// 	Text:  "test",
	// 	Stime: util.GetCurrTimeSecond().Add(60 * time.Second).Unix(),
	// 	Num:   10,
	// }
	// Manage.ServerEx.SetMsgHorselight(a)
	// go http.ListenAndServe("0.0.0.0:6060", nil)
}
