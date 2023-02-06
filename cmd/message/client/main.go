package main

import (
	"context"
	"time"

	"log"

	"runedance_douyin/kitex_gen/message/messageservice"

	"runedance_douyin/kitex_gen/message"

	"github.com/cloudwego/kitex/client"

	"runedance_douyin/pkg/tools"
)

func main() {
	client, err := messageservice.NewClient("test_client", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}

	for{
		token, _ := tools.GenToken("test_user", 102)
		token2, _ := tools.GenToken("test_toUser", 103)
		// test MessageAction
		// send message
		log.Println("user100 send message to user101")
		msgActionReq := &message.MessageActionRequest{
			Token: token,
			ToUserId: 103,
			ActionType: 1,
			Content: "send test message",
		}
		// log.Println(msgActionReq)
		msgActionResp, err := client.MessageAction(context.Background(), msgActionReq)
		if(err != nil){
			log.Fatal(err)
		}
		log.Println(msgActionResp)
		time.Sleep(time.Second)

		// reply
		log.Println("user101 reply to user100")
		msgActionReq2 := &message.MessageActionRequest{
			Token: token2,
			ToUserId: 102,
			ActionType: 1,
			Content: "reply test message",
		}
		msgActionResp2, err2 := client.MessageAction(context.Background(), msgActionReq2)
		if(err2 != nil){
			log.Fatal(err)
		}
		log.Println(msgActionResp2)
		time.Sleep(time.Second)

		//=======================================================

		// test GetMessageChat
		log.Println("get message chat between user100 and user101")
		msgChatReq := &message.GetMessageChatRequest{
			Token: token,
			ToUserId: 103,
		}
		log.Println(msgChatReq)
		resp2, err2 := client.GetMessageChat(context.Background(), msgChatReq)
		if(err2 != nil){
			log.Fatal(err2)
		}
		log.Println(resp2)
		time.Sleep(time.Second)
	}
}