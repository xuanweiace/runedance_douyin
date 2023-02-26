package service

import (
	"context"
	"log"
	"math"

	"runedance_douyin/kitex_gen/message"

	"runedance_douyin/cmd/message/dal/db_redis"

	"runedance_douyin/cmd/message/dal/db_mysql"

	"encoding/json"

	"time"

	"strconv"
)

type GetMessageChatService struct {
	ctx context.Context
}

func NewGetMessageChatService(ctx context.Context) *GetMessageChatService {
	return &GetMessageChatService{ctx: ctx}
}


// return the latest 20 messages or more if in redis
func (s *GetMessageChatService) GetMessageChat(ctx context.Context, userId int64, toUserId int64) ([]*message.Message, error){
	db_redis.InitCluster()
	var result_redis []*message.Message
	recordList, err := db_redis.GetMessageChatJson(ctx, userId, toUserId)
	if(err != nil){
		return result_redis, err
	}
	l_redis := len(recordList)												// keyname stores values in redis

	// 	decode json into Message struct
	for _, val := range recordList {
		
		var temp db_mysql.MessageRecord
		err := json.Unmarshal([]byte(val), &temp)
		if(err != nil){
			continue
		}
		msg := message.Message{
			Id: temp.Timestamp,
			Content: temp.Content,
			CreateTime: time.Unix(temp.Timestamp, 0).Format(time.UnixDate),
		}
		result_redis = append(result_redis, &msg)
	}
		
	if(l_redis >= 20){
		return result_redis, nil
	}
	var earliest_redis int64
	earliest_redis = math.MaxInt64
	if(l_redis > 0){
		earliest_redis = result_redis[0].Id			// the earlest message in redis
		log.Print("the earlest is " + strconv.FormatInt(earliest_redis, 10))
	}
	
	l_mysql := 20 - l_redis
	
	// if less than 20, access mysql db to get message chat
	log.Print("retrieve message chat from mysql")
	recordListSQL, err2 := db_mysql.GetMessageChat(ctx, userId, toUserId, l_mysql, earliest_redis)
	if(err2 != nil){
		return result_redis, err2
	}

	var result_mysql []*message.Message


	// message ordered from old to new
	for i := len(recordListSQL) - 1; i >= 0; i --{
		val := recordListSQL[i]
		timestamp := val.Timestamp
		msg := message.Message{
			Id : timestamp,
			Content : val.Content,
			CreateTime : time.Unix(timestamp, 0).Format(time.UnixDate),   // convert to readable time string
		}
		result_mysql = append(result_mysql, &msg)
	}

	// store mysql message chat to redis
	err3 := db_redis.LoadMessageChat(ctx, userId, toUserId, recordListSQL)

	// combine 
	result_mysql = append(result_mysql, result_redis...)

	return result_mysql, err3
}