package eventSourcing

var databaseConnection string

func SetDatabaseConnection(dataUrl string) {
	databaseConnection = dataUrl
}
