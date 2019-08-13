package ConstantCode

const (
	NsqRequestFailure = -6 //请求失败
	NotLogic          = -5 //没有逻辑处理它
	NotAction         = -4 //没有事件处理这个消息
	NotResult         = -3 //不回复
	LOGIC_ERROR       = -2 //没有逻辑处理
	Timeout           = -1 //超时
	Success           = 0  //成功
	User_STATUS_NOT   = 1  //用户状态错误
	Parm_Error        = 2  //参数错误
	User_DB_Error     = 3  //数据库错误
)

//用户
const (
	Player_Not_Exist  = 3000 //用户不存在
	Player_Hash_Error = 3001 //hash不对
)
