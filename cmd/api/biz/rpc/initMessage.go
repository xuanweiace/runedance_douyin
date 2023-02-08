package rpc

import (
	"context"

	"fmt"

	"runedance_douyin/kitex_gen/message"

	"runedance_douyin/kitex_gen/message/messageservice"

	constants "runedance_douyin/pkg/consts"
	"runedance_douyin/pkg/errnos"

	"github.com/cloudwego/kitex/client"
)

var messageClient messageservice.Client

func initMessage() {
	c, err := messageservice.NewClient(constants.MessageServiceName, client.WithHostPorts("127.0.0.1:9000"))

	if err != nil {
		panic(err)
	}
	messageClient = c
}

func MessageAction(ctx context.Context, userId int64, toUserId int64, content string) error {
	req := message.MessageActionRequest{
		UserId: userId,
		ToUserId: toUserId,
		Content: content,
	}
	resp, err := messageClient.MessageAction(ctx, &req)
	if(err != nil){
		return err
	}
	if(resp.StatusCode != errnos.CodeSuccess){
		return fmt.Errorf("[rpc.RelationAction] code=%v, msg=%v", resp.StatusCode, resp.StatusMsg)
	}
	return nil
}


func GetMessageChat(ctx context.Context, userId int64, toUserId int64) ([]*message.Message, error) {
	req := message.GetMessageChatRequest{
		UserId: userId,
		ToUserId: userId,
	}
	resp, err := messageClient.GetMessageChat(ctx, &req)
	if(err != nil){
		return nil, err
	}
	if(resp.StatusCode != errnos.CodeSuccess){
		return nil, fmt.Errorf("[rpc.RelationAction] code=%v, msg=%v", resp.StatusCode, resp.StatusMsg)
	}
	return resp.MsgList, nil
}