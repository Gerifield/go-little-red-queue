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

	fmt.Println("Start consume")
	for {
		res, err := queue.GetString("testKey", 5)
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
