package service

import (
	"context"
	"log"

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


// return the latest 20 messages or more if in redis
func (s *GetMessageChatService) GetMessageChat(ctx context.Context, userId int64, toUserId int64) ([]*message.Message, error){
	var result_redis []*message.Message

	recordList, err := db_redis.GetMessageChatJson(ctx, userId, toUserId)
	if(err != nil){
		return result_redis, err
	}
	l_redis := len(recordList)												// keyname stores values in redis
	// 	decode json into Message struct
	for _, val := range recordList {
		var msg message.Message
		err := json.Unmarshal([]byte(val), &msg)
		if(err != nil){
			continue
		}
		result_redis = append(result_redis, &msg)
	}
		
	if(l_redis >= 20){
		return result_redis, nil
	}

	l_mysql := 20 - l_redis
	
	// if less than 20, access mysql db to get message chat
	log.Print("retrieve message chat from mysql")
	recordListSQL, err2 := db_mysql.GetMessageChat(ctx, userId, toUserId, l_mysql)
	if(err2 != nil){
		return result_redis, err2
	}

	var result_mysql []*message.Message
	// message ordered from old to new
	for _, val := range recordListSQL {
		msg := message.Message{
			Id : val.Timestamp,
			Content : val.Content,
			CreateTime : val.CreateTime, 
		}
		result_mysql = append(result_mysql, &msg)
	}

	// store mysql message chat to redis
	err3 := db_redis.LoadMessageChat(ctx, userId, toUserId, result_mysql)

	// combine 
	result_mysql = append(result_mysql, result_redis...)

	return result_mysql, err3
}