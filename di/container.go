//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
)

var providerSet = wire.NewSet(
)