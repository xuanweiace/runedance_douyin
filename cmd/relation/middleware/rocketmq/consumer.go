package rmq

import (
	"context"
	"fmt"
	"runedance_douyin/cmd/relation/dal/db_mysql"
	constants "runedance_douyin/pkg/consts"
	"strconv"
	"strings"

	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

func consume(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
	for _, msg := range msgs {
		fmt.Printf("msg: %+v \n", msg)
		rel := db_mysql.Relation{}
		action_type := decode_relation_msg(msg.Body, &rel)
		if action_type == constants.ActionType_AddRelation {
			db_
		}
	}
	// 消费成功，进行ack确认
	return consumer.ConsumeSuccess, nil
}

func decode_relation_msg(b []byte, rel *db_mysql.Relation) int {
	s := string(b)
	ss := strings.Split(s, " ")
	action_type = strconv.ParseInt(ss[0], 10, 64)

	for i, s := range ss {
		i64, _ :=
			fmt.Println(i, i64)
	}
}
