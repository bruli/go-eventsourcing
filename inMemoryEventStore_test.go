package eventSourcing

import (
	"errors"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInMemoryEventStore(t *testing.T) {
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
		hand := InMemoryEventStore{}
		hand.Init()
		hand.DeclareEvent(&ev)
		hand.DeclareListener(&list, &ev)
		hand.ApplyNewEvent(&ev)
		assert.Error(t, hand.Save(&agg))
		assert.Equal(t, 1, len(list.HandleCalls()))
	})
	t.Run("it should dispatch event", func(t *testing.T) {
		ev := EventMock{
			NameFunc: func() string {
				return "event_two"
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
		hand := InMemoryEventStore{}
		hand.Init()
		hand.DeclareEvent(&ev)
		hand.DeclareListener(&list1, &ev)
		hand.DeclareListener(&list2, &ev)
		hand.ApplyNewEvent(&ev)
		assert.NoError(t, hand.Save(&agg))
		assert.Equal(t, 1, len(list1.HandleCalls()))
		assert.Equal(t, 1, len(list2.HandleCalls()))
	})
}
