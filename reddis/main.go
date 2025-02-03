package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

// Global Redis context
var ctx = context.Background()

// Function to connect to Redis
func connectRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis default port
	})
	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
	fmt.Println("Connected to Redis!")
	return client
}

// Function to demonstrate List operations
func listOperations(client *redis.Client) {
	fmt.Println("\n🔹 Working with Redis Lists")

	client.RPush(ctx, "fruits", "apple", "banana", "cherry")
	list, _ := client.LRange(ctx, "fruits", 0, -1).Result()
	fmt.Println("List elements:", list)

	popped, _ := client.LPop(ctx, "fruits").Result()
	fmt.Println("Popped element:", popped)

	newList, _ := client.LRange(ctx, "fruits", 0, -1).Result()
	fmt.Println("Updated list:", newList)
}

// Function to demonstrate Hash operations
func hashOperations(client *redis.Client) {
	fmt.Println("\n🔹 Working with Redis Hashes")

	client.HSet(ctx, "user:1001", "name", "Alice", "age", "30", "city", "New York")
	name, _ := client.HGet(ctx, "user:1001", "name").Result()
	fmt.Println("User Name:", name)

	userData, _ := client.HGetAll(ctx, "user:1001").Result()
	fmt.Println("Complete User Data:", userData)
}

// Function to demonstrate Key Expiration
func expireKeyOperations(client *redis.Client) {
	fmt.Println("\n🔹 Setting Expiry on Keys")

	client.Set(ctx, "session:xyz", "active", 5*time.Second)
	fmt.Println("Key 'session:xyz' set with 5 seconds expiry.")

	time.Sleep(3 * time.Second)
	ttl, _ := client.TTL(ctx, "session:xyz").Result()
	fmt.Println("Remaining TTL:", ttl)

	time.Sleep(3 * time.Second)
	val, err := client.Get(ctx, "session:xyz").Result()
	if err == redis.Nil {
		fmt.Println("Key expired!")
	} else {
		fmt.Println("Key still exists:", val)
	}
}

// Function to demonstrate Pub/Sub
func pubSubOperations(client *redis.Client) {
	fmt.Println("\n🔹 Working with Pub/Sub")

	// Subscribe to a channel
	pubsub := client.Subscribe(ctx, "news")
	ch := pubsub.Channel()

	// Goroutine to receive messages
	go func() {
		for msg := range ch {
			fmt.Println("📩 Received message:", msg.Payload)
		}
	}()

	// Publish messages
	client.Publish(ctx, "news", "Breaking News: Redis and Go!").Err()
	client.Publish(ctx, "news", "Another update: Pub/Sub is cool!").Err()

	// Give some time for the messages to be received
	time.Sleep(2 * time.Second)

	// Close subscription
	_ = pubsub.Close()
}

func main() {
	client := connectRedis()

	// Execute Redis operations
	listOperations(client)
	hashOperations(client)
	expireKeyOperations(client)
	pubSubOperations(client)

	// Close Redis connection
	client.Close()
	fmt.Println("\n✅ All Redis operations completed successfully!")
}
