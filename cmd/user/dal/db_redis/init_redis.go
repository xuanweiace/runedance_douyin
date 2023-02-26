package db_redis

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool

func GetRec() redis.Conn {
	conn := pool.Get()
	// _, err := conn.Do("AUTH", "douyin/123456")
	// if err != nil {
	// 	logrus.Errorln(err)
	// }
	fmt.Println("获取连接成功,conn", conn)
	return conn
}

func InitRedis(url string, pass string) {
	redisPool := &redis.Pool{
		//最大活跃连接数，0代表无限
		MaxActive: 888,
		//最大闲置连接数
		MaxIdle: 20,
		//闲置连接的超时时间
		IdleTimeout: time.Second * 100,
		//定义拨号获得连接的函数
		Dial: func() (redis.Conn, error) {
			return redis.DialURL(url)
		},
	}
	pool = redisPool
}
