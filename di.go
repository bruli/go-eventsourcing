package eventSourcing

type infrastructure struct {
	mysqlEventStoreRepository *mysqlEventStoreRepository
}

type handler struct {
	mysqlEventStoreHandler *mysqlEventStoreHandler
	listenersHandler       *listenersHandler
}

type pkgContainer struct {
	infrastructure infrastructure
	handler        handler
}

var container pkgContainer

func init() {
	container.infrastructure.mysqlEventStoreRepository = &mysqlEventStoreRepository{}
	container.handler.listenersHandler = &listenersHandler{}
	container.handler.mysqlEventStoreHandler = &mysqlEventStoreHandler{eventStore: container.infrastructure.mysqlEventStoreRepository,
		listenersHandler: container.handler.listenersHandler}
}
