package Events

import (
	"bigeChat/Code/ConstantCode"

	"github.com/buguang01/Logger"
	"github.com/buguang01/bige/event"
)

type LogicListenUser struct {
	Listen   bool
	MemberID int //用户ID
}

//子方法可以按消息来，这样可以把同模块的都放到一起
func (this *LogicListenUser) Hander(jsuser event.JsonMap) (result int) {
	result = ConstantCode.Success
	if this.Listen {
		Logger.SetListenKeyID(this.MemberID)
	} else {
		Logger.RemoveListenKeyID(this.MemberID)
	}
	return result
}
