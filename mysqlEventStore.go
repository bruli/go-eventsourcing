package eventSourcing

type MysqlEventStore struct {
	DatabaseUrl string
}

func (mes *MysqlEventStore) Init() {
	container.infrastructure.mysqlEventStoreRepository.databaseUrl = mes.DatabaseUrl
	container.handler.mysqlEventStoreHandler.init()
}

func (mes *MysqlEventStore) DeclareListener(list Listener, ev Event) {
	container.handler.mysqlEventStoreHandler.declareListener(list, ev)
}

func (mes *MysqlEventStore) DeclareEvent(ev Event) {
	container.handler.mysqlEventStoreHandler.declareEvent(ev)
}

func (mes *MysqlEventStore) ApplyNewEvent(e Event) {
	container.handler.mysqlEventStoreHandler.applyNewEvent(e)
}

func (mes *MysqlEventStore) Save(a Aggregate) error {
	return container.handler.mysqlEventStoreHandler.save(a)
}

func (mes *MysqlEventStore) Load(id string, agg Aggregate) error {
	return container.handler.mysqlEventStoreHandler.load(id, agg)
}
