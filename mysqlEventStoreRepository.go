package eventSourcing

import (
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type mysqlEventStoreRepository struct {
	databaseUrl string
	events      map[string]Event
}

func (m *mysqlEventStoreRepository) setEvents(ev map[string]Event) {
	m.events = ev
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

	args := []interface{}{message.id, message.payload.Payload(), message.recorderOn, message.payload.Name()}
	_, err = stmt.Exec(args...)
	return err
}

func (m *mysqlEventStoreRepository) load(ID string) (*domainMessages, error) {
	c, err := m.databaseConnection()
	if err != nil {
		return nil, err
	}
	defer c.Close()
	rows, err := c.Query("SELECT uuid, payload, recorded_on, type FROM events where uuid = ?", ID)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	result := domainMessages{}

	for rows.Next() {
		var (
			id         string
			payload    string
			recordedOn string
			typeEvent  string
		)

		rows.Scan(&id, &payload, &recordedOn, &typeEvent)

		event := m.getEvent(typeEvent)

		json.Unmarshal([]byte(payload), &event)

		recorded, _ := time.Parse(time.RFC3339, recordedOn)
		dm := domainMessage{id: id, recorderOn: recorded, payload: event}

		result = append(result, &dm)
	}

	return &result, err

}
func (m *mysqlEventStoreRepository) databaseConnection() (*sql.DB, error) {
	return sql.Open("mysql", m.databaseUrl)
}
func (m *mysqlEventStoreRepository) getEvent(eventName string) Event {
	return m.events[eventName]
}
