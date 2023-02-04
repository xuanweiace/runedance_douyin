package service

import (
	"context"
	"errors"
	"runedance_douyin/cmd/relation/dal/db_mysql"
	"runedance_douyin/cmd/relation/rpc"
	"runedance_douyin/kitex_gen/relation"
	constants "runedance_douyin/pkg/consts"
	"sync"
)

type ActionService struct {
	ctx context.Context
}

var actionService ActionService
var actionServiceOnce sync.Once

func GetActionServiceInstance(ctx context.Context) *ActionService {
	actionServiceOnce.Do(func() {
		actionService = ActionService{ctx: ctx}
	})
	return &actionService
}

// todo 方法名要改
func (a *ActionService) ExecuteAction(req *relation.RelationActionRequest) error {
	// check param
	user_id := req.FromUserId
	to_user_id := req.ToUserId

	if req.ActionType != constants.ActionType_AddRelation && req.ActionType != constants.ActionType_RemoveRelation {
		return errors.New("[action] req param err: ActionType must be 1 or 2")
	}
	// todo 这个逻辑？
	if user_id == to_user_id {
		return errors.New("[action] req param err: user_id can not be same")
	}

	// 校验用户是否存在
	if _, err := rpc.GetUser(user_id); err != nil {
		return err
	}
	if req.ActionType == constants.ActionType_AddRelation {
		if err := a.follow(user_id, to_user_id); err != nil {
			return err
		}
	} else if req.ActionType == constants.ActionType_RemoveRelation {
		if err := a.unfollow(user_id, to_user_id); err != nil {
			return err
		}
	}
	return nil
}
func (a *ActionService) follow(fansId, userId int64) (err error) {
	rela := &db_mysql.Relation{
		FansID: fansId,
		UserID: userId,
	}
	err = db_mysql.CreateRelation(rela)
	return
}

func (a *ActionService) unfollow(fansId, userId int64) (err error) {
	rela := &db_mysql.Relation{
		FansID: fansId,
		UserID: userId,
	}
	err = db_mysql.DeleteRelation(rela)
	return
}
