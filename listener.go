package eventSourcing

//Listener interface
type Listener interface {
	Handle(event Event) error
}
