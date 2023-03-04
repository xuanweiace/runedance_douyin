package rmq

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

var TOPIC_NAME string

const (
	GAP_TIME = 10 // 每两秒拉取一次
	MSG_NUM  = 5  // 每次拉取5个
)

var prod rocketmq.Producer
var cons rocketmq.PushConsumer

func Init() (err error) {
	err = initProducer()
	if err != nil {
		fmt.Println("[initProducer] err=", err)
		return err
	}
	go func() {
		err = initConsumer()
		if err != nil {
			fmt.Println("[initConsumer] err=", err)
		}
	}()

	return nil
}
func initProducer() error {
	p, err := rocketmq.NewProducer(
		producer.WithNameServer([]string{"127.0.0.1:9876"}),
	)
	if err != nil {
		panic(err)
	}
	err = p.Start()
	defer func() {
		err = p.Shutdown()
		if err != nil {
			fmt.Printf("shutdown producer error: %s", err.Error())
		}
	}()
	prod = p
	return err
}

func initConsumer() error {
	TOPIC_NAME = "relation_topic1"
	// 消费者主动拉取消息
	// not
	c, err := consumer.NewPullConsumer(
		consumer.WithGroupName("pull_consumer_group_test1"), //relation_consumer_group
		consumer.WithNsResolver(primitive.NewPassthroughResolver([]string{"127.0.0.1:9876"})))
	if err != nil {
		panic(err)
	}
	err = c.Start()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer func() {
		err = c.Shutdown()
		if err != nil {
			fmt.Println("Shutdown Pull Consumer error: ", err)
		}
	}()
	queue := primitive.MessageQueue{
		Topic:      TOPIC_NAME,
		BrokerName: "broker-a", // 使用broker的名称
		QueueId:    0,
	}

	offset := int64(0)
	for {
		resp, err := c.PullFrom(context.Background(), &queue, offset, MSG_NUM)
		// resp, err := c.Pull(context.Background(), TOPIC_NAME, consumer.MessageSelector{}, 2)
		fmt.Println("resp=", resp, ", err=", err)
		if err != nil {
			// if err == rocketmq.ErrRequestTimeout {
			// 	fmt.Printf("timeout\n")
			// 	time.Sleep(time.Second)
			// 	continue
			// }
			fmt.Printf("unexpected error: %v\n", err)
			return err
		}
		if resp.Status == primitive.PullFound {
			fmt.Printf("pull message success. nextOffset: %d\n", resp.NextBeginOffset)
			consume_batch_msg(context.Background(), resp.GetMessageExts()...)
			// for _, ext := range resp.GetMessageExts() {
			// 	fmt.Printf("pull msg: %s\n", ext.Body)
			// }
		}
		offset = resp.NextBeginOffset
		time.Sleep(time.Second * GAP_TIME)
	}

}
