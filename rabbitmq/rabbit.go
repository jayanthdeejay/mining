package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/jayanthdeejay/mining/config"
	amqp "github.com/rabbitmq/amqp091-go"
)

func BunnyOpen() (*amqp.Channel, *amqp.Queue, context.Context) {
	// Read the config file
	file, err := os.Open("../config/rabbit.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := config.BunnyConfig{}
	err = decoder.Decode(&config)
	if err != nil {
		panic(err)
	}

	// Create the connection to the RabbitMQ server
	uriString := fmt.Sprintf("amqp://%s:%s@%s:%s/", config.RabbitUser, config.RabbitPwd, config.RabbitHost, config.RabbitPort)
	conn, err := amqp.Dial(uriString)
	if err != nil {
		panic(err)
	}

	// Create a channel for sending and receiving messages
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	// Declare a queue to which the generated keys will be sent
	queue, err := ch.QueueDeclare(
		config.QueueName,        // queue name
		config.Durable,          // durable
		config.DeleteWhenUnused, // delete when unused
		config.Exclusive,        // exclusive
		config.NoWait,           // no-wait
		nil,                     // arguments
	)
	if err != nil {
		panic(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), time.Second)

	return ch, &queue, ctx
}
