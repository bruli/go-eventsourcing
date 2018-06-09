package eventSourcing

import (
	"time"
)

type mysqlEventStoreHandler struct {
	listeners        map[string][]Listener
	events           map[string][]Event
	eventStore       eventStoreRepository
	newEvents        events
	listenersHandler *listenersHandler
}

func (mes *mysqlEventStoreHandler) declareListener(list Listener, ev Event) {
	mes.listeners[ev.Name()] = append(mes.listeners[ev.Name()], list)
}

func (mes *mysqlEventStoreHandler) declareEvent(ev Event) {
	mes.events[ev.Name()] = append(mes.events[ev.Name()], ev)
}

func (mes *mysqlEventStoreHandler) init() {
	mes.listeners = make(map[string][]Listener)
	mes.events = make(map[string][]Event)
}

func (mes *mysqlEventStoreHandler) load(id string, agg Aggregate) error {
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

func (mes *mysqlEventStoreHandler) applyNewEvent(e Event) {
	mes.newEvents = append(mes.newEvents, e)
}

func (mes *mysqlEventStoreHandler) save(a Aggregate) error {
	for _, ev := range mes.newEvents {
		dm := domainMessage{id: a.GetID(), payload: ev, recorderOn: time.Now()}

		if err := mes.eventStore.save(&dm); err != nil {
			return err
		}

		if err := mes.handleListeners(ev); err != nil {
			return err
		}
	}

	return nil
}

func (mes *mysqlEventStoreHandler) handleListeners(ev Event) error {
	listHand := mes.listenersHandler
	listHand.setListeners(mes.listeners)
	return listHand.handle(ev)
}
