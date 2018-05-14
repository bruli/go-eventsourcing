package eventSourcing

import (
	"database/sql"
	"encoding/json"
	"time"
)

type MysqlEventStore struct {
	DatabaseUrl string
}

func (e *MysqlEventStore) save(message *domainMessage) error {
	con, err := e.databaseConnection()
	if err != nil {
		return err
	}
	defer con.Close()
	stmt, err := con.Prepare("INSERT INTO events(uuid, payload, recorded_on, type) VALUES (?, ?, ?, ?)")
	defer stmt.Close()
	if err != nil {
		return err
	}
	p, err := message.getPayload()

	if err != nil {
		return err
	}
	recorderOn := message.getRecorderOn()
	args := []interface{}{message.id, p, recorderOn, message.payload.Name()}
	_, err = stmt.Exec(args...)
	return err
}

func (e *MysqlEventStore) load(Id string) (*domainMessages, error) {
	c, err := e.databaseConnection()
	if err != nil {
		return nil, err
	}
	defer c.Close()
	rows, err := c.Query("SELECT uuid, payload, recorded_on, type FROM events where uuid = ?", Id)
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

		event := eventBus.getEvent(typeEvent)

		json.Unmarshal([]byte(payload), &event)

		recorded, _ := time.Parse(time.RFC3339, recordedOn)
		dm := domainMessage{id: id, recorderOn: recorded, payload: event}

		result.addMessage(&dm)
	}

	return &result, err

}
func (e *MysqlEventStore) databaseConnection() (*sql.DB, error) {
	return sql.Open("mysql", e.DatabaseUrl)
}
