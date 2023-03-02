/*
type PSubscribeCallback func (pattern, channel, message string)

type PSubscriber struct {
	client redis.PubSubConn
	cbMap map[string]PSubscribeCallback
}


func (c *PSubscriber) PConnect(ip string, port uint16) error {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	//conn, err := redis.Dial("tcp", ip + ":" + strconv.Itoa(int(port)))
	if err != nil {
		log.Println("redis dial failed.")
		return err
	}

	c.client = redis.PubSubConn{conn}
	c.cbMap = make(map[string]PSubscribeCallback)

	go func() {
		for {
			log.Println("wait...")
			switch res := c.client.Receive().(type) {
			case redis.Message:
				pattern := res.Pattern
				channel := string(res.Channel)
				message := string(res.Data)
				if (pattern == "__keyspace@0__:blog*"){
					switch  message {
					case "set":
						// do something
						fmt.Println("set", channel)
					case "del":
						fmt.Println("del", channel)
					case "expire":
						fmt.Println("expire", channel)
					case "expired":
						fmt.Println("expired", channel)
					}
				}
			case redis.Subscription:
				fmt.Printf("%s: %s %d\n", res.Channel, res.Kind, res.Count)
			case error:
				{
				log.Println("error handle...")
					return
				}
				continue
			}
		}
	}()
	return nil
}
func (c *PSubscriber)Psubscribe(channel interface{}, cb PSubscribeCallback) error {
	err := c.client.PSubscribe(channel)
	if err != nil{
		log.Println("redis Subscribe error.")
		return err
	}
	c.cbMap[channel.(string)] = cb
	return nil
}

func TestPubCallback(patter , chann, msg string){
	log.Println( "TestPubCallback patter : " + patter + " channel : ", chann, " message : ", msg)
}

func main() {

	log.Println("===========main start============")
	var psub PSubscriber
	e:=psub.PConnect("127.0.0.1", 6397)
	if e!=nil {
		return
	}
	e=psub.Psubscribe("__keyspace@0__:favoriteUid*", TestPubCallback)
	if e!=nil {
		return
	}
	for{
		time.Sleep(1 * time.Second)
	}
}*/

package main

import (
	//"github.com/go-redis/redis"
	"fmt"
	log "github.com/astaxie/beego/logs"
	"github.com/gomodule/redigo/redis"
	"strconv"
)

type PSubscribeCallback func(pattern, channel, message string)

type PSubscriber struct {
	client redis.PubSubConn
	cbMap  map[string]PSubscribeCallback
}

func (c *PSubscriber) PConnect(ip string, port uint16) {
	conn, err := redis.Dial("tcp", ip+":"+strconv.Itoa(int(port)))
	if err != nil {
		log.Critical("redis dial failed.")
	}

	c.client = redis.PubSubConn{conn}
	c.cbMap = make(map[string]PSubscribeCallback)

	go func() {
		for {
			log.Debug("wait...")
			switch res := c.client.Receive().(type) {
			case redis.Message:
				pattern := res.Pattern
				channel := string(res.Channel)
				message := string(res.Data)
				if pattern == "__keyspace@0__:cool" {
					switch message {
					case "set":
						// do something
						fmt.Println("set", channel)
					case "del":
						fmt.Println("del", channel)
					case "expire":
						fmt.Println("expire", channel)
					case "expired":
						fmt.Println("expired", channel)
					default:
						log.Info("123")
					}
				} else {
					log.Info(pattern)
				}
			case redis.Subscription:
				fmt.Printf("%s: %s %d\n", res.Channel, res.Kind, res.Count)
			case error:
				{
					//log.Error("error handle...")
					log.Debug(c.client.Receive())
					continue
				}

			}
		}
	}()

}
func (c *PSubscriber) Psubscribe(channel interface{}, cb PSubscribeCallback) {
	err := c.client.PSubscribe(channel)
	if err != nil {
		log.Critical("redis Subscribe error.")
	}

	c.cbMap[channel.(string)] = cb
}

func TestPubCallback(patter, chann, msg string) {
	log.Debug("TestPubCallback patter : "+patter+" channel : ", chann, " message : ", msg)
	if msg == "expire" {
		//放入消息队列
		log.Info("123")
	}
}

/*func main() {

	log.Info("===========main start============")

	var psub PSubscriber
	psub.PConnect("43.143.130.52", 6379)
	log.Info("connect完")
	psub.Psubscribe("__keyspace@0__:cool", TestPubCallback)
	log.Info("订阅完")
	for {
		time.Sleep(1 * time.Second)
	}

}*/
