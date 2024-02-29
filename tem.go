// // Task for you
// // Read about redis,
// // write a code in go to connect to redis,
// // Producer: generate random values of age and push it to redis
// // Consumer: consume values from redis and find average and total age.
// // Match the values from producer and consumer.
// // time.Sleep(1 * time.Second) // Wait for 1 second before consuming the next item
// package main

// import (
// 	"context"
// 	"fmt"
// 	"log/slog"
// 	"math/rand"
// 	"os"
// 	"strconv"
// 	"strings"
// 	"sync"
// 	"time"

// 	"github.com/joho/godotenv"
// 	"github.com/redis/go-redis/v9"
// )

// var ctx = context.Background()
// var totalAgeProduce, totalCharProduce = 0, 0 
// var totalAgeConsume, totalCharConsume = 0, 0
// var mutexProduce = &sync.Mutex{}

// func main() {
// 	err := godotenv.Load("config.env")
// 	if err != nil {
// 		slog.Error("Error loading .env file:", err)
// 	}
	
// 	// Access the environment variables
// 	dbHost := os.Getenv("DB_HOST")
// 	dbPort := os.Getenv("DB_PORT")
// 	dbPass := os.Getenv("DB_PASS")
// 	dbIndex, _ := strconv.Atoi(os.Getenv("DB_INDEX"))

// 	// Connection
// 	client := redis.NewClient(&redis.Options{
// 		Addr: fmt.Sprintf("%s:%s", dbHost, dbPort),
// 		Password: dbPass,
// 		DB: dbIndex,
// 	})

// 	fmt.Print("Producing...\n")
// 	createProducers(client, 1000, 100)
	
// 	fmt.Print("Consuming...\n")
// 	consumer(client, "people")
		
// 	fmt.Printf("Total Characters Produced: %d, Total Age Produced: %d", totalCharProduce, totalAgeProduce)
// 	fmt.Printf("\nTotal Characters Consumed: %d, Total Age Consumed: %d", totalCharConsume, totalAgeConsume)

// 	if totalAgeProduce == totalAgeConsume && totalCharProduce == totalCharConsume {
// 		fmt.Print("\nChecksum Matched!")
// 	} else {
// 		fmt.Print("\nChecksum Didn't Matched")
// 	}

// 	client.Close()
// }

// // producer: Producer function produces random names & age as string & integer respectively for provided count & then it pushes it in redis database as a list data structure
// func producer(client *redis.Client, count int) {
// 	for i:=0; i<count; i++ {

// 		// Preventing race conditions
// 		name := generateRandomName(rand.Intn(200)) //Generate random name of 0 to 200 characters
// 		age := rand.Intn(100) //Generate rantom age between 0 to 100
		
// 		// Adding total age & total characters that are being produced
// 		mutexProduce.Lock()
//         totalAgeProduce += age
//         totalCharProduce += len(name)
// 		mutexProduce.Unlock()
		
// 		fmt.Printf("Pushed:\n%s:%d\n", name, age)
// 		fmt.Print("\n")
		
// 		err := client.LPush(ctx, "people", fmt.Sprintf("%s:%d", name, age)).Err() //Push data into DB
// 		if err != nil {
// 			slog.Error("Error pushing to Redis:", err)
// 		}
		
// 		time.Sleep(1 * time.Millisecond)
// 	}
// }

// // createProducers: This function takes 3 arguments (Redis client, total number of producers & total producers per thread), it calls producer function to then produce via multi threading using go routines.
// func createProducers(client *redis.Client, totalProducers, producersPerThread int) {
// 	var wg sync.WaitGroup //A WaitGroup waits for a collection of goroutines to finish. The main goroutine calls Add to set the number of goroutines to wait for. Then each of the goroutines runs and calls Done when finished. At the same time, Wait can be used to block until all goroutines have finished.

// 	threads := totalProducers / producersPerThread
// 	for i := 0; i < threads; i++ {
// 		wg.Add(1)

// 		// go routine for multi-threading
// 		go func() {
// 			producer(client, producersPerThread)
// 			defer wg.Done()
// 		}()
// 	}

// 	remaining := totalProducers % producersPerThread
// 	if remaining > 0 {
// 		wg.Add(1)
// 		// go routine for multi-threading
// 		go func() {
// 			producer(client, remaining)
// 			defer wg.Done()
// 		}()
// 	}

// 	// Wait for all go routines to get completed
// 	wg.Wait()
// }

// // consumer: Consumer function pops the value from redis database & also calcuates total age count & average of age while consuming in each step 
// func consumer(client *redis.Client, key string) {
// 	for {
// 		exists, err := client.Exists(context.Background(), key).Result()
// 		if err != nil {
// 			slog.Error("List does not exist:", err)
// 		}

// 		if exists > 0 {
// 			val, err := client.LPop(context.Background(), key).Result()
// 			if err != nil {
// 				slog.Error("Error in popping the result:", err)
// 			}

// 			data := splitValue(val, ":")

// 			age, err := strconv.Atoi(data[1])
// 			if err != nil {
// 				slog.Error("Error in parsing integer:", err)
// 			}

// 			totalAgeConsume += age
// 			totalCharConsume += len(data[0])

// 			fmt.Printf("Popped:\n%s:%d\n", data[0], age)
// 			fmt.Print("\n")
// 		} else {
// 			break
// 		}

// 		time.Sleep(1 * time.Millisecond)
// 	}
// }


// // generateRandomName: This function generates random string of length passed as an argument.
// func generateRandomName(length int) string {
// 	const charSet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
// 	name := make([]byte, length) //Make empty array for provided length
// 	for i := range name {
// 		name[i] = charSet[rand.Intn(len(charSet))] //Generate random string for provided length
// 	}
// 	return string(name)
// }

// // splitValue: This function takes a string & separator as arguments & separates the string with provided separator & return string elements in form of array
// func splitValue(value, separator string) []string {
// 	return strings.Split(value, separator)
// }

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// // Task for you
// // Read about redis,
// // write a code in go to connect to redis,
// // Producer: generate random values of age and push it to redis
// // Consumer: consume values from redis and find average and total age.
// // Match the values from producer and consumer.
// // time.Sleep(1 * time.Second) // Wait for 1 second before consuming the next item
// package main

// import (
// 	"context"
// 	"fmt"
// 	"log/slog"
// 	"math/rand"
// 	"os"
// 	"strconv"
// 	"sync"
// 	"time"

// 	"github.com/joho/godotenv"
// 	"github.com/redis/go-redis/v9"
// )

// var ctx = context.Background()
// var totalAgeProduce, totalCharProduce = 0, 0 
// var totalAgeConsume, totalCharConsume = 0, 0
// var mutexProduce = &sync.Mutex{}

// func main() {
// 	err := godotenv.Load("config.env")
// 	if err != nil {
// 		slog.Error("Error loading .env file:", err)
// 	}
	
// 	// Access the environment variables
// 	dbHost := os.Getenv("DB_HOST")
// 	dbPort := os.Getenv("DB_PORT")
// 	dbPass := os.Getenv("DB_PASS")
// 	dbIndex, _ := strconv.Atoi(os.Getenv("DB_INDEX"))

// 	// Connection
// 	client := redis.NewClient(&redis.Options{
// 		Addr: fmt.Sprintf("%s:%s", dbHost, dbPort),
// 		Password: dbPass,
// 		DB: dbIndex,
// 	})

// 	fmt.Print("Producing...\n")
// 	createProducers(client, 1000, 100)

// 	time.Sleep(1 * time.Second) // Wait for second
	
// 	fmt.Print("Consuming...\n")
// 	consumer(client)
		
// 	fmt.Printf("Total Characters Produced: %d, Total Age Produced: %d", totalCharProduce, totalAgeProduce)
// 	fmt.Printf("\nTotal Characters Consumed: %d, Total Age Consumed: %d", totalCharConsume, totalAgeConsume)

// 	if totalAgeProduce == totalAgeConsume && totalCharProduce == totalCharConsume {
// 		fmt.Print("\nChecksum Matched!")
// 	} else {
// 		fmt.Print("\nChecksum Didn't Matched")
// 	}

// 	client.Close()
// }

// // producer: Producer function produces random names & age as string & integer respectively for provided count & then it pushes it in redis database as a list data structure
// func producer(client *redis.Client, count int) {
// 	for i:=0; i<count; i++ {

// 		// Preventing race conditions
// 		name := generateRandomName(rand.Intn(199) + 1) //Generate random name of 0 to 200 characters
// 		age := rand.Intn(99) + 1 //Generate rantom age between 0 to 100
		
// 		// Adding total age & total characters that are being produced
// 		mutexProduce.Lock()
//         totalAgeProduce += age
//         totalCharProduce += len(name)
// 		mutexProduce.Unlock()
		
// 		// fmt.Printf("Pushed:\n%s:%d\n", name, age)
// 		fmt.Print(age)
// 		fmt.Print("\n")
		
// 		err := client.HSet(ctx, "people", name, age).Err() //Push data into DB
// 		if err != nil {
// 			slog.Error("Error pushing to Redis:", err)
// 		}
		
// 		time.Sleep(1 * time.Millisecond)
// 	}
// }

// // createProducers: This function takes 3 arguments (Redis client, total number of producers & total producers per thread), it calls producer function to then produce via multi threading using go routines.
// func createProducers(client *redis.Client, totalProducers, producersPerThread int) {
// 	var wg sync.WaitGroup //A WaitGroup waits for a collection of goroutines to finish. The main goroutine calls Add to set the number of goroutines to wait for. Then each of the goroutines runs and calls Done when finished. At the same time, Wait can be used to block until all goroutines have finished.

// 	threads := totalProducers / producersPerThread
// 	for i := 0; i < threads; i++ {
// 		wg.Add(1)

// 		// go routine for multi-threading
// 		go func() {
// 			producer(client, producersPerThread)
// 			defer wg.Done()
// 		}()
// 	}

// 	remaining := totalProducers % producersPerThread
// 	if remaining > 0 {
// 		wg.Add(1)
// 		// go routine for multi-threading
// 		go func() {
// 			producer(client, remaining)
// 			defer wg.Done()
// 		}()
// 	}

// 	// Wait for all go routines to get completed
// 	wg.Wait()
// }

// // consumer: Consumer function pops the value from redis database & also calcuates total age count & average of age while consuming in each step 
// func consumer(client *redis.Client) {
// 	// Fetch all data of hash into map
// 	values, err := client.HGetAll(ctx, "people").Result()
// 	if err != nil {
// 		slog.Error("Error fetching hash from Redis:", err)
// 		return
// 	}

// 	// Iterate over the map & pop each fields from hash
// 	for character, ageStr := range values {
// 		age, err := strconv.Atoi(ageStr)
// 		if err != nil {
// 			slog.Error("Error parsing age:", err)
// 			time.Sleep(1 * time.Minute)
// 			continue
// 		}

//         totalAgeConsume += age
//         totalCharConsume += len(character)

// 		// fmt.Printf("Popped:\n%s:%d\n", character, age)
// 		fmt.Print(age)
// 		fmt.Print("\n")

// 		err = client.HDel(ctx, "people", character).Err()
// 		if err != nil {
// 			slog.Error("Error deleting a field from Redis:", err)
// 			time.Sleep(1 * time.Minute)
// 			continue
// 		}

// 		time.Sleep(1 * time.Millisecond)
// 	}
// }

// // generateRandomName: This function generates random string of length passed as an argument.
// func generateRandomName(length int) string {
// 	const charSet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
// 	name := make([]byte, length) //Make empty array for provided length
// 	for i := range name {
// 		name[i] = charSet[rand.Intn(len(charSet))] //Generate random string for provided length
// 	}
// 	return string(name)
// }

// // // splitValue: This function takes a string & separator as arguments & separates the string with provided separator & return string elements in form of array
// // func splitValue(value, separator string) []string {
// // 	return strings.Split(value, separator)
// // }