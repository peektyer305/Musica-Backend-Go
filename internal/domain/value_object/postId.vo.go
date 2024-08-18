package valueobject

import (
	uuid "github.com/satori/go.uuid"
)

type PostId struct {
	value uuid.UUID
}

func NewPostId(value string) (PostId, error) {
	id, err := uuid.FromString(value)
	if err != nil {
		return PostId{}, err
	}
	return PostId{value: id}, nil
}

func (p PostId) String() string {
	return p.value.String()
}

func (p PostId) Equals(other PostId) bool {
	return p.value == other.value
}

func (p PostId) Value() uuid.UUID {
	return p.value
}