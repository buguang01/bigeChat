package ChatModels

import (
	"bigeChat/Service"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/buguang01/util"
)

func init() {
	ChatEx = new(ChatManage)
	ChatEx.chatli = make(map[string]*ChatMD)
	ChatEx.maplock = new(sync.RWMutex)
}

const (
	CHAT_WORLD   = "world"   //世界
	CHAT_GROUP   = "group"   //群
	CHAT_PRIVATE = "private" //私有
)

var (
	ChatEx *ChatManage //聊天管理器
)

type ChatManage struct {
	chatli  map[string]*ChatMD
	maplock *sync.RWMutex
}

func (this *ChatManage) GetChat(name string) (result *ChatMD) {
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
			md.UpTime = util.GetCurrTimeSecond()
			Service.MemoryEx.AddListenMsg(md)
			return
		} else {
			if strings.Contains(name, CHAT_WORLD) {
				md = NewChatMD(name, 1)
			} else if strings.Contains(name, CHAT_GROUP) {
				md = NewChatMD(name, 2)
			} else if strings.Contains(name, CHAT_PRIVATE) {
				md = NewChatMD(name, 3)
			} else {
				panic(errors.New(fmt.Sprintf("Not exist chat %s.", name)))
			}
			this.chatli[name] = md
			result = md
			md.UpTime = util.GetCurrTimeSecond()
			Service.MemoryEx.AddListenMsg(md)
		}
	})

	return result
}

func (this *ChatManage) DelChat(name string) (result bool) {
	result = false
	util.UsingWiter(this.maplock, func() {
		md, ok := this.chatli[name]
		if ok {
			if len(md.pusList) > 0 {
				return
			} else if util.GetCurrTimeSecond().Sub(md.UpTime) <= time.Duration(Service.Sconf.MemoryConf.RunTime)*time.Second {
				return
			}
			delete(this.chatli, name)
			result = true
			return
		}
	})
	return result
}
