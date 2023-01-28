package service

import (
	"context"

	"strconv"

	"runedance_douyin/cmd/message/dal/db_redis"

)

type MessageActionService struct {
	ctx context.Context
}

func NewMessageActionService(ctx context.Context) *MessageActionService {
	return &MessageActionService{ctx: ctx}
}

func (s *MessageActionService) MessageAction(ctx context.Context, userId int64, toUserId int64, actionType int32, content string) error{
	err := db_redis.HandleMessageSend(ctx, strconv.Itoa(int(userId)),  strconv.Itoa(int(toUserId)), actionType, content)
	return err
}