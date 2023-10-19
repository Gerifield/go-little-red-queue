package main

import (
	"fmt"

	littleredqueue "github.com/gerifield/go-little-red-queue"
	"github.com/gomodule/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "192.168.99.100:32768")
	if err != nil {
		fmt.Println("Connection error", err)
		return
	}
	defer conn.Close()

	queue := littleredqueue.NewQueue[string](conn)

	fmt.Println("Start consume")
	for {
		res, err := queue.Get("testKey", 5)
		if err != nil {
			panic(err)
		}

		if res == "" {
			fmt.Println("Timeout!")
			break
		} else {
			fmt.Println("Res:", res)
		}
	}

	fmt.Println("Finished")
}
