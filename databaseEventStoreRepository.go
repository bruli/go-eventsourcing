package eventSourcing

//go:generate moq -out databaseEventStoreRepositoryMock.go . databaseEventStoreRepository
type databaseEventStoreRepository interface {
	save(message *domainMessage) error
	load(ID string) (*domainMessages, error)
	setEvents(ev map[string]Event)
}
