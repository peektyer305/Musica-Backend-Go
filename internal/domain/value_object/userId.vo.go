package valueobject

import (
	uuid "github.com/satori/go.uuid"
)

type UserId struct {
	value uuid.UUID
}

func NewUserId(value string) (UserId, error) {
	id, err := uuid.FromString(value)
	if err != nil {
		return UserId{}, err
	}
	return UserId{value: id}, nil
}

func (u UserId) String() string {
	return u.value.String()
}

func (u UserId) Equals(other UserId) bool {
	return u.value == other.value
}

func (u UserId) Value() uuid.UUID {
	return u.value
}

