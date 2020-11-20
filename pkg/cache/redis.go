package cache

import (
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/spf13/viper"
)

var RedisClient *redis.Pool

func Init() {
	// 从配置文件获取redis的ip以及db
	RedisHost := viper.GetString("redis.addr")
	RedisPassword := viper.GetString("redis.password")
	redisDB := viper.GetInt("redis.db")
	MaxIdle := viper.GetInt("redis.MaxIdle")
	MaxActive := viper.GetInt("redis.MaxActive")
	IdleTimeout := viper.GetInt("redis.IdleTimeout")
	// 建立连接池
	RedisClient = &redis.Pool{
		MaxIdle:     MaxIdle,   /*最大的空闲连接数*/
		MaxActive:   MaxActive, /*最大的激活连接数*/
		IdleTimeout: time.Duration(IdleTimeout) * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", RedisHost, redis.DialPassword(RedisPassword))
			if err != nil {
				return nil, err
			}
			// 选择db
			c.Do("SELECT", redisDB)
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
