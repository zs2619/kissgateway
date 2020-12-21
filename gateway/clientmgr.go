package gateway

import (
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/zs2619/kissnet-go"
)

type UserClinetMgr struct {
	userClinetIDMap map[string]*UserClinet
	connMap         map[kissnet.IConnection]*UserClinet
	num             int64
	mutex           sync.RWMutex
}

var userClinetMgr *UserClinetMgr = &UserClinetMgr{
	userClinetIDMap: make(map[string]*UserClinet),
	connMap:         make(map[kissnet.IConnection]*UserClinet),
	num:             int64(0),
}

func (this *UserClinetMgr) GetUserClientByUserID(UserID string) *UserClinet {
	this.mutex.RLock()
	defer this.mutex.RUnlock()
	if v, ok := this.userClinetIDMap[UserID]; ok {
		return v
	}
	return nil
}
func (this *UserClinetMgr) GetUserClientByConn(conn kissnet.IConnection) *UserClinet {
	this.mutex.RLock()
	defer this.mutex.RUnlock()
	if v, ok := this.connMap[conn]; ok {
		return v
	}
	return nil
}

func (this *UserClinetMgr) AddUserClient(userClientID string, conn kissnet.IConnection) *UserClinet {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	u := &UserClinet{
		Conn:         conn,
		UserClientID: userClientID,
	}
	this.connMap[u.Conn] = u
	this.userClinetIDMap[u.UserClientID] = u
	this.num++
	u.Init()
	return u
}

func (this *UserClinetMgr) DelUserClient(conn kissnet.IConnection) {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	u, ok := this.connMap[conn]
	if !ok {
		return
	}
	this.num--
	u.Dispose()

	delete(this.connMap, conn)
	delete(this.userClinetIDMap, u.UserClientID)
}
func (this *UserClinetMgr) Close() {
	logrus.Info("UserChatMgr Close")
	for k := range this.connMap {
		k.Close()
	}
}

func (this *UserClinetMgr) GetUserNum() int64 {
	this.mutex.RLock()
	defer this.mutex.RUnlock()
	return this.num
}
