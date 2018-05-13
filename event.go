package eventSourcing

type Event interface {
	Name() string
	Payload() []byte
}
