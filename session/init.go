package session

import "fmt"

// 中间件让用户去选择

var (
	sessionMgr SessionMgr
)

func Init(provider string, addr string, option ...string) (err error) {
	switch provider {
	case "memory":
		sessionMgr = NewMemorySessionMgr()
	case "redis":
		sessionMgr = NewRedisSessionMgr()
	default:
		fmt.Errorf("不支持")
		return
	}
	sessionMgr.Init(addr, option...)
	return
}
