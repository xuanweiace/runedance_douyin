package service

import (
	"context"

	"runedance_douyin/kitex_gen/message"

	"runedance_douyin/cmd/message/dal/db_redis"

	"strconv"

	"encoding/json"

	"time"
)
type UserMessageRecord struct {
	UserId		int64	 `redis:"userId"`
	ToUserId    int64    `redis:"toUserId"`
	Content     string   `redis:"content"`
	CreateTime  string   `redis:"createTime"`
}

type GetMessageChatService struct {
	ctx context.Context
}

func NewGetMessageChatService(ctx context.Context) *GetMessageChatService {
	return &GetMessageChatService{ctx: ctx}
}

func (s *GetMessageChatService) GetMessageChat(ctx context.Context, userId int64, toUserId int64) ([]*message.Message, error){
	var result []*message.Message
	var toResult []*message.Message
	var fromResult []*message.Message

	// "to" message
	recordList1, err := db_redis.GetMessageChatJson(ctx, strconv.Itoa(int(userId)), strconv.Itoa(int(toUserId)))
	if(err != nil){
		return result, err
	}
	// "from" message
	recordList2, err2 := db_redis.GetMessageChatJson(ctx, strconv.Itoa(int(toUserId)), strconv.Itoa(int(userId)))
	if(err2 != nil){
		return result, err2
	}

	// decode json into UserMessageRecord struct
	for _, val := range recordList1 {
		var record UserMessageRecord
		var msg message.Message
		err := json.Unmarshal([]byte(val), &record)
		if(err != nil){
			continue
		}
		// set msg fields
		msg.Id = time.Now().Unix()
		msg.Content = record.Content
		msg.CreateTime = &record.CreateTime
		toResult = append(toResult, &msg)
	}

	for _, val := range recordList2 {
		var record UserMessageRecord
		var msg message.Message
		err := json.Unmarshal([]byte(val), &record)
		if(err != nil){
			continue
		}
		// set msg fields
		msg.Id = time.Now().Unix()
		msg.Content = record.Content
		msg.CreateTime = &record.CreateTime
		fromResult = append(fromResult, &msg)
	}

	// sort message by createTime
	l1 := len(toResult)
	l2 := len(fromResult)
	i := 0
	j := 0
	for(i < l1 && j < l2){
		if(toResult[i].Id > fromResult[j].Id){
			result = append(result, toResult[i])
			i++
			continue
		}
		result = append(result, fromResult[j])
		j++
	}
	for(i < l1){
		result = append(result, toResult[j])
	}
	for(j < l2){
		result = append(result, fromResult[j])
	}

	return result, nil

}