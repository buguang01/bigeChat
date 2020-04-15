package usermodels

import (
	"time"

	"github.com/buguang01/bige/event"
)

const (
	SAVE_LV0 = 0
	SAVE_LV1 = 1 * time.Second
	SAVE_LV2 = 5 * time.Second
	SAVE_LV3 = 10 * time.Second
)

type UpDataArgs struct {
	DataList []event.ISqlDataModel //要更新的数据
}

func NewUpDataArgs() *UpDataArgs {
	result := new(UpDataArgs)
	result.DataList = make([]event.ISqlDataModel, 0, 50)
	return result
}

func (this *UpDataArgs) Add(sd event.ISqlDataModel) {
	this.DataList = append(this.DataList, sd)
}

func (this *UpDataArgs) UpData() {
	// Service.DBEx.AddMsg(this.DataList...)
}

type DalModel struct {
	event.SqlDataModel
}

func (this *DalModel) GetKeyID() int {
	return this.KeyID
	// if Service.Sconf.SqlConf.InitNum == 0 {
	// 	return this.KeyID
	// }
	// return this.KeyID % Service.Sconf.SqlConf.InitNum
}
