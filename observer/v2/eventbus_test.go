package v2

import (
	"fmt"
	"testing"
	"time"
)

func sub1(msg1, msg2 string) {
	time.Sleep(time.Microsecond * 1)
	fmt.Printf("sub1: %s %s\n", msg1, msg2)
}


func sub2(msg1, msg2 string) {
	fmt.Printf("sub2: %s %s\n", msg1, msg2)
}

func TestAsyncEventBusPublish(t *testing.T) {
	b := NewAsyncEventBus()
	b.Subscribe("topic:1", sub1)
	b.Subscribe("topic:1", sub2)
	b.Publish("topic:1", "test1", "test2")
	b.Publish("topic:1", "testA", "testB")
	time.Sleep(time.Second * 1)
}