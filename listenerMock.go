// Code generated by moq; DO NOT EDIT
// github.com/matryer/moq

package eventSourcing

import (
	"sync"
)

var (
	lockListenerMockHandle sync.RWMutex
)

// ListenerMock is a mock implementation of Listener.
//
//     func TestSomethingThatUsesListener(t *testing.T) {
//
//         // make and configure a mocked Listener
//         mockedListener := &ListenerMock{
//             HandleFunc: func(event Event) error {
// 	               panic("TODO: mock out the Handle method")
//             },
//         }
//
//         // TODO: use mockedListener in code that requires Listener
//         //       and then make assertions.
//
//     }
type ListenerMock struct {
	// HandleFunc mocks the Handle method.
	HandleFunc func(event Event) error

	// calls tracks calls to the methods.
	calls struct {
		// Handle holds details about calls to the Handle method.
		Handle []struct {
			// Event is the event argument value.
			Event Event
		}
	}
}

// Handle calls HandleFunc.
func (mock *ListenerMock) Handle(event Event) error {
	if mock.HandleFunc == nil {
		panic("moq: ListenerMock.HandleFunc is nil but Listener.Handle was just called")
	}
	callInfo := struct {
		Event Event
	}{
		Event: event,
	}
	lockListenerMockHandle.Lock()
	mock.calls.Handle = append(mock.calls.Handle, callInfo)
	lockListenerMockHandle.Unlock()
	return mock.HandleFunc(event)
}

// HandleCalls gets all the calls that were made to Handle.
// Check the length with:
//     len(mockedListener.HandleCalls())
func (mock *ListenerMock) HandleCalls() []struct {
	Event Event
} {
	var calls []struct {
		Event Event
	}
	lockListenerMockHandle.RLock()
	calls = mock.calls.Handle
	lockListenerMockHandle.RUnlock()
	return calls
}