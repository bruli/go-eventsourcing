package eventSourcing

import "time"

//go:generate moq -out aggregateMock.go . Aggregate
type Aggregate interface {
	GetID() string
	ReplayEvents(e []Event)
}

//go:generate moq -out eventMock.go . Event
type Event interface {
	Name() string
	Payload() []byte
}

//go:generate moq -out listenerMock.go . Listener
type Listener interface {
	Handle(event Event) error
}

type domainMessages []*domainMessage

func (dm domainMessages) getEvents() []Event {
	var ev []Event

	for _, m := range dm {
		ev = append(ev, m.payload)
	}

	return ev
}

type domainMessage struct {
	id         string
	payload    Event
	recorderOn time.Time
}

type events []Event
