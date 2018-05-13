package eventSourcing

var currentEvents events

func NewCurrentEvents() {
	currentEvents = events{}
}

type events struct {
	events []Event
}

func GetCurrentEvents() []Event {
	return currentEvents.events
}
