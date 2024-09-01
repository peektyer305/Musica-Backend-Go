package valueobject

import (
	"database/sql/driver"

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

func (u UserId) GetUUID() uuid.UUID {
	return u.value
}

// Scan は、SQLからの読み込みのために必要なメソッドです。
// Value は、SQLへの書き込みのために必要なメソッドです。
//valueobject.UserId 型に Scanner と Valuer インターフェースを実装することで、データベースとGo構造体間のデータ変換をサポートする必要があります
func (u *UserId) Scan(value interface{}) error {
	id, err := uuid.FromString(value.(string))
	if err != nil {
		return err
	}
	u.value = id
	return nil
}

func (u UserId) Value() (driver.Value, error) {
	return u.value.String(), nil
}



