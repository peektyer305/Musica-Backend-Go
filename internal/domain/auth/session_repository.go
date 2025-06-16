package auth

import (
	valueobject "Musica-Backend/internal/domain/value_object"
)

type SessionRepository interface {
	Save(session *Session) error
	Find(id valueobject.SessionId) (*Session, error)
	Delete(id valueobject.SessionId) error
}
