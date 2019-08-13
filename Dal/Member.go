package Dal

/*
这是一个例子；
*/
import (
	"time"

	"github.com/buguang01/bige/event"
)

type MemberMD struct {
	MemberID   int       //
	UserName   string    //
	CreateTime time.Time //
	BanTime    time.Time //
	HashKey    string    //
	ServerID   int       //
	ChanID     string    //渠道ID
	OpenID     string    //用户唯一标识
}

func (this *MemberMD) ToJson() event.JsonMap {
	js := make(event.JsonMap)
	js["Name"] = this.UserName
	js["SID"] = this.ServerID
	js["MID"] = this.MemberID

	return js
}
