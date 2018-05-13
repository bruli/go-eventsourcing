package eventSourcing

type eventStore interface {
	save(message *domainMessage) error
	load(ID string) ([]*domainMessage, error)
}
