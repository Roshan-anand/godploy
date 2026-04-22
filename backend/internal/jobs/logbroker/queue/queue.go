package logbrokerqueue

import (
	"github.com/Roshan-anand/godploy/internal/lib/sse"
	"github.com/google/uuid"
)

type PubData struct {
	ID  uuid.UUID
	Msg string
}

type Subscriber struct {
	SSE          *sse.SSE
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
func (l *LogBrokerQueue) SubscribeLogs(userID uuid.UUID, sub *Subscriber) {
	l.Subscribers[userID] = sub
}

// unsubscribe user from logs of the deployment
func (l *LogBrokerQueue) UnsubscribeLogs(userID uuid.UUID) {
	delete(l.Subscribers, userID)
}
