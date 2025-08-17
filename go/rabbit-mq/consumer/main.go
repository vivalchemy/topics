package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	amqp "github.com/rabbitmq/amqp091-go"
)

func InitializeMQ(url string, queueName string) (*amqp.Connection, *amqp.Channel, amqp.Queue) {
	conn, err := amqp.Dial(url)
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	q, err := ch.QueueDeclare(queueName, false, false, false, false, nil)
	failOnError(err, "Failed to declare a queue")

	return conn, ch, q
}

func main() {
	conn, ch, q := InitializeMQ("amqp://guest:guest@rabbitmq:5672/", "comments")
	defer ch.Close()
	defer conn.Close()

	msgs, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	failOnError(err, "Failed to register a consumer")

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	var doneChan chan struct{}

	msgCount := 0

	go func() {
		for {
			select {
			case d := <-msgs:
				msgCount++
				log.Printf("Received a message: %s", d.Body)
			case interrupt := <-signalChan:
				log.Printf("Received interrupt signal: %v", interrupt)
				doneChan <- struct{}{}
				return
			}
		}
	}()

	log.Printf("[*] Waiting for messages. To exit press CTRL+C")
	<-doneChan
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
