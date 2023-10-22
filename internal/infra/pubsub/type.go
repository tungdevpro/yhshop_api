package pubsub

type PubSub interface {
	Subscribe(topic string) <-chan interface{}
	Publish(topic string, data interface{})
	Close()
}
