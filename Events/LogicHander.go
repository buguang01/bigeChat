package Events

import (
	"bigeChat/Code/ConstantCode"

	"github.com/buguang01/Logger"
	"github.com/buguang01/bige/threads"
)

//userLogicRun 用户逻辑运行,这应该还要加一个参数，也就是你的用户ID或用户对象
//用来在发生异常时，调用下面的那个方法
func UserLogicRun(f func() int) int {
	result := ConstantCode.Success
	threads.Try(func() {
		result = f()
	}, func(err interface{}) {
		Logger.PFatal(err)
		result = ConstantCode.User_DB_Error
		// UserDbErrorHander(user)
		/*
			这里是如果传入的方法运行出错时，会给自己发一个同一的消息
			比如发一个重新加载用户数据的消息，用来清除用户内存的脏数据；
		*/

	}, nil)
	return result
}
