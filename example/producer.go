package main

import (
	"fmt"

	"github.com/Gerifield/go-little-red-queue"
	"github.com/garyburd/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "192.168.99.100:32768")
	if err != nil {
		fmt.Println("Connection error", err)
		return
	}
	defer conn.Close()

	queue := littleredqueue.NewQueue(conn)

	fmt.Println("Start publish")
	for i := 0; i < 1000; i++ {
		if i%2 == 0 {
			queue.PutNormal("testKey", "NORM")
		} else {
			queue.PutHigh("testKey", "HI")
		}

	}

	fmt.Println("Finished")
}
