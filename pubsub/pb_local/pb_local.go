package pblocal

import (
	"coffee_api/pubsub"
	"context"
	"log"
	"sync"
)

type localPubSub struct {
	messageQueue chan *pubsub.Message
	mapChannel   map[pubsub.Topic][]chan *pubsub.Message
	locker       *sync.RWMutex
}

func NewPubSub() *localPubSub {
	pb := &localPubSub{
		messageQueue: make(chan *pubsub.Message, 10000),
		mapChannel:   make(map[pubsub.Topic][]chan *pubsub.Message),
		locker:       new(sync.RWMutex),
	}

	pb.run()

	return pb
}

func (ps *localPubSub) Publish(ctx context.Context, channel pubsub.Topic, data *pubsub.Message) error {
	data.SetChannel(channel)

	go func() {
		// defer commons.Recover(appCtx)
		ps.messageQueue <- data
		log.Println("New event published:", data.String())
	}()
	return nil
}

func (ps *localPubSub) Subscribe(ctx context.Context, channel pubsub.Topic) (ch <-chan *pubsub.Message, close func()) {
	c := make(chan *pubsub.Message)

	ps.locker.Lock()
	if val, ok := ps.mapChannel[channel]; ok {
		val = append(ps.mapChannel[channel], c)
		ps.mapChannel[channel] = val
	} else {
		ps.mapChannel[channel] = []chan *pubsub.Message{c}
	}
	ps.locker.Unlock()

	return c, func() {
		log.Println("Unsubscribe")

		if chans, ok := ps.mapChannel[channel]; ok {
			for i := range chans {
				if chans[i] == c {
					chans = append(chans[:i], chans[i+1:]...)

					ps.locker.Lock()
					ps.mapChannel[channel] = chans
					ps.locker.Unlock()
					break
				}
			}
		}
	}

}

func (ps *localPubSub) run() error {
	log.Println("Pubsub started")

	return nil
}
