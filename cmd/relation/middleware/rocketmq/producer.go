package rmq

import (
	"context"
	"fmt"
	"strconv"

	"github.com/apache/rocketmq-client-go/v2/primitive"
)

func SendActionMsg(action_type, fansId, userId int64) error {
	s := strconv.FormatInt(action_type, 10) + strconv.FormatInt(fansId, 10) + strconv.FormatInt(userId, 10)
	msg := &primitive.Message{
		Topic: TOPIC_NAME,
		Body:  []byte(s),
	}
	res, err := prod.SendSync(context.Background(), msg)
	if err != nil {
		panic(err)
	}
	fmt.Println("[SendActionMsg],result:", res.String())
	return err
}
