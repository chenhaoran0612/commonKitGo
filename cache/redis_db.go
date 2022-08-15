package cache

import (
	"errors"
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

const (
	Session    = 0
	Cache      = 1
	BigData    = 2
	TokenRedis = 3
)

var redisPoll map[int]*redis.Pool

type redisConfig struct {
	address  string
	port     string
	password string
}

func RedisConfig(address, password, port string) *redisConfig {
	redisPoll = make(map[int]*redis.Pool)
	return &redisConfig{
		address:  address,
		password: password,
		port:     port,
	}
}
func (r *redisConfig) InitDataBase(dataBase int, maxIdle int, maxActive int, idleTimeout int) error {
	if maxActive <= 0 {
		maxActive = 10
	}
	if maxIdle <= 0 {
		maxIdle = 3
	}
	if idleTimeout < 0 {
		idleTimeout = 10
	}
	redisPoll[dataBase] = &redis.Pool{
		MaxIdle:     maxIdle,   //最大的空闲连接数，表示即使没有redis连接时依然可以保持N个空闲的连接，而不被清除，随时处于待命状态。
		MaxActive:   maxActive, //最大的连接数，表示同时最多有N个连接。0表示不限制。
		IdleTimeout: time.Duration(idleTimeout) * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", r.address, r.port), redis.DialDatabase(dataBase), redis.DialPassword(r.password))
			if err != nil {
				return nil, err
			}
			if r.password != "" {
				if _, err := c.Do("AUTH", r.password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	fmt.Println(fmt.Sprintf("Redis DataBase:%d Init OK", dataBase))
	return nil
}

func NewSession() *session {
	return &session{Session}
}

func NewCache() *cache {
	return &cache{Cache}
}

func getRedisPoll(db int) (*redis.Pool, error) {
	if redis, ok := redisPoll[db]; ok {
		return redis, nil
	} else {
		return nil, errors.New("redis DataBase not exist")
	}
}
