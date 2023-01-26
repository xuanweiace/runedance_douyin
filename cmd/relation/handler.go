package main

import (
	"context"
	"runedance_douyin/cmd/relation/service"
	relation "runedance_douyin/kitex_gen/relation"
	"runedance_douyin/pkg/errnos"
)

// RelationServiceImpl implements the last service interface defined in the IDL.
type RelationServiceImpl struct{}

// RelationAction implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationAction(ctx context.Context, req *relation.RelationActionRequest) (resp *relation.RelationActionResponse, err error) {
	resp = relation.NewRelationActionResponse()
	err = service.GetActionServiceInstance(ctx).ExecuteAction(req)
	if err != nil {
		resp.BaseResp.StatusCode = errnos.CodeServiceErr
		er := err.Error() //todo 只能这样写？
		resp.BaseResp.StatusMsg = &er
	} else {
		resp.BaseResp = &relation.BaseResponse{
			StatusCode: 0,
			StatusMsg:  nil,
		}
	}
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
