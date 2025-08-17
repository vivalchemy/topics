package main

import (
	"encoding/json"
	"log"

	"github.com/IBM/sarama"
	"github.com/gofiber/fiber/v3"
)

type Comment struct {
	Text string `form:"text" json:"text"`
}

func PushCommentToQueue(topic string, data []byte) error {
	brokersUrl := []string{"kafka-broker:29092"}
	producer, err := ConnectProducer(brokersUrl)
	if err != nil {
		return err
	}
	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(data),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		return err
	}

	log.Printf("Message produced to topic %s [partition=%d, offset=%d]\n", topic, partition, offset)

	return nil
}

func ConnectProducer(brokersUrl []string) (sarama.SyncProducer, error) {
	// TODO: Add error handling
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	conn, err := sarama.NewSyncProducer(brokersUrl, config)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func createComment(c fiber.Ctx) error {
	cmt := new(Comment)

	if err := c.Bind().Body(cmt); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"message": "Invalid request body",
		})
	}

	cmtInBytes, err := json.Marshal(cmt)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": "Failed to marshal comment",
		})
	}

	err = PushCommentToQueue("comment", cmtInBytes)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": "Failed to push comment to queue",
		})
	}

	// Respond on success
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
		"message": "Comment created successfully",
		"comment": cmt,
	})
}

func main() {
	log.Println("Starting server...")
	app := fiber.New()

	api := app.Group("/api/v1")
	api.Post("/comments", createComment)

	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}
