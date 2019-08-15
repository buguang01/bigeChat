package main

import (
	"bigeChat/Conf"
	"bigeChat/Flag"
	"bigeChat/Routes"
	"bigeChat/Service"
	"flag"
	"io/ioutil"
	"os"
	"sync"

	"github.com/buguang01/Logger"
	"github.com/buguang01/bige/json"
	"github.com/buguang01/bige/model"
	"github.com/buguang01/bige/module"
	"github.com/buguang01/bige/runserver"
	"github.com/buguang01/bige/threads"
)

func main() {
	Service.Sconf = new(Service.ServiceConf)
	var conf = Service.Sconf

	if !flag.Parsed() {
		flag.Parse()
	}

	f, err := os.Open(*Flag.Flagc)
	if err != nil {
		panic(err)
	}
	b, _ := ioutil.ReadAll(f)
	f.Close()

	json.Unmarshal(b, &conf)
	Logger.Init(conf.LogLv, conf.LogPath, conf.LogMode)
	defer Logger.LogClose()

	// Service.MysqlEx = model.NewMysqlAccess(&conf.DBConf)
	// defer Service.MysqlEx.Close()
	// if err := Service.MysqlEx.Ping(); err != nil {
	// 	Logger.PError(err, "")
	// 	return
	// }
	Service.RedisEx = model.NewRedisAccess(&conf.RedisConf)
	defer Service.RedisEx.Close()
	c := Service.RedisEx.GetConn()
	if err := c.Err(); err != nil {
		Logger.PError(err, "")
		return
	}
	c.Close()

	Service.GameEx = runserver.NewGameService(&conf.GameConf)
	Service.GameEx.ServiceStopHander = Service.ServiceStop

	// Service.DBEx = module.NewSqlDataModule(&conf.SqlConf, Service.MysqlEx.GetDB())
	Service.LogicEx = module.NewLogicModule(&conf.LogicConf)
	Service.MemoryEx = module.NewMemoryModule(&conf.MemoryConf)

	Service.WebSocketEx = module.NewWSModule(&conf.WsocketConf)
	Service.NsqdEx = module.NewNsqdModule(&conf.NsqdConf, conf.GameConf.ServiceID)

	// Service.HTTPEx = module.NewHTTPModule(&conf.HttpConf)

	// Service.WorkIDEx = util.NewIDGenerator().SetWorkerId(conf.WorkerId)
	// Service.WorkIDEx.Init()
	Service.GoTreandEx = threads.NewThreadGo()

	// Service.GameEx.AddModule(Service.DBEx)
	Service.GameEx.AddModule(Service.NsqdEx)
	Service.GameEx.AddModule(Service.LogicEx)
	Service.GameEx.AddModule(Service.MemoryEx)
	Service.GameEx.AddModule(Service.WebSocketEx)
	// Service.GameEx.AddModule(Service.HTTPEx)
	InitData()
	Service.GameEx.Run()

}

//InitData 用来初始化服务器的其他东西
//加载配置表
//加载数据库数据
//其他数据的加载
//初始化一些信息
//需要写入redis的操作等
func InitData() {
	Routes.WebSocketInit()
	Routes.NsqdInit()
	wg := new(sync.WaitGroup)
	wg.Add(1)
	threads.GoTry(func() {
		Conf.InitLoad(wg)
	}, nil, nil)
	wg.Wait()

	// Routes.HTTPInit()

	// go http.ListenAndServe("0.0.0.0:6060", nil)
}
