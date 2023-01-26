package service

import (
	"context"
	"errors"
	"runedance_douyin/cmd/relation/dal/db_mysql"
	"runedance_douyin/kitex_gen/relation"
	constants "runedance_douyin/pkg/consts"
)

type ActionService struct {
	ctx context.Context
}

// todo 需要加*吗
// todo 用Once来保证单例
func GetActionServiceInstance(ctx context.Context) *ActionService {
	return &ActionService{ctx: ctx}
}

// todo 方法名要改
func (a *ActionService) ExecuteAction(req *relation.RelationActionRequest) error {
	// check param
	user_id, err := extract_user_id_from_jwt_token(req.Token)
	if err != nil {
		return err
	}
	to_user_id, err := extract_user_id_from_jwt_token(req.ToUserId)
	if err != nil {
		return err
	}

	if req.ActionType != constants.ActionType_AddRelation && req.ActionType != constants.ActionType_RemoveRelationRelation {
		return errors.New("[action] req param err: ActionType must be 1 or 2")
	}

	//todo 校验用户是否存在
	if req.ActionType == constants.ActionType_AddRelation {
		a.follow(user_id, to_user_id)
	} else if req.ActionType == constants.ActionType_RemoveRelation {
		a.unfollow(user_id, to_user_id)
	}
}
func (a *ActionService) follow(fansId, userId int64) {
	relation := &db_mysql.Relation{
		FansID: fansId,
		UserID: userId,
	}
	db_mysql.CreateRelation(relation)
}

func (a *ActionService) unfollow(fansId, userId int64) {
	relation := &db_mysql.Relation{
		FansID: fansId,
		UserID: userId,
	}
	db_mysql.DeleteRelation(relation)
}
