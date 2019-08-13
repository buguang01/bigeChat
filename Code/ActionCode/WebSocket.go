package ActionCode

//定义来自WebSocket的消息
const (
	// Ws_User_Test = 1001 //例子消息

	Ws_Chat_Notice    = 5001 //聊天通知
	Ws_Chat_Register  = 5002 //注册连接
	Ws_Chat_Pus       = 5003 //监听聊天频道
	Ws_Chat_CancelPus = 5004 //取消监听聊天频道
	Ws_Chat_DiDa      = 5005 //心跳
	Ws_Chat_SendMsg   = 5010 //发聊天信息

)
