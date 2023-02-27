package db_redis

import (
	"errors"
	"fmt"

	"github.com/gomodule/redigo/redis"
	"github.com/sirupsen/logrus"
)

const (
	DefaultExpirationTime int = 1800
)

func RedisCacheString(key string, value string, t int) error {
	conn := GetRec()
	defer CloseConn(conn)
	logrus.Println("set redis: key:", key, "value:", value)
	_, err := conn.Do("SET", key, value)
	if err != nil {
		return err
	}
	_, err = conn.Do("EXPIRE", key, t)
	if err != nil {
		return err
	}
	return nil
}

func RedisGetValue(key string) (string, error) {
	conn := GetRec()
	defer CloseConn(conn)
	value, err := conn.Do("GET", key)
	logrus.Println("查询redis value:", value, "err:", err)
	if err != nil {
		fmt.Println("redis get value err: ", err.Error())
		return "", err
	}
	if value == nil {
		return "", errors.New("value为空")
	}

	return string(value.([]uint8)), err

}

// RedisKeyFlush
// @Description: 刷新过期时间
// @param conn	redis connection
// @param k	键值
// @return error
func RedisKeyFlush(k interface{}) error {
	conn := GetRec()
	defer CloseConn(conn)
	_, err := conn.Do("expire", k, DefaultExpirationTime)
	if err != nil {
		return err
	}
	return err
}

// RedisCheckKey
// @Description: 判断键是否有效
// @param conn	redis connection
// @param k	键值
// @return bool	true:有效 | false:过期
// @return error
func RedisCheckKey(k string) (bool, error) {
	//当key不存在时，返回-2，当key存在但没有设置剩余生存时间时，返回-1。否则，以毫秒为单位，返回key的剩余生存时间
	conn := GetRec()
	defer CloseConn(conn)
	r, err := redis.Int(conn.Do("TTL", k))
	if err != nil {
		return false, err
	}
	if r == -2 {
		return false, nil
	}
	return true, nil
}

// RedisDeleteKey
// @Description: 删除键
// @receiver rec
// @param k
// @return error
func RedisDeleteKey(k string) error {
	//当key不存在时，返回-2，当key存在但没有设置剩余生存时间时，返回-1。否则，以毫秒为单位，返回key的剩余生存时间
	conn := GetRec()
	defer CloseConn(conn)
	_, err := conn.Do("DEL", k)
	if err != nil {
		return err
	}
	return nil
}

// RedisDoKV
// redis操作：action name value
func RedisDoKV(action string, name, value interface{}) error {
	con := GetRec()
	defer CloseConn(con)
	_, err := con.Do(action, name, value)
	if err != nil {
		return err
	}
	return nil
}

// RedisDoHash
// redis操作：action name key value
func RedisDoHash(action string, name, key, value interface{}) error {
	con := GetRec()
	defer CloseConn(con)
	_, err := con.Do(action, name, key, value)
	if err != nil {
		return err
	}
	return nil
}

// RedisDo
// redis操作，切片传值，啥都可以做
// 包装了Do方法，多了错误处理
func RedisDo(action string, values ...interface{}) (reply interface{}, err error) {
	con := GetRec()
	defer CloseConn(con)
	do, err := con.Do(action, values...)
	if err != nil {
		logrus.Errorln("redsi do:", action, values, "is false")
	}
	return do, err
}

// RedisKeyExists
// key存在判断
func RedisKeyExists(key interface{}) bool {
	con := GetRec()
	defer CloseConn(con)
	do, _ := con.Do("EXISTS", key)
	if do.(int64) == 1 {
		return true
	}
	return false
}

func GetAllKV(hash string) (m map[string]string, err error) {
	con := GetRec()
	defer CloseConn(con)
	result, err := redis.Values(con.Do("hgetall", hash))
	if err != nil {
		return nil, err
	}
	m = make(map[string]string, len(result)/2)
	var key string
	for i, v := range result {
		if i&1 == 0 {
			//read key
			key = string(v.([]byte))
		} else {
			//read value
			m[key] = string(v.([]byte))
		}
	}
	return m, nil
}

// CloseConn 包装close方法,多了错误处理
func CloseConn(conn redis.Conn) {
	err := conn.Close()
	if err != nil {
		logrus.Errorln(err.Error())
	}
}
