package eventSourcing

type inMemoryEventStoreHandler struct {
	listeners        map[string][]Listener
	events           map[string]Event
	newEvents        events
	listenersHandler *listenersHandler
}

func (imes *inMemoryEventStoreHandler) declareListener(list Listener, ev Event) {
	imes.listeners[ev.Name()] = append(imes.listeners[ev.Name()], list)
}

func (imes *inMemoryEventStoreHandler) declareEvent(ev Event) {
	imes.events[ev.Name()] = ev
}

func (imes *inMemoryEventStoreHandler) init() {
	imes.listeners = make(map[string][]Listener)
	imes.events = make(map[string]Event)
}

func (imes *inMemoryEventStoreHandler) applyNewEvent(e Event) {
	imes.newEvents = append(imes.newEvents, e)
}

func (imes *inMemoryEventStoreHandler) save(a Aggregate) error {
	a.ReplayEvents(imes.newEvents)
	for _, ev := range imes.newEvents {
		if err := imes.handleListeners(ev); err != nil {
			return err
		}
	}

	imes.newEvents = nil
	return nil
}

func (imes *inMemoryEventStoreHandler) handleListeners(ev Event) error {
	imes.listenersHandler.listeners = imes.listeners
	return imes.listenersHandler.handle(ev)
}
