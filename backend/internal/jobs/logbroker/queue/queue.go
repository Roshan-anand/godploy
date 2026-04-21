package logbrokerqueue

import "github.com/google/uuid"

type PubData struct {
	ID  uuid.UUID
	Msg string
}

type LogBrokerQueue struct {
	Pub chan *PubData
}

// initializes the log broker queue
func InitLogBrokerQueue() *LogBrokerQueue {
	pub := make(chan *PubData, 10)

	return &LogBrokerQueue{
		Pub: pub,
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
