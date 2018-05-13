package eventSourcing

import (
	"time"
)

type domainMessage struct {
	id         string
	payload    Event
	recorderOn time.Time
}
