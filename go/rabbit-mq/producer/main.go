package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/gofiber/fiber/v3"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Comment struct {
	Text string `form:"text" json:"text"`
}

func InitializeMQ(url string, queueName string) (*amqp.Connection, *amqp.Channel, amqp.Queue) {
	conn, err := amqp.Dial(url)
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	q, err := ch.QueueDeclare(queueName, false, false, false, false, nil)
	failOnError(err, "Failed to declare a queue")

	return conn, ch, q
}

func PushCommentToMQ(ch *amqp.Channel, q amqp.Queue, cmtInBytes []byte) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return ch.PublishWithContext(ctx, "", q.Name, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        cmtInBytes,
	})
}

func createCommentHandler(ch *amqp.Channel, q amqp.Queue) fiber.Handler {
	return func(c fiber.Ctx) error {
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

		err = PushCommentToMQ(ch, q, cmtInBytes)
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
}

func main() {
	conn, ch, q := InitializeMQ("amqp://guest:guest@rabbitmq:5672/", "comments")
	defer ch.Close()
	defer conn.Close()

	log.Println("Starting server...")
	app := fiber.New()

	api := app.Group("/api/v1")
	api.Post("/comments", createCommentHandler(ch, q))

	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
