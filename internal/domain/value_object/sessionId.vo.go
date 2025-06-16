package valueobject

import (
	uuid "github.com/satori/go.uuid"
)

type SessionId struct {
	value uuid.UUID
}

func NewSessionId() SessionId {
	return SessionId{
		value: uuid.NewV4(),
	}
}

func (id SessionId) Value() uuid.UUID {
	return id.value
}
