package rpc

import (
	"context"

	"fmt"

	"runedance_douyin/kitex_gen/message"

	"runedance_douyin/kitex_gen/message/messageservice"

	constants "runedance_douyin/pkg/consts"
	"runedance_douyin/pkg/errnos"

	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"time"
	"github.com/cloudwego/kitex/pkg/retry"
	"log"
)

var messageClient messageservice.Client

func InitMessage() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	c, err := messageservice.NewClient(constants.MessageServiceName,
		client.WithMuxConnection(1),
		client.WithRPCTimeout(3*time.Second),
		client.WithConnectTimeout(500*time.Microsecond),
		client.WithFailureRetry(retry.NewFailurePolicy()),
		client.WithResolver(r),
	)

	if err != nil {
		panic(err)
	}
	log.Println("successfully start message service")
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
		return fmt.Errorf("[rpc.MessageAction] code=%v, msg=%v", resp.StatusCode, resp.StatusMsg)
	}
	return nil
}


func GetMessageChat(ctx context.Context, userId int64, toUserId int64) ([]*message.Message, error) {
	req := message.GetMessageChatRequest{
		UserId: userId,
		ToUserId: toUserId,
	}
	resp, err := messageClient.GetMessageChat(ctx, &req)
	if(err != nil){
		return nil, err
	}
	if(resp.StatusCode != errnos.CodeSuccess){
		return nil, fmt.Errorf("[rpc.MessageAction] code=%v, msg=%v", resp.StatusCode, resp.StatusMsg)
	}
	return resp.MsgList, nil
}