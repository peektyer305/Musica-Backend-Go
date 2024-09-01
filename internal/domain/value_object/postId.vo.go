package valueobject

import (
	"database/sql/driver"

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

func (p PostId) GetUUID() uuid.UUID {
	return p.value
}
// Scan は、SQLからの読み込みのために必要なメソッドです。
// Value は、SQLへの書き込みのために必要なメソッドです。
//valueobject.PostId 型に Scanner と Valuer インターフェースを実装することで、データベースとGo構造体間のデータ変換をサポートする必要があります
func (p *PostId) Scan(value interface{}) error {
	id, err := uuid.FromString(value.(string))
	if err != nil {
		return err
	}
	p.value = id
	return nil
}

func (p PostId) Value() (driver.Value, error) {
	return p.value.String(), nil
}