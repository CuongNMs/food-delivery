package localpubsub

import (
	"context"
	"food-delivery/common"
	"food-delivery/pubsub"
	"log"
	"sync"
)

type localPubsub struct {
	messageQueue chan *pubsub.Message
	mapChannel   map[pubsub.Topic][]chan *pubsub.Message
	locker       *sync.RWMutex
}

func NewPubsub() *localPubsub {
	pb := &localPubsub{
		messageQueue: make(chan *pubsub.Message, 10000),
		mapChannel:   make(map[pubsub.Topic][]chan *pubsub.Message),
		locker:       new(sync.RWMutex),
	}
	pb.run()
	return pb
}

func (l *localPubsub) Publish(ctx context.Context, topic pubsub.Topic, data *pubsub.Message) error {
	data.SetChannel(topic)
	go func() {
		defer common.AppRecover()
		l.messageQueue <- data
		log.Println("New event published: ", data.String(), "with data:", data.Data())
	}()
	return nil
}

func (l *localPubsub) Subscribe(ctx context.Context, topic pubsub.Topic) (ch <-chan *pubsub.Message, close func()) {
	c := make(chan *pubsub.Message)
	l.locker.Lock()
	if val, ok := l.mapChannel[topic]; ok {
		val = append(l.mapChannel[topic], c)
		l.mapChannel[topic] = val
	} else {
		l.mapChannel[topic] = []chan *pubsub.Message{c}
	}
	l.locker.Unlock()

	return c, func() {
		log.Println("Unsubscribe")

		if chans, ok := l.mapChannel[topic]; ok {
			for i := range chans {
				if chans[i] == c {
					chans = append(chans[:i], chans[:i+1]...)
					l.locker.Lock()
					l.mapChannel[topic] = chans
					l.locker.Unlock()
					break
				}
			}
		}
	}

}

func (l *localPubsub) run() error {
	log.Println("Start pubsub")
	go func() {
		for {
			mess := <-l.messageQueue
			log.Println("Message dequeue: ", mess)
			if subs, ok := l.mapChannel[mess.Channel()]; ok {
				for i := range subs {
					go func(c chan *pubsub.Message) {
						c <- mess
					}(subs[i])
				}
			}
		}
	}()
	return nil
}
