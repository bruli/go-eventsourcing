package eventSourcing

import (
	"errors"
	"github.com/manveru/faker"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDatabaseEventStore(t *testing.T) {
	t.Run("it should return error when load return error", func(t *testing.T) {
		eventSt := &databaseEventStoreRepositoryMock{}
		eventSt.setEventsFunc = func(ev map[string]Event) {
		}
		eventSt.loadFunc = func(ID string) (*domainMessages, error) {
			return nil, errors.New("error")
		}
		listHand := listenersHandler{}

		databaseES := databaseEventStoreHandler{eventStore: eventSt, listenersHandler: &listHand}
		agg := &AggregateMock{}
		assert.Error(t, databaseES.load(uuid.NewV4().String(), agg))

	})

	t.Run("it should return error nil without domain messages", func(t *testing.T) {
		eventSt := &databaseEventStoreRepositoryMock{}
		eventSt.setEventsFunc = func(ev map[string]Event) {
		}
		eventSt.loadFunc = func(ID string) (*domainMessages, error) {
			return nil, nil
		}
		listHand := listenersHandler{}

		databaseES := databaseEventStoreHandler{eventStore: eventSt, listenersHandler: &listHand}
		agg := &AggregateMock{}
		assert.Nil(t, databaseES.load(uuid.NewV4().String(), agg))

	})
	t.Run("it should replay events in load", func(t *testing.T) {
		eventSt := &databaseEventStoreRepositoryMock{}
		eventSt.setEventsFunc = func(ev map[string]Event) {
		}
		eventSt.loadFunc = func(ID string) (*domainMessages, error) {
			dms := domainMessagesStub()
			return &dms, nil
		}
		listHand := listenersHandler{}

		databaseES := databaseEventStoreHandler{eventStore: eventSt, listenersHandler: &listHand}
		agg := &AggregateMock{}
		agg.ReplayEventsFunc = func(e []Event) {
		}
		assert.Nil(t, databaseES.load(uuid.NewV4().String(), agg))

	})

	t.Run("it should return nil without events", func(t *testing.T) {
		eventSt := &databaseEventStoreRepositoryMock{}
		eventSt.setEventsFunc = func(ev map[string]Event) {
		}
		eventSt.saveFunc = func(message *domainMessage) error {
			return errors.New("error")
		}
		listHand := listenersHandler{}

		databaseES := databaseEventStoreHandler{eventStore: eventSt, listenersHandler: &listHand}
		agg := &AggregateMock{}
		agg.ReplayEventsFunc = func(e []Event) {
		}
		assert.Nil(t, databaseES.save(agg))

	})
	t.Run("it should return error when save returns error", func(t *testing.T) {
		ev := EventMock{}
		eventSt := &databaseEventStoreRepositoryMock{}
		eventSt.saveFunc = func(message *domainMessage) error {
			return errors.New("error")
		}
		listHand := listenersHandler{}

		databaseES := databaseEventStoreHandler{eventStore: eventSt, listenersHandler: &listHand}
		agg := &AggregateMock{}
		agg.GetIDFunc = func() string {
			return uuid.NewV4().String()
		}
		agg.ReplayEventsFunc = func(e []Event) {
		}

		databaseES.applyNewEvent(&ev)
		assert.Error(t, databaseES.save(agg))

	})
	t.Run("it should return error when listeners returns error", func(t *testing.T) {
		ev := &EventMock{}
		ev.NameFunc = func() string {
			return "eventito"
		}
		list1 := ListenerMock{
			HandleFunc: func(event Event) error {
				return nil
			},
		}
		list2 := ListenerMock{
			HandleFunc: func(event Event) error {
				return errors.New("error")
			},
		}
		eventSt := &databaseEventStoreRepositoryMock{}
		eventSt.saveFunc = func(message *domainMessage) error {
			return nil
		}
		listHand := listenersHandler{}

		databaseES := databaseEventStoreHandler{eventStore: eventSt, listenersHandler: &listHand}
		databaseES.init()
		databaseES.declareListener(&list1, ev)
		databaseES.declareListener(&list2, ev)
		databaseES.declareEvent(ev)
		agg := &AggregateMock{}
		agg.GetIDFunc = func() string {
			return uuid.NewV4().String()
		}
		agg.ReplayEventsFunc = func(e []Event) {
		}

		databaseES.applyNewEvent(ev)
		assert.Error(t, databaseES.save(agg))

	})
	t.Run("it should save new event and call listeners", func(t *testing.T) {
		ev := &EventMock{}
		ev.NameFunc = func() string {
			return "eventito"
		}
		list1 := ListenerMock{
			HandleFunc: func(event Event) error {
				return nil
			},
		}
		list2 := ListenerMock{
			HandleFunc: func(event Event) error {
				return nil
			},
		}
		eventSt := &databaseEventStoreRepositoryMock{}
		eventSt.saveFunc = func(message *domainMessage) error {
			return nil
		}
		listHand := listenersHandler{}

		databaseES := databaseEventStoreHandler{eventStore: eventSt, listenersHandler: &listHand}
		databaseES.init()
		databaseES.declareListener(&list1, ev)
		databaseES.declareListener(&list2, ev)
		databaseES.declareEvent(ev)
		agg := &AggregateMock{}
		agg.GetIDFunc = func() string {
			return uuid.NewV4().String()
		}
		agg.ReplayEventsFunc = func(e []Event) {
		}

		databaseES.applyNewEvent(ev)
		assert.Nil(t, databaseES.save(agg))

	})
	t.Run("it should save two events", func(t *testing.T) {
		ev := &EventMock{}
		ev.NameFunc = func() string {
			return "eventito"
		}
		event2 := &EventMock{}
		event2.NameFunc = func() string {
			return "other.eventito"
		}
		list1 := ListenerMock{
			HandleFunc: func(event Event) error {
				return nil
			},
		}
		list2 := ListenerMock{
			HandleFunc: func(event Event) error {
				return nil
			},
		}
		list3 := ListenerMock{
			HandleFunc: func(event Event) error {
				return nil
			},
		}
		eventSt := &databaseEventStoreRepositoryMock{}
		eventSt.saveFunc = func(message *domainMessage) error {
			return nil
		}
		listHand := listenersHandler{}

		databaseES := databaseEventStoreHandler{eventStore: eventSt, listenersHandler: &listHand}
		databaseES.init()
		databaseES.declareListener(&list1, ev)
		databaseES.declareListener(&list2, ev)
		databaseES.declareEvent(ev)
		agg := &AggregateMock{}
		agg.GetIDFunc = func() string {
			return uuid.NewV4().String()
		}
		agg.ReplayEventsFunc = func(e []Event) {
		}

		databaseES.applyNewEvent(ev)
		assert.Nil(t, databaseES.save(agg))
		assert.Equal(t, 1, len(list1.HandleCalls()))
		assert.Equal(t, 1, len(list2.HandleCalls()))

		databaseES.declareEvent(event2)
		databaseES.declareListener(&list3, event2)
		databaseES.applyNewEvent(event2)
		assert.Nil(t, databaseES.save(agg))
		assert.Equal(t, 2, len(eventSt.saveCalls()))
		assert.Equal(t, 1, len(list3.HandleCalls()))
	})
}

func domainMessagesStub() domainMessages {
	var dms domainMessages
	for i := 0; 2 > i; i++ {
		dm := domainMessageStub()
		dms = append(dms, &dm)
	}

	return dms
}
func domainMessageStub() domainMessage {
	ev := EventMock{
		NameFunc: func() string {
			return getFaker().UserName()
		},
		PayloadFunc: func() []byte {
			return []byte(`{"id": "1111", "age": 43}`)
		},
	}
	return domainMessage{id: uuid.NewV4().String(), recorderOn: time.Now(), payload: &ev}
}

func getFaker() *faker.Faker {
	f, _ := faker.New("en")
	return f
}
