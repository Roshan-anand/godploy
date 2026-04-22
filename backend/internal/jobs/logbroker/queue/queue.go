package logbrokerqueue

import (
	"net/http"

	"github.com/google/uuid"
)

type PubData struct {
	ID  uuid.UUID
	Msg string
}

type Subscriber struct {
	w            http.ResponseWriter
	DeploymentID uuid.UUID
}

type LogBrokerQueue struct {
	Pub         chan *PubData
	Subscribers map[uuid.UUID]*Subscriber
}

// initializes the log broker queue
func InitLogBrokerQueue() *LogBrokerQueue {
	pub := make(chan *PubData, 10)
	sub := make(map[uuid.UUID]*Subscriber, 10)

	return &LogBrokerQueue{
		Pub:         pub,
		Subscribers: sub,
	}
}

// closes the publisher channel
func (l *LogBrokerQueue) Close() {
	close(l.Pub)
}

// push log data to the publisher channel
func (l *LogBrokerQueue) PublishLog(data *PubData) {
	l.Pub <- data
}

// subscribe to logs of the deployment
func (l *LogBrokerQueue) SubscribeLogs(id uuid.UUID, sub *Subscriber) {
	l.Subscribers[id] = sub
}

// unsubscribe user from logs of the deployment
func (l *LogBrokerQueue) UnsubscribeLogs(id uuid.UUID) {
	delete(l.Subscribers, id)
}
