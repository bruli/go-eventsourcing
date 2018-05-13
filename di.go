package eventSourcing

var container eventSourcingContainer

type eventSourcingContainer struct {
	infrastructure infrastructure
}

type infrastructure struct {
	eventRepository eventStore
}

func init() {
	container.infrastructure.eventRepository = &eventStoreRepository{}
}
