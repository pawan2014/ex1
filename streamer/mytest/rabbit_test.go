package mytest

import (
	"testing"

	"github.com/ex1/streamer/config"
	"github.com/streadway/amqp"
)

func TestMessageSend1(t *testing.T) {
	// MQ Connection
	rabbitconn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	defer config.CloaseAll()
	defer rabbitconn.Close()
	if err != nil {
		panic(err)
	}
	c := config.NewRabbitConfig1(rabbitconn)
	c.Configure()
	c.Publish("test")

}
