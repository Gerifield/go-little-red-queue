package littleredqueue

import (
	"errors"
	"fmt"

	"github.com/gomodule/redigo/redis"
)

const (
	PRIORITY_HIGH   = "high"
	PRIORITY_NORMAL = "normal"
	PRIORITY_LOW    = "low"
)

type Queue struct {
	Redis  redis.Conn
	Prefix string
}

//Create a new queue
func NewQueue(conn redis.Conn) *Queue {
	return &Queue{
		Redis:  conn,
		Prefix: "queue",
	}
}

//Create a queue with a specific prefix
func NewQueueWithPrefix(conn redis.Conn, prefix string) *Queue {
	return &Queue{
		Redis:  conn,
		Prefix: prefix,
	}
}

//Get an element form queue
//At timeout the result and the error'll be nil
func (q *Queue) Get(key string, timeout int64) (interface{}, error) {
	params := []interface{}{
		q.getQueueKey(key, PRIORITY_HIGH),
		q.getQueueKey(key, PRIORITY_NORMAL),
		q.getQueueKey(key, PRIORITY_LOW),
		timeout,
	}

	res, err := q.Redis.Do("BLPOP", params...)

	if res != nil && err == nil {
		if a, ok := res.([]interface{}); ok && len(a) == 2 {
			return a[1], nil
		}
	}
	return nil, err
}

//Get bytes from queue
func (q *Queue) GetBytes(key string, timeout int64) ([]byte, error) {
	r, err := q.Get(key, timeout)

	if err == nil && r != nil {
		if res, ok := r.([]byte); ok {
			return res, nil
		} else {
			return nil, errors.New("Conversion error")
		}
	}

	return nil, err
}

//Get a string element from the queue
func (q *Queue) GetString(key string, timeout int64) (string, error) {
	r, err := q.GetBytes(key, timeout)
	return string(r), err
}

// Put a value into the queue
func (q *Queue) put(key, priority string, value interface{}) (int64, error) {
	return redis.Int64(q.Redis.Do("RPUSH", q.getQueueKey(key, priority), value))
}

//Put an element into the queue with normal priority
func (q *Queue) PutNormal(key string, value interface{}) (int64, error) {
	return q.put(key, PRIORITY_NORMAL, value)
}

//Put an element into the queue with high priority
func (q *Queue) PutHigh(key string, value interface{}) (int64, error) {
	return q.put(key, PRIORITY_HIGH, value)
}

//Put an element into the queue with low priority
func (q *Queue) PutLow(key string, value interface{}) (int64, error) {
	return q.put(key, PRIORITY_LOW, value)
}

func (q *Queue) getQueueKey(key, priority string) string {
	return fmt.Sprintf("%s:%s:%s", q.Prefix, priority, key)
}
