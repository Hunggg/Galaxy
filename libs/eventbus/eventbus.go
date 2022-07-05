package eventbus

import (
	"math/rand"
	"strconv"
	"sync"

	"github.com/cskr/pubsub"
)

const DefaultBuffered = 50

type Subscription struct {
	Id       string
	Topic    string
	Receiver chan interface{}
	ErrCh    chan error
}

type SubscriptionKeeper struct {
	sync.Mutex
	ps             *pubsub.PubSub
	bufferedLength int
	subscriptions  map[string]Subscription
}

func NewSubscriptionKeeper(bufferedLen int) *SubscriptionKeeper {
	return &SubscriptionKeeper{
		bufferedLength: bufferedLen,
		ps:             pubsub.New(bufferedLen),
		subscriptions:  make(map[string]Subscription),
	}
}

// Subscribe subscribe on topic, use Receiver channel to receive messages
func (s *SubscriptionKeeper) Subscribe(topic string, errCh chan error) Subscription {
	s.Lock()
	defer s.Unlock()
	sub := s.ps.Sub(topic)

	id := generateRandomId()
	subscription := Subscription{
		Id:       id,
		Topic:    topic,
		Receiver: sub,
		ErrCh:    errCh,
	}

	s.subscriptions[id] = subscription
	return subscription
}

// Unsubscribe unsubscribe using subscription id
func (s *SubscriptionKeeper) Unsubscribe(id string) {
	s.Lock()
	defer s.Unlock()
	if _, ok := s.subscriptions[id]; ok {
		s.ps.Unsub(s.subscriptions[id].Receiver)
		delete(s.subscriptions, id)
	}
}

// Publish Publish with specific topic, should not block
func (s *SubscriptionKeeper) Publish(topic string, message interface{}) {
	s.ps.Pub(message, topic)
}

func (s *SubscriptionKeeper) GetAllSubsciptions() map[string]Subscription {
	s.Lock()
	defer s.Unlock()
	var result = make(map[string]Subscription)
	for key, item := range s.subscriptions {
		result[key] = item
	}
	return result
}

func (s *SubscriptionKeeper) GetAllTopics() []string {
	s.Lock()
	defer s.Unlock()
	var result = make([]string, 0)
	for _, sub := range s.subscriptions {
		if !isInList(sub.Topic, result) {
			result = append(result, sub.Topic)
		}
	}
	return result
}

func generateRandomId() string {
	id := rand.Int()
	return strconv.Itoa(id)
}

func isInList(item string, arr []string) bool {
	for i := range arr {
		if item == arr[i] {
			return true
		}
	}
	return false
}
