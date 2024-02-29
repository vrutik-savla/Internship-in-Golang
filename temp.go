// package main

// import (
// 	"log"
// 	"os"

// 	. "github.com/remiges-tech/logharbour/logharbour"
// )

// // Example of using logharbour package
// func main() {
// 	// Open a file for logging.
// 	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()

// 	// Create a fallback writer that uses the file as the primary writer and stdout as the fallback.
// 	fallbackWriter := NewFallbackWriter(file, os.Stdout)

	

// 	// Initialize the logger.
// 	logger := NewLogger(LoggerContext, "MyApp", fallbackWriter)

// 	// Create a new logger with various fields set.
// 	logger = logger.WithModule("Module1").
// 		WithWho("John Doe").
// 		WithStatus(Success).
// 		WithRemoteIP("192.168.1.1")

// 	// Use the new logger to log an activity.
// 	logger.LogActivity("User logged in", map[string]any{"username": "john"})

// 	// Log a data change entry.
// 	logger.LogDataChange("User updated profile", ChangeInfo{
// 		Entity:    "User",
// 		Operation: "Update",
// 	})
// }

// // Example struct representing your JSON data
// type Person struct {
// 	Name string `json:"name"`
// 	Age  int    `json:"age"`
// }

// // Serialize converts a struct to JSON string
// func Serialize(person Person) string {
// 	data, _ := json.Marshal(person)
// 	return string(data)
// }

// // Deserialize converts a JSON string to a struct
// func Deserialize(data string) Person {
// 	var person Person
// 	_ = json.Unmarshal([]byte(data), &person)
// 	return person
// }

// func main() {
// 	ctx := context.Background()

// 	// Connect to Redis
// 	client := redis.NewClient(&redis.Options{
// 		Addr: "localhost:6379", // Replace with your Redis server address
// 		Password: "",
// 		DB: 2,
// 	})

// 	// Close the connection before the main function exits
// 	defer client.Close()

// 	// Example struct to store in Redis
// 	person := Person{Name: "John", Age: 30}

// 	// Serialize the struct to JSON string and store in Redis
// 	serialized := Serialize(person)
// 	err := client.Set(ctx, "person", serialized, 0).Err()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Retrieve from Redis and deserialize back to struct
// 	data, err := client.Get(ctx, "person").Result()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	retrievedPerson := Deserialize(data)
// 	fmt.Println("Retrieved Person:")
// 	fmt.Println(retrievedPerson)
// }



// var totalProducers, producersPerThread int
/// fmt.Println("Enter total number of Producers: ")
// fmt.Scanln(&totalProducers)
// fmt.Println("Enter total number of Producers per Thread: ")
// fmt.Scanln(&producersPerThread)
/* 
package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	ctx := context.Background()

	// err := client.Set(ctx, "foo", "bar", 0).Err()
	// if err != nil {
	// 	panic(err)
	// }
	// val, err := client.Get(ctx, "foo").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("foo", val)

	session := map[string]string{
		"name": "Vrutik",
		"surname": "Savla",
		"company": "Remiges",
		"age": "20",
	}
	for k, v := range session {
		err := client.HSet(ctx, "user-session:123", k, v).Err()
		if err != nil {
			panic(err)
		}
	}
	userSession := client.HGetAll(ctx, "user-session:123").Val()
	fmt.Println(userSession)
}
*/