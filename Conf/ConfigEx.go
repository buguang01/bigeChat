package Conf

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"
	"sync"

	"github.com/buguang01/Logger"
	"github.com/buguang01/util"
)

/*
这是我自己写的一个读配置表的逻辑，里面解配置表的部分要按你自己的修改
希望对你有帮助
*/
const (
	CONFIG_PATH = "config"
)

var (
	ConfExample *ConfigModel
)

//ConfigModel配置表
type ConfigModel struct {
	*ConfigFilter
}

func InitLoad(wg *sync.WaitGroup) {
	defer wg.Done()
	conf := new(ConfigModel)
	Logger.PStatus("Config Load Filter.")
	conf.ConfigFilter = InitConfigFilter(CONFIG_PATH)
	ChackConfig(conf)
	ConfExample = conf
	Logger.PStatus("Config Load All.")
}

func ChackConfig(cf *ConfigModel) {
	result := false

	// for _, v := range cf.HBossList {
	// 	if _, ok := cf.ItemList[v.RewardItemID]; !ok {
	// 		Logger.PDebug("HookBoss.%d.RewardItemID=%d", v.UID, v.RewardItemID)
	// 		result = true
	// 	}
	// 	for _, k := range v.RewardByPop.Data {
	// 		if _, ok := cf.DropList[k]; !ok {
	// 			Logger.PDebug("HookBoss.%d.RewardByPop.DropID=%d", v.UID, k)
	// 			result = true
	// 		}
	// 	}
	// }

	if result {
		panic(errors.New("Chack Config Fatal."))
	}
}

func loadfile(filestr string) *ConfigTableRows {
	file, err := os.Open(filestr)
	if err != nil {
		panic(err)
	}
	barr, _ := ioutil.ReadAll(file)
	bstr := string(barr)
	// fmt.Println(bstr)
	strr := []rune(bstr)
	for i, _ := range strr {
		strr[i] = strr[i] ^ 255
	}
	bstr = string(strr)
	bstr = strings.ReplaceAll(
		strings.ReplaceAll(
			strings.ReplaceAll(bstr, " ", ""),
			"\"", "",
		),
		"\r", "",
	)
	// fmt.Println(bstr)
	strtemp := strings.Split(bstr, "\n")
	table := NewConfigTable(len(strtemp))
	for _, drow := range strtemp {
		rowstr := strings.Split(drow, "\t")
		if len(rowstr) <= 1 && rowstr[0] == "" {
			if table.Count() == 0 {
				return nil
			} else {
				return table
			}
		} else if strings.Contains(rowstr[0], "#") {
			continue
		} else {
			if table.Columns == nil {
				table.Columns = rowstr
			} else {
				table.AddRow(rowstr)
			}
		}
	}
	return table
}

type ConfigTableRows struct {
	Columns []string
	Rows    []*ConfigRow
}

func (this *ConfigTableRows) Count() int {
	return len(this.Rows)
}

func NewConfigTable(rownum int) *ConfigTableRows {
	result := new(ConfigTableRows)
	result.Rows = make([]*ConfigRow, 0, rownum)
	return result
}

func (this *ConfigTableRows) AddRow(arr []string) {
	row := new(ConfigRow)
	row.Columns = this.Columns
	row.Row = make([]*util.String, len(arr))
	for i, v := range arr {
		row.Row[i] = util.NewString(v)
	}
	// row.Row = arr
	this.Rows = append(this.Rows, row)
}

type ConfigRow struct {
	Columns []string
	Row     []*util.String
}

func (this *ConfigRow) GetData(name string) *util.String {
	index := -1
	for i, n := range this.Columns {
		if n == name {
			index = i
			break
		}
	}
	if index > -1 {
		return this.Row[index]
	}
	Logger.PError(nil, "ConfigRow.GetData(%s);Not Column.", name)
	return util.NewString("")
}
