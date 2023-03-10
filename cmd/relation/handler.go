package main

import (
	"context"
	"runedance_douyin/cmd/relation/service"
	relation "runedance_douyin/kitex_gen/relation"
	"runedance_douyin/pkg/errnos"

	"github.com/sirupsen/logrus"
)

// RelationServiceImpl implements the last service interface defined in the IDL.
type RelationServiceImpl struct{}

// RelationAction implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationAction(ctx context.Context, req *relation.RelationActionRequest) (resp *relation.RelationActionResponse, err error) {
	logrus.Info("RelationAction")
	//幂等接口
	resp = relation.NewRelationActionResponse()
	err = service.GetActionServiceInstance(ctx).ExecuteAction(req)
	if err != nil {
		resp.StatusCode = errnos.CodeServiceErr
		er := err.Error()
		resp.StatusMsg = &er
	} else {
		resp.StatusCode = 0
		resp.StatusMsg = nil
	}
	return
}

// GetFollowList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFollowList(ctx context.Context, req *relation.GetFollowListRequest) (resp *relation.GetFollowListResponse, err error) {
	logrus.Info("GetFollowList")
	resp = relation.NewGetFollowListResponse()
	userList, err := service.GetQueryServiceInstance(ctx).GetFollowList(req)
	if err != nil {
		resp.StatusCode = errnos.CodeServiceErr
		er := err.Error()
		resp.StatusMsg = &er
	} else {
		resp.StatusCode = 0
		resp.StatusMsg = nil
		resp.UserList = userList
	}
	return
}

// GetFollowerList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFollowerList(ctx context.Context, req *relation.GetFollowerListRequest) (resp *relation.GetFollowerListResponse, err error) {
	logrus.Info("GetFollowerList")
	resp = relation.NewGetFollowerListResponse()
	followerList, err := service.GetQueryServiceInstance(ctx).GetFollowerList(req)
	if err != nil {
		resp.StatusCode = errnos.CodeServiceErr
		er := err.Error()
		resp.StatusMsg = &er
	} else {
		resp.StatusCode = 0
		resp.StatusMsg = nil
		resp.UserList = followerList
	}
	return
}

// GetFriendList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFriendList(ctx context.Context, req *relation.GetFriendListRequest) (resp *relation.GetFriendListResponse, err error) {
	logrus.Info("test")
	resp = relation.NewGetFriendListResponse()
	friendList, err := service.GetQueryServiceInstance(ctx).GetFriendList(req)
	if err != nil {
		resp.StatusCode = errnos.CodeServiceErr
		er := err.Error()
		resp.StatusMsg = &er
	} else {
		resp.StatusCode = 0
		resp.StatusMsg = nil

		resp.UserList = friendList
	}
	return
}

// ExistRelation implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) ExistRelation(ctx context.Context, req *relation.ExistRelationRequest) (resp *relation.ExistRelationResponse, err error) {
	resp = relation.NewExistRelationResponse()
	existed, err := service.GetQueryServiceInstance(ctx).ExistRelation(req)
	if err != nil {
		resp.StatusCode = errnos.CodeServiceErr
		er := err.Error()
		resp.StatusMsg = &er
	} else {
		resp.StatusCode = 0
		resp.StatusMsg = nil
		resp.Existed = existed
	}

	return
}
