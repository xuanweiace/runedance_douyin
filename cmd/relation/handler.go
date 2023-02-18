package main

import (
	"context"
	relation "runedance_douyin/kitex_gen/relation"
	"runedance_douyin/cmd/relation/service"
	"runedance_douyin/pkg/errnos"
)

// RelationServiceImpl implements the last service interface defined in the IDL.
type RelationServiceImpl struct{}

// RelationAction implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationAction(ctx context.Context, req *relation.RelationActionRequest) (resp *relation.RelationActionResponse, err error) {

	return
}

// GetFollowList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFollowList(ctx context.Context, req *relation.GetFollowListRequest) (resp *relation.GetFollowListResponse, err error) {
	// TODO: Your code here...
	return
}

// GetFollowerList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFollowerList(ctx context.Context, req *relation.GetFollowerListRequest) (resp *relation.GetFollowerListResponse, err error) {
	// TODO: Your code here...
	return
}

// GetFriendList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFriendList(ctx context.Context, req *relation.GetFriendListRequest) (resp *relation.GetFriendListResponse, err error) {
	// TODO: Your code here...
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
