package eventSourcing

type mysqlEventStoreRepository struct {
	databaseUrl string
}

func (m *mysqlEventStoreRepository) save(message *domainMessage) error {
	panic("implement me")
}

func (m *mysqlEventStoreRepository) load(ID string) (*domainMessages, error) {
	panic("implement me")
}
