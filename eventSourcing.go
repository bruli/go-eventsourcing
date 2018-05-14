package eventSourcing

import (
	"time"
)

type EventSourcing struct {
	EventStore eventStoreI
}

func (es *EventSourcing) ApplyNewEvent(e Event) {
	ev := &currentEvents
	ev.events = append(ev.events, e)
}

func (es *EventSourcing) Save(a Aggregate) error {
	for _, r := range currentEvents.events {
		d := domainMessage{id: a.GetID(), payload: r, recorderOn: time.Now()}
		esr := es.EventStore
		if err := esr.save(&d); err != nil {
			return err
		}

		if err := handle(r); err != nil {
			return err
		}
	}

	return nil
}

func (es *EventSourcing) DeclareListener(listener Listener, event Event) {
	eventBus.addListener(listener, event)
}

func (es *EventSourcing) DeclareEvent(e Event) {
	eventBus.addEvent(e)
}

func (es *EventSourcing) Load(id string, ag Aggregate) error {
	esr := es.EventStore
	dm, err := esr.load(id)
	if err != nil {
		return nil
	}

	if 0 == len(dm.messages) {
		return err
	}

	ag.ReplayEvents(dm.getEvents())

	return nil
}
