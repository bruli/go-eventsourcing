package eventSourcing

type InMemoryEventStore struct {
}

func (imes *InMemoryEventStore) Init() {
	container.handler.inMemoryEventStoreHandler.init()
}

func (imes *InMemoryEventStore) DeclareListener(list Listener, ev Event) {
	container.handler.inMemoryEventStoreHandler.declareListener(list, ev)
}

func (imes *InMemoryEventStore) DeclareEvent(ev Event) {
	container.handler.inMemoryEventStoreHandler.declareEvent(ev)
}

func (imes *InMemoryEventStore) ApplyNewEvent(ev Event) {
	container.handler.inMemoryEventStoreHandler.applyNewEvent(ev)
}

func (imes *InMemoryEventStore) Save(a Aggregate) error {
	return container.handler.inMemoryEventStoreHandler.save(a)
}
