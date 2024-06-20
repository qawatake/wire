//go:build wireinject
// +build wireinject

package main

import (
	"context"

	"github.com/google/wire"
)

func initializeX(context.Context) X {
	panic(wire.Build(Set))
}
