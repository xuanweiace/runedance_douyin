package service

import (
	"context"

	"runedance_douyin/kitex_gen/message"

	"runedance_douyin/cmd/message/dal/db_redis"

	"runedance_douyin/cmd/message/dal/db_mysql"

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

	if(len(recordList) != 0){													// keyname stores values in redis
	// 	decode json into Message struct
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

	// access mysql db to get message chat
	recordListSQL, err2 := db_mysql.GetMessageChat(ctx, userId, toUserId)
	if(err2 != nil){
		return result, err2
	}
	for _, val := range recordListSQL {
		msg := message.Message{
			Id : val.Timestamp,
			Content : val.Content,
			CreateTime : val.CreateTime, 
		}
		result = append(result, &msg)
	}

	// store message chat to redis
	err3 := db_redis.LoadMessageChat(ctx, userId, toUserId, result)
	if(err3 != nil){
		return result, err3
	}
	return result, nil
}