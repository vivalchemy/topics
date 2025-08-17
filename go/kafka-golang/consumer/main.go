package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/IBM/sarama"
)

func CreateConsumer(brokersUrl []string) (sarama.Consumer, error) {
	// TODO: Add error handling
	config := sarama.NewConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest // pass it to the function so that it can resume where it left off
	config.Consumer.Return.Errors = true

	conn, err := sarama.NewConsumer(brokersUrl, config)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func main() {
	topic := "comment"
	worker, err := CreateConsumer([]string{"kafka-broker:29092"})
	if err != nil {
		log.Fatal(err)
	}

	consumer, err := worker.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("consumer started")

	//graceful shutdown
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	doneCh := make(chan struct{})

	msgCount := 0
	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				log.Println("error:", err)
			case msg := <-consumer.Messages():
				msgCount++
				log.Printf("Message on %s [offset=%d, key=%s, value=%s]\n", msg.Topic, msg.Offset, string(msg.Key), string(msg.Value))
			case interrupt := <-sigchan:
				log.Println("shutting down consumer interup detected", interrupt.String())
				consumer.AsyncClose()
				doneCh <- struct{}{}
				return
			}

		}
	}()
	<-doneCh

	log.Println("consumer stopped")
	log.Println("Processed message count:", msgCount)

	if err := worker.Close(); err != nil {
		panic(err)
	}

}
