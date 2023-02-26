package config

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type BunnyConfig struct {
	RabbitUser       string     `json:"rabbit_user"`
	RabbitPwd        string     `json:"rabbit_pwd"`
	ProducerHost     string     `json:"producer_host"`
	ConsumerHost     string     `json:"consumer_host"`
	RabbitPort       string     `json:"rabbit_port"`
	QueueName        string     `json:"queue_name"`
	Durable          bool       `json:"durable"`
	DeleteWhenUnused bool       `json:"delete_when_unused"`
	Exclusive        bool       `json:"exclusive"`
	NoWait           bool       `json:"no_wait"`
	Args             amqp.Table `json:"args"`
	ConsumerName     string     `json:"consumer_name"`
	AutoAck          bool       `json:"auto_ack"`
	NoLocal          bool       `json:"no_local"`
	MsgNoWait        bool       `json:"msg_no_wait"`
}
