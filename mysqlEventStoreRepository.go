package eventSourcing

import "database/sql"

type mysqlEventStoreRepository struct {
	databaseUrl string
}

func (m *mysqlEventStoreRepository) save(message *domainMessage) error {
	con, err := m.databaseConnection()
	if err != nil {
		return err
	}
	defer con.Close()
	stmt, err := con.Prepare("INSERT INTO events(uuid, payload, recorded_on, type) VALUES (?, ?, ?, ?)")
	defer stmt.Close()
	if err != nil {
		return err
	}

	args := []interface{}{message.id, message.payload, message.recorderOn, message.payload.Name()}
	_, err = stmt.Exec(args...)
	return err
}

func (m *mysqlEventStoreRepository) load(ID string) (*domainMessages, error) {
	panic("implement me")
}
func (m *mysqlEventStoreRepository) databaseConnection() (*sql.DB, error) {
	return sql.Open("mysql", m.databaseUrl)
}
