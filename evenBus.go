package eventSourcing

type pkgEventBus struct {
	listeners map[string][]Listener
	events    map[string][]Event
}

var eventBus pkgEventBus

func init() {
	eventBus = pkgEventBus{}
	eventBus.listeners = make(map[string][]Listener)
	eventBus.events = make(map[string][]Event)
}

func (e *pkgEventBus) addListener(listener Listener, event Event) {
	n := append(e.listeners[event.Name()], listener)
	e.listeners[event.Name()] = n
}

func (e *pkgEventBus) addEvent(ev Event) {
	n := append(e.events[ev.Name()], ev)
	e.events[ev.Name()] = n
}

func (e *pkgEventBus) getListeners(event Event) []Listener {
	return e.listeners[event.Name()]
}

func (e *pkgEventBus) getEvent(eventName string) Event {
	n := e.events[eventName]

	return n[0]
}
