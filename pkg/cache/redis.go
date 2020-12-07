package cache

import (
	"apiapi/pkg/config"
	"time"

	"github.com/gomodule/redigo/redis"
)

var RedisClient *redis.Pool

func Init() {
	// 从配置文件获取redis的ip以及db
	cfg := config.C.Redis
	// 建立连接池
	RedisClient = &redis.Pool{
		MaxIdle:     cfg.MaxIdle,   /*最大的空闲连接数*/
		MaxActive:   cfg.MaxActive, /*最大的激活连接数*/
		IdleTimeout: time.Duration(cfg.IdleTimeout) * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", cfg.Addr, redis.DialPassword(cfg.Password))
			if err != nil {
				return nil, err
			}
			// 选择db
			c.Do("SELECT", cfg.DB)
			return c, nil
		},
	}
}

//用法
// rc := RedisClient.Get()
// defer rc.Close()
// 操作
//_, err := rc.Do("SET", "key", "value")
//value, err := redis.String(rc.Do("lpop", "kuaishou"))
