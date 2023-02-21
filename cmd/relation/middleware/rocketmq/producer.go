package rmq

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/apache/rocketmq-client-go/v2/primitive"
)

func SendActionMsg(action_type, fansId, userId int64) error {
	s := strconv.FormatInt(action_type, 10) + strconv.FormatInt(fansId, 10) + strconv.FormatInt(userId, 10)
	msg := &primitive.Message{
		Topic: "test",
		Body:  []byte(s),
	}
	st := time.Now()
	res, err := prod.SendSync(context.Background(), msg)

	ed := time.Now()
	ed.Second()
	fmt.Println(ed.UnixNano() - st.UnixNano())
	if err != nil {
		panic(err)
	}
	fmt.Println("发送成功,result:", res.String())
	return err
}
