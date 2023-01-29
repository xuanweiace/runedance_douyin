package service

import (
	"context"

	"runedance_douyin/kitex_gen/message"

	"runedance_douyin/cmd/message/dal/db_redis"

	"encoding/json"

)

type GetMessageChatService struct {
	ctx context.Context
}

func NewGetMessageChatService(ctx context.Context) *GetMessageChatService {
	return &GetMessageChatService{ctx: ctx}
}

func (s *GetMessageChatService) GetMessageChat(ctx context.Context, userId int64, toUserId int64) ([]*message.Message, error){
	var result []*message.Message

	recordList, err := db_redis.GetMessageChatJson(ctx, userId, toUserId)
	if(err != nil){
		return result, err
	}

	// decode json into Message struct
	for _, val := range recordList {
		var msg message.Message
		err := json.Unmarshal([]byte(val), &msg)
		if(err != nil){
			continue
		}
		result = append(result, &msg)
	}

	return result, nil
}