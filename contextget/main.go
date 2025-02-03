package main

import (
	"context"
	"fmt"
)

func main() {
	// Step 1: Create a base context (background context)
	ctx := context.Background()

	// Step 2: Add a value to the context using context.WithValue
	key := "userID"           // Key (can be any type, but typically a unique type)
	value := 12345            // Value to store
	ctx = context.WithValue(ctx, key, value)

	// Step 3: Retrieve the value using context.Value
	retrievedValue := ctx.Value(key)

	// Step 4: Check if value exists
	if retrievedValue != nil {
		fmt.Printf("Retrieved value: %v\n", retrievedValue)
	} else {
		fmt.Println("No value found for key.")
	}
}
