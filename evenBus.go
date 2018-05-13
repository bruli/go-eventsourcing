package eventSourcing

type eventBus struct {
	listeners map[string][]Listener
	events    map[string][]Event
}

var stdEventBus eventBus

func init() {
	stdEventBus = eventBus{}
	stdEventBus.listeners = make(map[string][]Listener)
	stdEventBus.events = make(map[string][]Event)
}

func (e *eventBus) addListener(listener Listener, event Event) {
	n := append(e.listeners[event.Name()], listener)
	e.listeners[event.Name()] = n
}

func (e *eventBus) addEvent(ev Event) {
	n := append(e.events[ev.Name()], ev)
	e.events[ev.Name()] = n
}

func (e *eventBus) getListeners(event Event) []Listener {
	return e.listeners[event.Name()]
}

func (e *eventBus) getEvent(eventName string) Event {
	n := e.events[eventName]

	return n[0]
}
