package eventSourcing

import (
	"errors"
	"github.com/manveru/faker"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestMysqlEventStore(t *testing.T) {
	t.Run("it should return error when load return error", func(t *testing.T) {
		eventSt := &eventStoreRepositoryMock{}
		eventSt.loadFunc = func(ID string) (*domainMessages, error) {
			return nil, errors.New("error")
		}
		listHand := listenersHandler{}

		mysqlES := mysqlEventStoreHandler{eventStore: eventSt, listenersHandler: &listHand}
		agg := &AggregateMock{}
		assert.Error(t, mysqlES.load(uuid.NewV4().String(), agg))

	})

	t.Run("it should return error nil without domain messages", func(t *testing.T) {
		eventSt := &eventStoreRepositoryMock{}
		eventSt.loadFunc = func(ID string) (*domainMessages, error) {
			return nil, nil
		}
		listHand := listenersHandler{}

		mysqlES := mysqlEventStoreHandler{eventStore: eventSt, listenersHandler: &listHand}
		agg := &AggregateMock{}
		assert.Nil(t, mysqlES.load(uuid.NewV4().String(), agg))

	})
	t.Run("it should replay events in load", func(t *testing.T) {
		eventSt := &eventStoreRepositoryMock{}
		eventSt.loadFunc = func(ID string) (*domainMessages, error) {
			dms := domainMessagesStub()
			return &dms, nil
		}
		listHand := listenersHandler{}

		mysqlES := mysqlEventStoreHandler{eventStore: eventSt, listenersHandler: &listHand}
		agg := &AggregateMock{}
		agg.ReplayEventsFunc = func(e []Event) {
		}
		assert.Nil(t, mysqlES.load(uuid.NewV4().String(), agg))

	})

	t.Run("it should return nil without events", func(t *testing.T) {
		eventSt := &eventStoreRepositoryMock{}
		eventSt.saveFunc = func(message *domainMessage) error {
			return errors.New("error")
		}
		listHand := listenersHandler{}

		mysqlES := mysqlEventStoreHandler{eventStore: eventSt, listenersHandler: &listHand}
		agg := &AggregateMock{}
		assert.Nil(t, mysqlES.save(agg))

	})
	t.Run("it should return error when save returns error", func(t *testing.T) {
		ev := EventMock{}
		eventSt := &eventStoreRepositoryMock{}
		eventSt.saveFunc = func(message *domainMessage) error {
			return errors.New("error")
		}
		listHand := listenersHandler{}

		mysqlES := mysqlEventStoreHandler{eventStore: eventSt, listenersHandler: &listHand}
		agg := &AggregateMock{}
		agg.GetIDFunc = func() string {
			return uuid.NewV4().String()
		}

		mysqlES.applyNewEvent(&ev)
		assert.Error(t, mysqlES.save(agg))

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
		eventSt := &eventStoreRepositoryMock{}
		eventSt.saveFunc = func(message *domainMessage) error {
			return nil
		}
		listHand := listenersHandler{}

		mysqlES := mysqlEventStoreHandler{eventStore: eventSt, listenersHandler: &listHand}
		mysqlES.init()
		mysqlES.declareListener(&list1, ev)
		mysqlES.declareListener(&list2, ev)
		mysqlES.declareEvent(ev)
		agg := &AggregateMock{}
		agg.GetIDFunc = func() string {
			return uuid.NewV4().String()
		}

		mysqlES.applyNewEvent(ev)
		assert.Error(t, mysqlES.save(agg))

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
		eventSt := &eventStoreRepositoryMock{}
		eventSt.saveFunc = func(message *domainMessage) error {
			return nil
		}
		listHand := listenersHandler{}

		mysqlES := mysqlEventStoreHandler{eventStore: eventSt, listenersHandler: &listHand}
		mysqlES.init()
		mysqlES.declareListener(&list1, ev)
		mysqlES.declareListener(&list2, ev)
		mysqlES.declareEvent(ev)
		agg := &AggregateMock{}
		agg.GetIDFunc = func() string {
			return uuid.NewV4().String()
		}

		mysqlES.applyNewEvent(ev)
		assert.Nil(t, mysqlES.save(agg))

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
