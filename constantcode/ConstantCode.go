package constantcode

const (
	NsqRequestFailure   = -6 //请求失败
	NotLogic            = -5 //没有逻辑处理它
	NotAction           = -4 //没有事件处理这个消息
	NotResult           = -3 //不回复
	LOGIC_Unknown_Error = -2 //逻辑处理错误
	Timeout             = -1 //超时
	Success             = 0  //成功
	User_STATUS_NOT     = 1  //用户状态错误
	Parm_Error          = 2  //参数错误
	User_DB_Error       = 3  //数据库错误
)
