package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	_ "github.com/joho/godotenv/autoload"
	"github.com/redis/go-redis/v9"
)

var ctx context.Context = context.Background()

type CacheClient interface {
	Set(key string, data *UserToken) error
	Get(key string) (*UserToken, error)
	Del(keys ...string) error
	DelOne(key string) error
}

type UserToken struct {
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	ExpiryDate time.Time `json:"expiry_date"`
}

func NewUserToken(name string, email string) (string, *UserToken, error) {
	randomUUID, err := uuid.NewRandom()
	if err != nil {
		log.Println("Randon uuid is created found")
		return "", nil, err
	}

	timeDay := 24 * time.Hour
	token := UserToken{
		Name:       name,
		Email:      email,
		ExpiryDate: time.Now().Add(50 * timeDay),
	}

	return randomUUID.String(), &token, nil
}

func main() {
	var url string
	if os.Getenv("RUNNING_IN_DOCKER") == "true" {
		log.Println("Running in docker container")
		url = "redis:6379" // matches docker-compose service name
	} else {
		log.Println("Running locally")
		url = "127.0.0.1:6379"
	}

	clientOpts := redis.Options{
		Addr:     url,
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	}

	rdsClient, err := NewRedisClient(&clientOpts, 5*time.Second)
	if err != nil {
		log.Fatalf(err.Error(), "Error creating redis client")

	}

	uuid, token, err := NewUserToken("vivian", "vivian@gmail.com")
	if err != nil {
		log.Println("Error creating user token")
		return
	}

	err = rdsClient.Set(uuid, token)
	if err != nil {
		log.Println(err, "Error setting user token")
		return
	}

	token, err = rdsClient.Get(uuid)
	if err != nil {
		log.Println(err, "Error getting user token")
		return
	}
	log.Println(token)

	err = rdsClient.DelOne(uuid)
	if err != nil {
		log.Println(err, "Error deleting user token")
		return
	}

	// this should give error
	token, err = rdsClient.Get(uuid)
	if err != nil {
		log.Println(err, "Error getting user token")
		return
	}
	log.Println(token)
}
