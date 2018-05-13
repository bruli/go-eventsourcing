package eventSourcing

type Aggregate interface {
	GetID() string
	ReplayEvents(e []Event)
}
