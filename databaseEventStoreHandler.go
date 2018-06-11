package eventSourcing

import (
	"time"
)

type databaseEventStoreHandler struct {
	listeners        map[string][]Listener
	events           map[string][]Event
	eventStore       eventStoreRepository
	newEvents        events
	listenersHandler *listenersHandler
}

func (mes *databaseEventStoreHandler) declareListener(list Listener, ev Event) {
	mes.listeners[ev.Name()] = append(mes.listeners[ev.Name()], list)
}

func (mes *databaseEventStoreHandler) declareEvent(ev Event) {
	mes.events[ev.Name()] = append(mes.events[ev.Name()], ev)
}

func (mes *databaseEventStoreHandler) init() {
	mes.listeners = make(map[string][]Listener)
	mes.events = make(map[string][]Event)
}

func (mes *databaseEventStoreHandler) load(id string, agg Aggregate) error {
	dm, err := mes.eventStore.load(id)
	if err != nil {
		return err
	}

	if dm == nil {
		return nil
	}

	agg.ReplayEvents(dm.getEvents())

	return nil
}

func (mes *databaseEventStoreHandler) applyNewEvent(e Event) {
	mes.newEvents = append(mes.newEvents, e)
}

func (mes *databaseEventStoreHandler) save(a Aggregate) error {
	a.ReplayEvents(mes.newEvents)
	for _, ev := range mes.newEvents {
		dm := domainMessage{id: a.GetID(), payload: ev, recorderOn: time.Now()}

		if err := mes.eventStore.save(&dm); err != nil {
			return err
		}

		if err := mes.handleListeners(ev); err != nil {
			return err
		}
	}

	mes.newEvents = nil
	return nil
}

func (mes *databaseEventStoreHandler) handleListeners(ev Event) error {
	mes.listenersHandler.listeners = mes.listeners
	return mes.listenersHandler.handle(ev)
}
