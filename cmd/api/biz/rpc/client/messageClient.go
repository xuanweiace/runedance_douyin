package main

import (
	"context"
	"log"
	"runedance_douyin/pkg/tools"
	"time"
	"runedance_douyin/cmd/api/biz/rpc"
)

func main(){
	rpc.InitMessage()
	token1, _ := tools.GenToken("test_user1", 100)
	token2, _ := tools.GenToken("test_user2", 101)

	for {
		log.Println("user100 send message to user101")
		
		log.Print("token: " + token1)
		err := rpc.MessageAction(context.Background(), 100, 101,  "send test message")
		if(err == nil){
			log.Print("success")
		}
		time.Sleep(time.Second)

		///////////////////////////////////////////
		log.Println("user101 reply to user100")
		
		log.Print("token: " + token2)
		err = rpc.MessageAction(context.Background(), 101, 100, "reply test message")
		if(err == nil){
			log.Print("success")
		}
		time.Sleep(time.Second)

		///////////////////////////////////////////

		log.Println("get message chat between user100 and user101")
		msgList, err2 := rpc.GetMessageChat(context.Background(), 100, 101)
		if(err2 != nil){
			log.Fatal(err2)
		}
		log.Println(msgList)
		time.Sleep(time.Second)
	}

}

