// Code generated by moq; DO NOT EDIT
// github.com/matryer/moq

package eventSourcing

import (
	"sync"
)

var (
	lockdatabaseEventStoreRepositoryMockload      sync.RWMutex
	lockdatabaseEventStoreRepositoryMocksave      sync.RWMutex
	lockdatabaseEventStoreRepositoryMocksetEvents sync.RWMutex
)

// databaseEventStoreRepositoryMock is a mock implementation of databaseEventStoreRepository.
//
//     func TestSomethingThatUsesdatabaseEventStoreRepository(t *testing.T) {
//
//         // make and configure a mocked databaseEventStoreRepository
//         mockeddatabaseEventStoreRepository := &databaseEventStoreRepositoryMock{
//             loadFunc: func(ID string) (*domainMessages, error) {
// 	               panic("TODO: mock out the load method")
//             },
//             saveFunc: func(message *domainMessage) error {
// 	               panic("TODO: mock out the save method")
//             },
//             setEventsFunc: func(ev map[string]Event)  {
// 	               panic("TODO: mock out the setEvents method")
//             },
//         }
//
//         // TODO: use mockeddatabaseEventStoreRepository in code that requires databaseEventStoreRepository
//         //       and then make assertions.
//
//     }
type databaseEventStoreRepositoryMock struct {
	// loadFunc mocks the load method.
	loadFunc func(ID string) (*domainMessages, error)

	// saveFunc mocks the save method.
	saveFunc func(message *domainMessage) error

	// setEventsFunc mocks the setEvents method.
	setEventsFunc func(ev map[string]Event)

	// calls tracks calls to the methods.
	calls struct {
		// load holds details about calls to the load method.
		load []struct {
			// ID is the ID argument value.
			ID string
		}
		// save holds details about calls to the save method.
		save []struct {
			// Message is the message argument value.
			Message *domainMessage
		}
		// setEvents holds details about calls to the setEvents method.
		setEvents []struct {
			// Ev is the ev argument value.
			Ev map[string]Event
		}
	}
}

// load calls loadFunc.
func (mock *databaseEventStoreRepositoryMock) load(ID string) (*domainMessages, error) {
	if mock.loadFunc == nil {
		panic("moq: databaseEventStoreRepositoryMock.loadFunc is nil but databaseEventStoreRepository.load was just called")
	}
	callInfo := struct {
		ID string
	}{
		ID: ID,
	}
	lockdatabaseEventStoreRepositoryMockload.Lock()
	mock.calls.load = append(mock.calls.load, callInfo)
	lockdatabaseEventStoreRepositoryMockload.Unlock()
	return mock.loadFunc(ID)
}

// loadCalls gets all the calls that were made to load.
// Check the length with:
//     len(mockeddatabaseEventStoreRepository.loadCalls())
func (mock *databaseEventStoreRepositoryMock) loadCalls() []struct {
	ID string
} {
	var calls []struct {
		ID string
	}
	lockdatabaseEventStoreRepositoryMockload.RLock()
	calls = mock.calls.load
	lockdatabaseEventStoreRepositoryMockload.RUnlock()
	return calls
}

// save calls saveFunc.
func (mock *databaseEventStoreRepositoryMock) save(message *domainMessage) error {
	if mock.saveFunc == nil {
		panic("moq: databaseEventStoreRepositoryMock.saveFunc is nil but databaseEventStoreRepository.save was just called")
	}
	callInfo := struct {
		Message *domainMessage
	}{
		Message: message,
	}
	lockdatabaseEventStoreRepositoryMocksave.Lock()
	mock.calls.save = append(mock.calls.save, callInfo)
	lockdatabaseEventStoreRepositoryMocksave.Unlock()
	return mock.saveFunc(message)
}

// saveCalls gets all the calls that were made to save.
// Check the length with:
//     len(mockeddatabaseEventStoreRepository.saveCalls())
func (mock *databaseEventStoreRepositoryMock) saveCalls() []struct {
	Message *domainMessage
} {
	var calls []struct {
		Message *domainMessage
	}
	lockdatabaseEventStoreRepositoryMocksave.RLock()
	calls = mock.calls.save
	lockdatabaseEventStoreRepositoryMocksave.RUnlock()
	return calls
}

// setEvents calls setEventsFunc.
func (mock *databaseEventStoreRepositoryMock) setEvents(ev map[string]Event) {
	if mock.setEventsFunc == nil {
		panic("moq: databaseEventStoreRepositoryMock.setEventsFunc is nil but databaseEventStoreRepository.setEvents was just called")
	}
	callInfo := struct {
		Ev map[string]Event
	}{
		Ev: ev,
	}
	lockdatabaseEventStoreRepositoryMocksetEvents.Lock()
	mock.calls.setEvents = append(mock.calls.setEvents, callInfo)
	lockdatabaseEventStoreRepositoryMocksetEvents.Unlock()
	mock.setEventsFunc(ev)
}

// setEventsCalls gets all the calls that were made to setEvents.
// Check the length with:
//     len(mockeddatabaseEventStoreRepository.setEventsCalls())
func (mock *databaseEventStoreRepositoryMock) setEventsCalls() []struct {
	Ev map[string]Event
} {
	var calls []struct {
		Ev map[string]Event
	}
	lockdatabaseEventStoreRepositoryMocksetEvents.RLock()
	calls = mock.calls.setEvents
	lockdatabaseEventStoreRepositoryMocksetEvents.RUnlock()
	return calls
}