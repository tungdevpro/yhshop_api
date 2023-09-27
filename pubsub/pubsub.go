package pubsub

import "context"

type Topic string

type Pubsub interface {
	Publish(context.Context, Topic, *Message) error
	Subscribe(context.Context, Topic) (ch <-chan *Message, close func())
}
