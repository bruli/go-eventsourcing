package eventSourcing

//go:generate moq -out eventStoreRepositoryMock.go . eventStoreRepository
type eventStoreRepository interface {
	save(message *domainMessage) error
	load(ID string) (*domainMessages, error)
}
