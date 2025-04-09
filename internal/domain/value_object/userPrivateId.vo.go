package valueobject

import (
	"database/sql/driver"

	uuid "github.com/satori/go.uuid"
)

type UserPrivateId struct {
	value uuid.UUID
}

func NewUserPrivateId(value string) (UserPrivateId, error) {
	id, err := uuid.FromString(value)
	if err != nil {
		return UserPrivateId{}, err
	}
	return UserPrivateId{value: id}, nil
}

func (u UserPrivateId) String() string {
	return u.value.String()
}

func (u UserPrivateId) Equals(other UserId) bool {
	return u.value == other.value
}

func (u UserPrivateId) GetUUID() uuid.UUID {
	return u.value
}

// Scan は、SQLからの読み込みのために必要なメソッドです。
// Value は、SQLへの書き込みのために必要なメソッドです。
//valueobject.UserId 型に Scanner と Valuer インターフェースを実装することで、データベースとGo構造体間のデータ変換をサポートする必要があります
func (u *UserPrivateId) Scan(value interface{}) error {
	id, err := uuid.FromString(value.(string))
	if err != nil {
		return err
	}
	u.value = id
	return nil
}

func (u UserPrivateId) Value() (driver.Value, error) {
	return u.value.String(), nil
}



