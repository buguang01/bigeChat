package Conf_test

import (
	"bigeChat/Conf"
	"testing"

	"github.com/buguang01/Logger"
)

func TestFilter(t *testing.T) {
	Logger.Init(Logger.LogLevelmainlevel, "../bin/logs", Logger.LogModeFmt)
	defer Logger.LogClose()
	CONFIG_PATH := "../bin/config"
	conf := new(Conf.ConfigModel)
	Logger.PStatus("Config Load Filter.")
	conf.ConfigFilter = Conf.InitConfigFilter(CONFIG_PATH)
	Conf.ConfExample = conf

	for i := 0; i < 100000; i++ {
		arrstr := []string{"我操 你妈B，我  日  你 就是喜欢你", "中国人就是神", "日本人就是狗", "毛主席万万岁", "天才自在功",
			"小孩子别BB", "民进党 是真 的傻吗"}
		for _, data := range arrstr {
			Conf.FilterChack(data)
			// Logger.PInfo(data)
			// Logger.PInfo(datan)
		}
	}

}
