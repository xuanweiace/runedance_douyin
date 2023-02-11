package rpc

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	"runedance_douyin/kitex_gen/relation"
	"runedance_douyin/kitex_gen/relation/relationservice"
	constants "runedance_douyin/pkg/consts"
	"time"
)

var relationClient relationservice.Client

func initRelation() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	c, err := relationservice.NewClient(constants.RelationServiceName,
		client.WithMuxConnection(1),
		client.WithRPCTimeout(3*time.Second),
		client.WithConnectTimeout(500*time.Microsecond), // 50ms会超时
		client.WithFailureRetry(retry.NewFailurePolicy()),
		client.WithResolver(r),
	)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	relationClient = c
	return
}

// ExistRelation implements the RelationServiceImpl interface.
func ExistRelation(fromUserId, ToUserId int64) (bool, error) {
	req := &relation.ExistRelationRequest{
		fromUserId,
		ToUserId,
	}
	resp, err := relationClient.ExistRelation(context.Background(), req)
	return resp.Existed, err
}
