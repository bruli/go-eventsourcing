package eventSourcing

import (
	"errors"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInMemoryEventStoreHandler(t *testing.T) {
	t.Run("it should return error when a listener returns error", func(t *testing.T) {
		ev := EventMock{
			NameFunc: func() string {
				return "event_one"
			},
		}

		agg := AggregateMock{
			GetIDFunc: func() string {
				return uuid.NewV4().String()
			},
			ReplayEventsFunc: func(e []Event) {
			},
		}
		list := ListenerMock{
			HandleFunc: func(event Event) error {
				return errors.New("error")
			},
		}
		listHand := listenersHandler{}
		hand := inMemoryEventStoreHandler{listenersHandler: &listHand}
		hand.init()
		hand.declareEvent(&ev)
		hand.declareListener(&list, &ev)
		hand.applyNewEvent(&ev)
		assert.Error(t, hand.save(&agg))
	})
	t.Run("it should dispatch event", func(t *testing.T) {
		ev := EventMock{
			NameFunc: func() string {
				return "event_one"
			},
		}

		agg := AggregateMock{
			GetIDFunc: func() string {
				return uuid.NewV4().String()
			},
			ReplayEventsFunc: func(e []Event) {
			},
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
		listHand := listenersHandler{}
		hand := inMemoryEventStoreHandler{listenersHandler: &listHand}
		hand.init()
		hand.declareEvent(&ev)
		hand.declareListener(&list1, &ev)
		hand.declareListener(&list2, &ev)
		hand.applyNewEvent(&ev)
		assert.NoError(t, hand.save(&agg))
		assert.Equal(t, 1, len(list1.HandleCalls()))
		assert.Equal(t, 1, len(list2.HandleCalls()))
	})
}
