package littleredqueue

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

const (
	PRIORITY_HIGH   = "high"
	PRIORITY_NORMAL = "normal"
	PRIORITY_LOW    = "low"
)

type Queue[T any] struct {
	Redis  redis.Conn
	Prefix string
}

// NewQueue create a new queue
func NewQueue[T any](conn redis.Conn) *Queue[T] {
	return &Queue[T]{
		Redis:  conn,
		Prefix: "queue",
	}
}

// NewQueueWithPrefix create a queue with a specific prefix
func NewQueueWithPrefix[T any](conn redis.Conn, prefix string) *Queue[T] {
	return &Queue[T]{
		Redis:  conn,
		Prefix: prefix,
	}
}

// Get an element form queue
// At timeout the result and the error will be nil
func (q *Queue[T]) Get(key string, timeout int64) (T, error) {
	params := []interface{}{
		q.getQueueKey(key, PRIORITY_HIGH),
		q.getQueueKey(key, PRIORITY_NORMAL),
		q.getQueueKey(key, PRIORITY_LOW),
		timeout,
	}

	res, err := q.Redis.Do("BLPOP", params...)

	if res != nil && err == nil {
		if a, ok := res.([]interface{}); ok && len(a) == 2 {
			return a[1].(T), nil
		}
	}

	// return empty value if nothing found
	var ret T

	return ret, err
}

// Put a value into the queue
func (q *Queue[T]) put(key, priority string, value T) (int64, error) {
	return redis.Int64(q.Redis.Do("RPUSH", q.getQueueKey(key, priority), value))
}

// PutNormal put an element into the queue with normal priority
func (q *Queue[T]) PutNormal(key string, value T) (int64, error) {
	return q.put(key, PRIORITY_NORMAL, value)
}

// PutHigh put an element into the queue with high priority
func (q *Queue[T]) PutHigh(key string, value T) (int64, error) {
	return q.put(key, PRIORITY_HIGH, value)
}

// PutLow put an element into the queue with low priority
func (q *Queue[T]) PutLow(key string, value T) (int64, error) {
	return q.put(key, PRIORITY_LOW, value)
}

func (q *Queue[T]) getQueueKey(key, priority string) string {
	return fmt.Sprintf("%s:%s:%s", q.Prefix, priority, key)
}
