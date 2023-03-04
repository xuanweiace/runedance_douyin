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

func consume_batch_msg(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
	var err error
	do_map := make(map[string]struct{}, 10)
	undo_map := make(map[string]struct{}, 10)

	for _, msg := range msgs {
		fmt.Printf("msg: %+v \n", msg)
		rel := db_mysql.Relation{}
		action_type := decode_relation_msg(msg.Body, &rel)
		if action_type == constants.ActionType_AddRelation {
			do_map[gen_key(&rel)] = struct{}{}
			delete(undo_map, gen_key(&rel))
		} else if action_type == constants.ActionType_RemoveRelation {
			undo_map[gen_key(&rel)] = struct{}{}
			delete(do_map, gen_key(&rel))
		}
	}
	// todo 批量消费 写入mysql
	do_rels := make([]*db_mysql.Relation, 0)
	for k, _ := range do_map {
		s := strings.Split(k, "-")
		fid, _ := strconv.ParseInt(s[0], 10, 64)
		uid, _ := strconv.ParseInt(s[1], 10, 64)
		do_rels = append(do_rels, &db_mysql.Relation{FansID: fid, UserID: uid})
	}
	err = db_mysql.BatchCreate(do_rels)

	undo_rels := make([]*db_mysql.Relation, 0)
	for k, _ := range undo_map {
		s := strings.Split(k, "-")
		fid, _ := strconv.ParseInt(s[0], 10, 64)
		uid, _ := strconv.ParseInt(s[1], 10, 64)
		undo_rels = append(undo_rels, &db_mysql.Relation{FansID: fid, UserID: uid})
	}

	err = db_mysql.BatchDelete(undo_rels)
	if err == nil {
		// 消费成功，进行ack确认
		return consumer.ConsumeSuccess, nil
	} else {
		return consumer.ConsumeRetryLater, err
	}

}

func decode_relation_msg(b []byte, rel *db_mysql.Relation) int {
	s := string(b)
	ss := strings.Split(s, " ")
	if len(ss) != 3 {
		return -1
	}
	action_type, _ := strconv.ParseInt(ss[0], 10, 64)
	rel.FansID, _ = strconv.ParseInt(ss[1], 10, 64)
	rel.UserID, _ = strconv.ParseInt(ss[2], 10, 64)

	return int(action_type)
}

func gen_key(rel *db_mysql.Relation) string {
	strconv.FormatInt(rel.FansID, 10)
	return strconv.FormatInt(rel.FansID, 10) + "-" + strconv.FormatInt(rel.UserID, 10)
}
