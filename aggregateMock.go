// Code generated by moq; DO NOT EDIT
// github.com/matryer/moq

package eventSourcing

import (
	"sync"
)

var (
	lockAggregateMockGetID        sync.RWMutex
	lockAggregateMockReplayEvents sync.RWMutex
)

// AggregateMock is a mock implementation of Aggregate.
//
//     func TestSomethingThatUsesAggregate(t *testing.T) {
//
//         // make and configure a mocked Aggregate
//         mockedAggregate := &AggregateMock{
//             GetIDFunc: func() string {
// 	               panic("TODO: mock out the GetID method")
//             },
//             ReplayEventsFunc: func(e []Event)  {
// 	               panic("TODO: mock out the ReplayEvents method")
//             },
//         }
//
//         // TODO: use mockedAggregate in code that requires Aggregate
//         //       and then make assertions.
//
//     }
type AggregateMock struct {
	// GetIDFunc mocks the GetID method.
	GetIDFunc func() string

	// ReplayEventsFunc mocks the ReplayEvents method.
	ReplayEventsFunc func(e []Event)

	// calls tracks calls to the methods.
	calls struct {
		// GetID holds details about calls to the GetID method.
		GetID []struct {
		}
		// ReplayEvents holds details about calls to the ReplayEvents method.
		ReplayEvents []struct {
			// E is the e argument value.
			E []Event
		}
	}
}

// GetID calls GetIDFunc.
func (mock *AggregateMock) GetID() string {
	if mock.GetIDFunc == nil {
		panic("moq: AggregateMock.GetIDFunc is nil but Aggregate.GetID was just called")
	}
	callInfo := struct {
	}{}
	lockAggregateMockGetID.Lock()
	mock.calls.GetID = append(mock.calls.GetID, callInfo)
	lockAggregateMockGetID.Unlock()
	return mock.GetIDFunc()
}

// GetIDCalls gets all the calls that were made to GetID.
// Check the length with:
//     len(mockedAggregate.GetIDCalls())
func (mock *AggregateMock) GetIDCalls() []struct {
} {
	var calls []struct {
	}
	lockAggregateMockGetID.RLock()
	calls = mock.calls.GetID
	lockAggregateMockGetID.RUnlock()
	return calls
}

// ReplayEvents calls ReplayEventsFunc.
func (mock *AggregateMock) ReplayEvents(e []Event) {
	if mock.ReplayEventsFunc == nil {
		panic("moq: AggregateMock.ReplayEventsFunc is nil but Aggregate.ReplayEvents was just called")
	}
	callInfo := struct {
		E []Event
	}{
		E: e,
	}
	lockAggregateMockReplayEvents.Lock()
	mock.calls.ReplayEvents = append(mock.calls.ReplayEvents, callInfo)
	lockAggregateMockReplayEvents.Unlock()
	mock.ReplayEventsFunc(e)
}

// ReplayEventsCalls gets all the calls that were made to ReplayEvents.
// Check the length with:
//     len(mockedAggregate.ReplayEventsCalls())
func (mock *AggregateMock) ReplayEventsCalls() []struct {
	E []Event
} {
	var calls []struct {
		E []Event
	}
	lockAggregateMockReplayEvents.RLock()
	calls = mock.calls.ReplayEvents
	lockAggregateMockReplayEvents.RUnlock()
	return calls
}