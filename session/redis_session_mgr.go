package session

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	uuid "github.com/satori/go.uuid"
)

type RedisSessionMgr struct {
	addr       string
	passwd     string
	redisDB    *redis.Client
	rwLock     sync.RWMutex
	sessionMap map[string]Session
}

func (r *RedisSessionMgr) Init(addr string, options ...string) (err error) {
	// 若有参数
	if len(options) > 0 {
		r.passwd = options[0]
	}
	// 创建链接
	r.redisDB, _ = myRedisClient(addr, r.passwd)
	r.addr = addr
	return
}

func myRedisClient(addr string, password string) (*redis.Client, error) {
	redisDB := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
		PoolSize: 100, // 连接池大小
	})
	fmt.Printf("Connecting Redis : %v\n", addr)

	ctxt, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	res, err := redisDB.Ping(ctxt).Result()
	if err != nil {
		fmt.Printf("redis connect failed, err%s\n", err)
		return nil, err
	}
	fmt.Printf("Connect Successful! Ping => %v\n", res)
	return redisDB, nil
}

func (r *RedisSessionMgr) CreateSession() (session Session, err error) {
	r.rwLock.Lock()
	defer r.rwLock.Unlock()
	id := uuid.NewV4()
	sessionId := id.String()
	r.sessionMap[sessionId] = NewRedisSession(sessionId, r.redisDB)
	return
}

func (r *RedisSessionMgr) Get(sessionId string) (session Session, err error) {
	r.rwLock.Lock()
	defer r.rwLock.Unlock()
	session, ok := r.sessionMap[sessionId]
	if !ok {
		errors.New("session not exists")
	}
	return
}

func NewRedisSessionMgr() SessionMgr {
	sr := &RedisSessionMgr{
		sessionMap: make(map[string]Session, 32),
	}
	return sr
}
