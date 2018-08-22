package config

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

type RabbitConfig struct {
	mydb      *amqp.Connection
	myqueue   amqp.Queue
	mychannel *amqp.Channel
}

var rabbit RabbitConfig

func (r *RabbitConfig) Configure() {
	var err, err1 error
	r.mychannel, err = r.mydb.Channel()
	failOnError(err, "Failed to open a channel")

	r.myqueue, err1 = r.mychannel.QueueDeclare(
		"hello", // name
		true,    // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	err = r.mychannel.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err1, "Failed to declare a queue")

	// PUBLISH
}

func Publish(msg string) {
	rabbit.Publish(msg)
}
func (r *RabbitConfig) Publish(msg string) {
	err := r.mychannel.Publish(
		"",             // exchange
		r.myqueue.Name, // routing key
		false,          // mandatory
		false,          // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(msg),
		})
	if err != nil {
		log.Fatalf("%s: %s", "Error while sending message", err)
	}
}
func NewRabbitConfig(conn *amqp.Connection) *RabbitConfig {
	rabbit = RabbitConfig{mydb: conn}
	return &rabbit
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func CloaseAll() {
	rabbit.mychannel.Close()
}
