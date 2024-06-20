//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

func initializeX() X {
	panic(wire.Build(Set))
}
