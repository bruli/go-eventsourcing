package eventSourcing

import (
	"time"
)

func ApplyNewEvent(e Event) {
	ev := &currentEvents
	ev.events = append(ev.events, e)
}

func Save(a Aggregate) error {
	for _, r := range currentEvents.events {
		d := domainMessage{a.GetID(), r, time.Now()}

		esr := container.infrastructure.eventRepository
		err := esr.save(&d)

		if err != nil {
			return err
		}

		err = handle(r)

		if err != nil {
			return err
		}
	}

	return nil
}

func DeclareListener(listener Listener, event Event) {
	stdEventBus.addListener(listener, event)
}

func DeclareEvent(e Event) {
	stdEventBus.addEvent(e)
}

func Load(id string, ag Aggregate) error {
	esr := container.infrastructure.eventRepository
	dm, err := esr.load(id)

	if 0 == len(dm) {
		return err
	}

	ag.ReplayEvents(getEvents(dm))

	return err
}

func getEvents(messages []*domainMessage) []Event {
	var ev []Event

	for _, m := range messages {
		e := append(ev, m.payload)
		ev = e
	}

	return ev

}
