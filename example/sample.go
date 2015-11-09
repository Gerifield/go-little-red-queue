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

	resb, err := queue.Get("test", 1)
	fmt.Println("Res", resb, "Err", err)

	res, err := queue.PutNormal("test2", "almafa")
	fmt.Println("Res", res, "Err", err)

	res, err = queue.PutNormal("test2", 5)
	fmt.Println("Res", res, "Err", err)

	resb, err = queue.GetString("test2", 1)
	fmt.Println("Res", resb, "Err", err)
	resb, err = queue.GetString("test2", 1)
	fmt.Println("Res", resb, "Err", err)
	resb, err = queue.GetString("test2", 1)
	fmt.Println("Res", resb, "Err", err)
}
