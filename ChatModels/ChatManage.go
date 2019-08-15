package ChatModels

import (
	"bigeChat/Service"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/buguang01/bige/event"
	"github.com/buguang01/util"
)

func init() {
	ChatEx = new(ChatManage)
	ChatEx.chatli = make(map[string]IChatMD)
	ChatEx.maplock = new(sync.RWMutex)
}

/*
频道名字的规则是前缀+ID
例：
群频道：
101群频道
name:"group,101"

私聊频道为：
100001与100002聊天
name:"private,100001,100002"
私聊时，较小的用户ID在前面
*/
const (
	CHAT_WORLD   = "world"   //世界
	CHAT_GROUP   = "group"   //群
	CHAT_PRIVATE = "private" //私有
	CHAT_SYSTEM  = "system"  //系统
)

var (
	ChatEx *ChatManage //聊天管理器
)

type ChatManage struct {
	chatli   map[string]IChatMD
	playerli map[int]*PlayerChatMD
	maplock  *sync.RWMutex
}

//获取频道
func (this *ChatManage) GetChat(name string) (result IChatMD) {
	result = nil
	util.UsingRead(this.maplock, func() {
		if md, ok := this.chatli[name]; ok {
			result = md
			Service.MemoryEx.AddListenMsg(md)
		}
	})
	if result != nil {
		return result
	}
	util.UsingWiter(this.maplock, func() {
		if md, ok := this.chatli[name]; ok {
			result = md
			md.SetUpTime(util.GetCurrTimeSecond())
			Service.MemoryEx.AddListenMsg(md)
			return
		} else {
			if strings.Contains(name, CHAT_WORLD) {
				md = NewChatMD(name, 1)
			} else if strings.Contains(name, CHAT_GROUP) {
				md = NewChatMD(name, 2)
			} else if strings.Contains(name, CHAT_PRIVATE) {
				md = NewChatMD(name, 3)
				arr := strings.Split(name, ",")
				for i := 1; i <= 2; i++ {
					pmd := this.getPlayerLi(util.NewString(arr[i]).ToIntV())
					pmd.ChatAdd(md)
				}

			} else if strings.Contains(name, CHAT_SYSTEM) {
				md = NewChatMD(name, 4)
			} else {
				panic(errors.New(fmt.Sprintf("Not exist chat %s.", name)))
			}
			this.chatli[name] = md
			result = md
			md.SetUpTime(util.GetCurrTimeSecond())
			Service.MemoryEx.AddListenMsg(md)
		}
	})

	return result
}

//删频道
func (this *ChatManage) DelChat(name string) (result bool) {
	result = false
	util.UsingWiter(this.maplock, func() {
		md, ok := this.chatli[name]
		if ok {
			if md.GetLenPusList() > 0 {
				return
			} else if util.GetCurrTimeSecond().Sub(md.GetUpTime()) <= time.Duration(Service.Sconf.MemoryConf.RunTime)*time.Second {
				return
			}
			delete(this.chatli, name)
			md.SetUpTime(util.GetMinDateTime())
			if md.GetTypeChat() == 3 { //私聊
				arr := strings.Split(name, ",")
				for i := 1; i <= 2; i++ {
					mid := util.NewString(arr[i]).ToIntV()
					pmd := this.getPlayerLi(mid)
					pmd.ChatDel(md)
					if pmd.Wsmd == nil && len(pmd.Chatli) == 0 {
						delete(this.playerli, mid)
					}
				}
			}
			result = true
			return
		}
	})
	return result
}

//注册用户的私聊频道
func (this *ChatManage) RegisterPlayer(mid int, wsmd *event.WebSocketModel) {
	util.UsingWiter(this.maplock, func() {
		pmd := this.getPlayerLi(mid)
		pmd.Wsmd = wsmd
		for _, md := range pmd.Chatli {
			md.PusAdd(wsmd)
		}
	})
}

//用户离线
func (this *ChatManage) OfflinePlayer(mid int, wsmd *event.WebSocketModel) {
	util.UsingWiter(this.maplock, func() {
		pmd := this.getPlayerLi(mid)
		if pmd.Wsmd == wsmd {
			pmd.Wsmd = nil
		}
		for _, md := range pmd.Chatli {
			md.PusDel(wsmd)
		}
		if pmd.Wsmd == nil && len(pmd.Chatli) == 0 {
			delete(this.playerli, mid)
		}
	})
}

//用户私聊频道管理器
func (this *ChatManage) getPlayerLi(mid int) (result *PlayerChatMD) {
	result, ok := this.playerli[mid]
	if !ok {
		result = NewPlayerChatMD(mid)
		this.playerli[mid] = result
	}
	return result
}
