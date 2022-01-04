package v2

import (
	"fmt"
	"reflect"
	"sync"
)

type Bus interface {
	Subscribe(topic string, handler interface{}) error
	Publish(topic string, args ...interface{})
}

type AsyncEventBus struct {
	handlers	map[string][]reflect.Value
	lock		sync.Mutex
}

func NewAsyncEventBus() *AsyncEventBus {
	return &AsyncEventBus {
		handlers:	map[string][]reflect.Value{},
		lock:		sync.Mutex{},
	}
}

func (b *AsyncEventBus) Subscribe(topic string, fn interface{}) error {
	b.lock.Lock()
	defer b.lock.Unlock()

	v := reflect.ValueOf(fn)
	if v.Type().Kind() != reflect.Func {
		return fmt.Errorf("fn is not a func")
	}

	h, ok := b.handlers[topic]
	if !ok {
		h = []reflect.Value{}
	}
	h = append(h, v)
	b.handlers[topic] = h

	return nil
}

//Publish 异步执行，并且不会等待返回结果
func (b *AsyncEventBus) Publish(topic string, args ...interface{}) {
	handlers, ok := b.handlers[topic]
	if !ok {
		fmt.Println("not found handler in topic:", topic)
		return
	}
	params := make([]reflect.Value, len(args))
	for i, arg := range args {
		params[i] = reflect.ValueOf(arg)
	}

	for i := range handlers {
		go handlers[i].Call(params)
	}
}