package eventSourcing

import (
	"errors"
	"time"
)

type MysqlEventStore struct {
	DatabaseUrl      string
	Listeners        map[string][]Listener
	Events           map[string][]Event
	eventStore       eventStore
	newEvents        events
	listenersHandler *listenersHandler
}

func (mes *MysqlEventStore) DeclareListener(list Listener, ev Event) {
	mes.Listeners[ev.Name()] = append(mes.Listeners[ev.Name()], list)
}

func (mes *MysqlEventStore) DeclareEvent(ev Event) {
	mes.Events[ev.Name()] = append(mes.Events[ev.Name()], ev)
}

func (mes *MysqlEventStore) validate() error {
	var err error
	if "" == mes.DatabaseUrl {
		err = errors.New("Invalid database url")
	}

	return err
}

func (mes *MysqlEventStore) Init() {
	mes.Listeners = make(map[string][]Listener)
	mes.Events = make(map[string][]Event)
}

func (mes *MysqlEventStore) Load(id string, agg Aggregate) error {
	if err := mes.validate(); err != nil {
		return err
	}
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

func (mes *MysqlEventStore) ApplyNewEvent(e Event) {
	mes.newEvents = append(mes.newEvents, e)
}

func (mes *MysqlEventStore) Save(a Aggregate) error {
	if err := mes.validate(); err != nil {
		return err
	}
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

func (mes *MysqlEventStore) handleListeners(ev Event) error {
	listHand := mes.listenersHandler
	listHand.setListeners(mes.Listeners)
	return listHand.handle(ev)
}
