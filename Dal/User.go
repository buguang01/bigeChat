package Dal

type UserModel struct {
	UserInfo string          //用户信息
	ChatLi   map[string]bool //监听的频道
}

func NewUserModel() *UserModel {
	md := new(UserModel)
	md.ChatLi = make(map[string]bool)
	return md
}
