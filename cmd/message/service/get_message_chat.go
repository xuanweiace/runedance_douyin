package service

import (
	"context"

	"runedance_douyin/kitex_gen/message"

	"runedance_douyin/cmd/message/dal/db_redis"

	"strconv"

)

type GetMessageChatService struct {
	ctx context.Context
}

func NewGetMessageChatService(ctx context.Context) *GetMessageChatService {
	return &GetMessageChatService{ctx: ctx}
}

func (s *GetMessageChatService) GetMessageChat(ctx context.Context, userId int64, toUserId int64) ([]*message.Message, error){
	messageList, err := db_redis.GetMessageChat(ctx, strconv.Itoa(int(userId)), strconv.Itoa(int(toUserId)))
	return messageList, err
}