package eventSourcing

type infrastructure struct {
	mysqlEventStoreRepository *mysqlEventStoreRepository
}

type handler struct {
	mysqlEventStoreHandler    *databaseEventStoreHandler
	listenersHandler          *listenersHandler
	inMemoryEventStoreHandler *inMemoryEventStoreHandler
}

type pkgContainer struct {
	infrastructure infrastructure
	handler        handler
}

var container pkgContainer

func init() {
	container.infrastructure.mysqlEventStoreRepository = &mysqlEventStoreRepository{}
	container.handler.listenersHandler = &listenersHandler{}
	container.handler.mysqlEventStoreHandler = &databaseEventStoreHandler{eventStore: container.infrastructure.mysqlEventStoreRepository,
		listenersHandler: container.handler.listenersHandler}
	container.handler.inMemoryEventStoreHandler = &inMemoryEventStoreHandler{listenersHandler: container.handler.listenersHandler}
}
