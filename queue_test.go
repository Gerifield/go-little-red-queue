package littleredqueue

import (
	"errors"
	"reflect"
	"testing"

	"github.com/rafaeljusto/redigomock"
)

func TestGet(t *testing.T) {
	conn := redigomock.NewConn()
	// cmd := conn.Command("BLPOP", "prefix:high:testk1", "prefix:normal:testk1", "prefix:low:testk1", 10)

	//With result
	conn.GenericCommand("BLPOP").Expect([]interface{}{"key", "val"})
	q := NewQueue(conn)
	res, err := q.Get("testk1", 10)
	if err != nil {
		t.Error("Should return without error")
	}
	if res != "val" {
		t.Error("Wrong result: ", res)
	}

	//With error
	conn.GenericCommand("BLPOP").ExpectError(errors.New("Connection error"))
	_, err = q.Get("testk2", 0)
	if err == nil {
		t.Error("Should return with error")
	}
}

func TestGetBytes(t *testing.T) {
	conn := redigomock.NewConn()

	conn.GenericCommand("BLPOP").Expect([]interface{}{"key", []byte("val")})
	q := NewQueue(conn)
	res, err := q.GetBytes("testk1", 10)
	if err != nil {
		t.Error("Should return without error")
	}
	if !reflect.DeepEqual(res, []byte("val")) {
		t.Error("Wrong result", res)
	}

	//Test error
	conn.GenericCommand("BLPOP").ExpectError(errors.New("Connection error"))
	res, err = q.GetBytes("testk1", 10)
	if err == nil {
		t.Error("Should return with error")
	}

	//Test conversion error
	conn.GenericCommand("BLPOP").Expect([]interface{}{"key", 5})
	res, err = q.GetBytes("testk1", 10)
	if err == nil {
		t.Error("Should return with error")
	}

	//Test string conversion
	conn.GenericCommand("BLPOP").Expect([]interface{}{"key", []byte("val")})
	res2, err := q.GetString("testk1", 10)
	if err != nil {
		t.Error("Should return without error")
	}
	if res2 != "val" {
		t.Error("Wrong result", res2)
	}
}

func TestNewQueueWithPrefix(t *testing.T) {

	conn := redigomock.NewConn()
	q := NewQueueWithPrefix(conn, "test")

	if q.Prefix != "test" {
		t.Error("Wrong prefix")
	}
}

func TestNewQueue(t *testing.T) {
	conn := redigomock.NewConn()
	q := NewQueue(conn)

	if q.Prefix != "queue" {
		t.Error("Wrong prefix")
	}
}

func TestPuts(t *testing.T) {
	conn := redigomock.NewConn()
	q := NewQueue(conn)

	//Low
	conn.GenericCommand("RPUSH").Expect(int64(1))
	l, err := q.PutLow("test1", "val1")
	if err != nil {
		t.Error("Should not return error")
	}
	if l != 1 {
		t.Error("Should return 1")
	}

	//Normal
	conn.GenericCommand("RPUSH").Expect(int64(1))
	l, err = q.PutNormal("test1", "val1")
	if err != nil {
		t.Error("Should not return error")
	}
	if l != 1 {
		t.Error("Should return 1")
	}

	//High
	conn.GenericCommand("RPUSH").Expect(int64(1))
	l, err = q.PutHigh("test1", "val1")
	if err != nil {
		t.Error("Should not return error")
	}
	if l != 1 {
		t.Error("Should return 1")
	}

}
