package events

import (
	"bigeChat/constantcode"
	"bigeChat/msgmodels"

	"github.com/buguang01/Logger"
)

type LogicListen struct {
	Listen   bool
	MemberID int //用户ID
}

//子方法可以按消息来，这样可以把同模块的都放到一起
func (this *LogicListen) Hander() (jsuser msgmodels.H, result int) {
	result = constantcode.Success
	jsuser = make(msgmodels.H)
	if this.Listen {
		Logger.SetListenKeyID(this.MemberID)
	} else {
		Logger.RemoveListenKeyID(this.MemberID)
	}
	return
}
