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

	fmt.Println("Start publish")
	for i := 0; i < 1000; i++ {
		if i%2 == 0 {
			_, _ = queue.PutNormal("testKey", "NORM")
		} else {
			_, _ = queue.PutHigh("testKey", "HI")
		}

	}

	fmt.Println("Finished")
}
