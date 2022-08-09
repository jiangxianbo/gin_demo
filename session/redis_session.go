package session

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"sync"
)

type RedisSession struct {
	sessionId string
	// 设置session，可以先放在内存的map中
	client *redis.Client
	// 批量导入redis，提升性能
	sessionMap map[string]interface{}
	// 读写锁
	rwLock sync.RWMutex
	// 记录内存中map是否被操作
	flag int
}

// 用常量去定义状态
const (
	SessionFlagNone = iota
	SessionFlagModify
)

var ctx = context.Background()

// 构造函数
func NewRedisSession(id string, client *redis.Client) *RedisSession {
	s := &RedisSession{
		sessionId:  id,
		client:     client,
		sessionMap: make(map[string]interface{}),
		flag:       SessionFlagNone,
	}
	return s
}

func (r RedisSession) Set(key string, value interface{}) (err error) {
	r.rwLock.Lock()
	defer r.rwLock.Unlock()
	// 设置值
	r.sessionMap[key] = value
	// 标记记录
	r.flag = SessionFlagModify
	return
}

// session存储到redis
func (r RedisSession) Save() (err error) {
	r.rwLock.Lock()
	defer r.rwLock.Unlock()
	// 若标记没变，不操作
	if r.flag != SessionFlagModify {
		return
	}
	// 内存中sessionMap序列化
	data, err := json.Marshal(r.sessionMap)
	if err != nil {
		return
	}
	// 获取redis链接
	err = r.client.Set(ctx, r.sessionId, data, -1).Err()
	if err != nil {
		return
	}
	// 改状态
	r.flag = SessionFlagNone
	return
}

func (r RedisSession) Get(key string) (result interface{}, err error) {
	r.rwLock.Lock()
	defer r.rwLock.Unlock()
	// 先判断内存
	result, ok := r.sessionMap[key]
	if !ok {

	}
	return
}

func (r RedisSession) loadFromRedis() (result interface{}, err error) {
	data, err := r.client.Get(ctx, r.sessionId).Result()
	if err != nil {
		return
	}
	// 去到的东西，反序列化到内存map
	err = json.Unmarshal([]byte(data), r.sessionMap)
	if err != nil {
		return
	}
	return
}

func (r RedisSession) Del(key string) (err error) {
	r.rwLock.Lock()
	defer r.rwLock.Unlock()
	r.flag = SessionFlagModify
	delete(r.sessionMap, key)
	return
}
