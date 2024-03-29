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

	resb, err := queue.Get("test", 1)
	fmt.Println("Res", resb, "Err", err)

	res, err := queue.PutNormal("test2", "almafa")
	fmt.Println("Res", res, "Err", err)

	queue2 := littleredqueue.NewQueue[int](conn)
	res, err = queue2.PutNormal("test2", 5)
	fmt.Println("Res", res, "Err", err)
}
