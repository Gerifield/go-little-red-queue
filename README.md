# go-little-red-queue [![Build Status](https://next.travis-ci.org/Gerifield/go-little-red-queue.svg?branch=master)](https://next.travis-ci.org/Gerifield/go-little-red-queue) [![Documentation](https://godoc.org/github.com/Gerifield/go-little-red-queue?status.svg)](https://godoc.org/github.com/Gerifield/go-little-red-queue)
Go version of the little-red-queue php library
[PHP version](https://github.com/Gerifield/little-red-queue)

It supports priority, at 3 levels: low, normal, high
These use different queues, the get method'll choose a value from them.

## Install

Via go get

``` bash
$ go get github.com/Gerifield/go-little-red-queue
```
## Example

``` go
package main
import (
	"github.com/Gerifield/go-little-red-queue"
	"github.com/garyburd/redigo/redis"
)

queue := littleredqueue.NewQueue(redis.Dial("tcp", "host:port"))

//Put an element into the queue with normal priority
l, _ := queue.PutNormal("key", "value")
fmt.Printf("QueueLength: %d", l)

//Put an element into the queue with high priority
l, err := queue.PutHigh("key", "value")
fmt.Printf("QueueLength: %d\n", l)

//Get a value with 5 seconds timeout, 0 means infite blocking
res, err := queue.GetString("test2", 5)
if err != nil {
	panic(err)
}
if res == nil {
	fmt.Println("Timeout")
} else {
	fmt.Printf("Value: %s\n", res)
}
```

## Testing

``` bash
$ go test .
```

## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.
