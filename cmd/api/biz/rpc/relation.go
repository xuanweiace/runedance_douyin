package rpc

import (
	"context"
	"fmt"
	"log"
	"runedance_douyin/kitex_gen/relation"
	"runedance_douyin/kitex_gen/relation/relationservice"
	constants "runedance_douyin/pkg/consts"
	"runedance_douyin/pkg/errnos"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var relationClient relationservice.Client

func initRelation() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	c, err := relationservice.NewClient(constants.RelationServiceName,
		client.WithMuxConnection(1),
		client.WithRPCTimeout(3*time.Second),
		client.WithConnectTimeout(500*time.Microsecond),
		client.WithFailureRetry(retry.NewFailurePolicy()),
		client.WithResolver(r),
	)

	if err != nil {
		panic(err)
	}
	log.Println("[api服务 initRelation] relation服务启动成功")
	relationClient = c
}

func RelationAction(ctx context.Context, from_id, to_id int64, action_type int32) error {
	req := relation.RelationActionRequest{
		FromUserId: from_id,
		ToUserId:   to_id,
		ActionType: action_type,
	}
	resp, err := relationClient.RelationAction(ctx, &req)
	if err != nil {
		return err
	}
	if resp.StatusCode != errnos.CodeSuccess {
		return fmt.Errorf("[rpc.RelationAction] code=%v, msg=%v", resp.StatusCode, resp.StatusMsg)
	}
	return nil

}

func GetFollowList(ctx context.Context, user_id int64) ([]*relation.User, error) {
	req := relation.GetFollowListRequest{
		UserId: user_id,
	}
	resp, err := relationClient.GetFollowList(context.Background(), &req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != errnos.CodeSuccess {
		return nil, fmt.Errorf("[rpc.GetFollowList] code=%v, msg=%v", resp.StatusCode, resp.StatusMsg)
	}

	return resp.UserList, nil

}

func GetFollowerList(ctx context.Context, user_id int64) ([]*relation.User, error) {
	req := relation.GetFollowerListRequest{
		UserId: user_id,
	}
	resp, err := relationClient.GetFollowerList(ctx, &req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != errnos.CodeSuccess {
		return nil, fmt.Errorf("[rpc.GetFollowerList] code=%v, msg=%v", resp.StatusCode, resp.StatusMsg)
	}

	return resp.UserList, nil

}

func GetFriendList(ctx context.Context, user_id int64) ([]*relation.User, error) {
	req := relation.GetFriendListRequest{
		UserId: user_id,
	}
	resp, err := relationClient.GetFriendList(ctx, &req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != errnos.CodeSuccess {
		return nil, fmt.Errorf("[rpc.GetFriendList] code=%v, msg=%v", resp.StatusCode, resp.StatusMsg)
	}

	return resp.UserList, nil

}
