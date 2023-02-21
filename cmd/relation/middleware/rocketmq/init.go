package rmq

import (
	"fmt"
	constants "runedance_douyin/pkg/consts"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/producer"
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
		producer.WithNameServer([]string{constants.NameServerAddr}),
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
	c, err := rocketmq.NewPushConsumer(
		consumer.WithGroupName("relation_consumer_group"),
		consumer.WithNameServer([]string{constants.NameServerAddr}),
	)
	if err != nil {
		panic(err)
	}
	err = c.Subscribe("relation", consumer.MessageSelector{}, consume)
	if err != nil {
		panic(err)
	}
	err = c.Start()
	if err != nil {
		panic(err)
	}
	defer func() {
		err = c.Shutdown()
		if err != nil {

			fmt.Printf("shutdown Consumer error: %s", err.Error())
		}
	}()
	<-(chan interface{})(nil)
	return err
}
