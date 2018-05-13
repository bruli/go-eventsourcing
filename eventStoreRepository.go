package eventSourcing

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

var errEmptyPayload = errors.New("Empty payload")

type eventStoreRepository struct {
}

func (e *eventStoreRepository) save(message *domainMessage) error {
	con, err := defaultDatabaseConnection()
	if err != nil {
		return err
	}
	defer con.Close()
	stmt, err := con.Prepare("INSERT INTO events(uuid, payload, recorded_on, type) VALUES (?, ?, ?, ?)")
	defer stmt.Close()
	if err != nil {
		return err
	}
	p, err := getPayload(message)

	if err != nil {
		return err
	}
	recorderOn := getRecorderOn(message)
	args := []interface{}{message.id, p, recorderOn, message.payload.Name()}
	_, err = stmt.Exec(args...)
	return err
}

func (e *eventStoreRepository) load(Id string) ([]*domainMessage, error) {
	c, err := defaultDatabaseConnection()
	if err != nil {
		return nil, err
	}
	defer c.Close()
	rows, err := c.Query("SELECT uuid, payload, recorded_on, type FROM events where uuid = ?", Id)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	var result []*domainMessage

	for rows.Next() {
		var (
			id         string
			payload    string
			recordedOn string
			typeEvent  string
		)

		rows.Scan(&id, &payload, &recordedOn, &typeEvent)

		event := stdEventBus.getEvent(typeEvent)

		json.Unmarshal([]byte(payload), &event)

		recorded, _ := time.Parse(time.RFC3339, recordedOn)
		dm := domainMessage{id, event, recorded}

		result = append(result, &dm)
	}

	return result, err

}

func getPayload(message *domainMessage) (string, error) {
	p := string(message.payload.Payload())

	if 3 > len(p) {
		return p, errEmptyPayload
	}
	return p, nil
}
func getRecorderOn(message *domainMessage) string {
	y, M, d := message.recorderOn.Date()
	h, m, s := message.recorderOn.Clock()
	u := message.recorderOn.Nanosecond()
	return fmt.Sprintf("%d-%d-%d %d:%d:%d.%d", y, M, d, h, m, s, u)
}

func defaultDatabaseConnection() (*sql.DB, error) {
	return sql.Open("mysql", databaseConnection)
}
