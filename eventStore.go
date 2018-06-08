package eventSourcing

//go:generate moq -out eventStoreMock.go . eventStore
type eventStore interface {
	save(message *domainMessage) error
	load(ID string) (*domainMessages, error)
}
