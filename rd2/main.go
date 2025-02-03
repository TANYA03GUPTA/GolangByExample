package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis"
)

var ctx = context.Background()

func main() {
 // Connect to Redis
 client := redis.NewClient(&redis.Options{
  Addr:     "localhost:6379", // Replace with your Redis server address
  Password: "",                // No password for local development
  DB:       0,                 // Default DB
 })

 // Ping the Redis server to check the connection
 pong, err := client.Ping(ctx).Result()
 if err != nil {
  log.Fatal("Error connecting to Redis:", err)
 }
 fmt.Println("Connected to Redis:", pong)
}