package eventSourcing

import (
	"time"
	"errors"
)

type domainMessages struct {
	messages []*domainMessage
}

func (dms *domainMessages) addMessage(m *domainMessage)  {
	n := append(dms.messages, m)
	dms.messages = n
}

func (dms *domainMessages) getEvents() []Event  {
	var ev []Event

	for _, m := range dms.messages {
		e := append(ev, m.payload)
		ev = e
	}

	return ev
}
type domainMessage struct {
	id         string
	payload    Event
	recorderOn time.Time
}

func (dm *domainMessage) getRecorderOn() string  {
	return dm.recorderOn.Format(time.RFC3339Nano)
}

func (dm *domainMessage)getPayload() (string, error) {
	p := string(dm.payload.Payload())

	if 3 > len(p) {
		return "", errors.New("Empty payload")
	}
	return p, nil
}
