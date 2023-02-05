package pkg

import "log"

import (
	"github.com/gomodule/redigo/redis"
)

var rec redis.Conn

func GetRec() redis.Conn {
	return rec
}

func InitRedis(url string) {
	con, err := redis.DialURL(url)
	if err != nil {
		log.Fatalln(err)
		return
	}
	rec = con
}
