package service

import (
	"context"

	"runedance_douyin/cmd/message/dal/db_redis"

	"runedance_douyin/cmd/message/middleware/asytask"
)

type MessageActionService struct {
	ctx context.Context
}

func NewMessageActionService(ctx context.Context) *MessageActionService {
	return &MessageActionService{ctx: ctx}
}

func (s *MessageActionService) MessageAction(ctx context.Context, userId int64, toUserId int64, actionType int32, content string) error{
	err := db_redis.HandleMessageSend(ctx, userId,  toUserId, actionType, content)
	if(err != nil){
		return err
	}

	// set timed task to update mysql which will implement after 300s
	err2 := asytask.AddNewTask(ctx, userId, toUserId)
	return err2
}

