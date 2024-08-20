package cockroach

import (
	"sync"

	"gorm.io/gorm"
)

var gormPostgresInstance *gorm.DB
var gomePostgresOnce = sync.Once{}

