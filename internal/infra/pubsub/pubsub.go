package pubsub

import (
	"sync"
)

type pubsub struct {
	mu     sync.RWMutex
	subs   map[string][]chan interface{}
	closed bool
}

func NewPubSub() *pubsub {
	ps := &pubsub{}
	ps.subs = make(map[string][]chan interface{})
	return ps
}

func (ps *pubsub) Subscribe(topic string) <-chan interface{} {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	ch := make(chan interface{}, 1)
	ps.subs[topic] = append(ps.subs[topic], ch)
	return ch
}

func (ps *pubsub) Publish(topic string, data interface{}) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	if ps.closed {
		return
	}

	for _, ch := range ps.subs[topic] {
		go func(ch chan interface{}) {
			ch <- data
		}(ch)
	}
}
