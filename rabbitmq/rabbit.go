package rabbitmq

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jayanthdeejay/mining/config"
	amqp "github.com/rabbitmq/amqp091-go"
)

func BunnyCon() (*amqp.Channel, *amqp.Queue) {
	// Read the config file
	file, err := os.Open("config.json")
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
	uriString := fmt.Sprintf("amqp://%s:%s@%s:%d/", config.RabbitUser, config.RabbitPwd, config.ProducerHost, config.RabbitPort)
	conn, err := amqp.Dial(uriString)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Create a channel for sending and receiving messages
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

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

	return ch, &queue
}
