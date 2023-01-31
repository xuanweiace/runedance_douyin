package middleware

import (
	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	conn  *amqp.Connection
	mqurl string
}

// 常见mq ActiveMQ、RabbitMQ、RocketMQ、Kafka（维护成本高）
// rabbitmq 异步场景 将信息写入数据库+发送成功邮件 解耦 削峰
func InitRabbiltMQ() {
	//

}
