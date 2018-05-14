package eventSourcing

//go:generate moq -out eventStoreMock.go . eventStoreI
type eventStoreI interface {
	save(message *domainMessage) error
	load(ID string) (*domainMessages, error)
}
