package rpc

import (
	"context"
	"fmt"
	"runedance_douyin/kitex_gen/relation"
	"runedance_douyin/kitex_gen/relation/relationservice"
	constants "runedance_douyin/pkg/consts"
	"runedance_douyin/pkg/errnos"

	"github.com/cloudwego/kitex/client"
)

var relationClient relationservice.Client

func initRelation() {
	c, err := relationservice.NewClient(constants.UserServiceName, client.WithHostPorts("127.0.0.1:9000"))

	if err != nil {
		panic(err)
	}

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
