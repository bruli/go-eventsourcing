package eventSourcing

type infrastructure struct {
	mysqlEventStoreRepository *mysqlEventStoreRepository
}

type handler struct {
	mysqlEventStoreHandler *mysqlEventStoreHandler
}

type pkgContainer struct {
	infrastructure infrastructure
	handler        handler
}

var container pkgContainer

func init() {
	container.infrastructure.mysqlEventStoreRepository = &mysqlEventStoreRepository{}
	container.handler.mysqlEventStoreHandler = &mysqlEventStoreHandler{}
}
