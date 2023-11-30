package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/go-redis/redis"
)

func main() {
	// Connect to Redis
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// Open the text file
	file, err := os.Open("/media/jay/slimshady/addresses/addresses.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read the file line by line and add each string to a Redis set
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		err := client.SAdd("addresses", text).Err()
		if err != nil {
			panic(err)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Print the number of strings added to the Redis set
	count, err := client.SCard("addresses").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d strings added to Redis set 'addresses'\n", count)
}

