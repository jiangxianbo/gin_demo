package session

import (
	"errors"
	"sync"

	uuid "github.com/satori/go.uuid"
)

// 定义对象
type MemorySessionMgr struct {
	sessionMap map[string]Session
	rwLock     sync.RWMutex
}

// 构造函数
func NewMemorySessionMgr() *MemorySessionMgr {
	sr := &MemorySessionMgr{
		sessionMap: make(map[string]Session),
	}
	return sr
}

func (m *MemorySessionMgr) Init(addr string, options ...string) (err error) {
	return
}

func (m *MemorySessionMgr) CreateSession() (session Session, err error) {
	m.rwLock.Lock()
	defer m.rwLock.Unlock()
	// go get github.com/satori/go.uuid
	// 用uuid作为sessionId
	id := uuid.NewV4()
	// 转成string
	sessionId := id.String()
	// 创建session
	session = NewMemorySession(sessionId)
	// 加入到大map
	m.sessionMap[sessionId] = session
	return
}

func (m MemorySessionMgr) Get(sessionId string) (session Session, err error) {
	m.rwLock.Lock()
	defer m.rwLock.Unlock()
	session, ok := m.sessionMap[sessionId]
	if !ok {
		err = errors.New("session not exists")
		return
	}
	return
}
