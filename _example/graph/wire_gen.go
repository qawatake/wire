// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"context"
)

// Injectors from injector.go:

func initializeX(contextContext context.Context) X {
	a := NewA(contextContext)
	b := NewB(a)
	i2 := _wireCValue
	d := NewD(i2)
	mainE := _wireEValue
	f := F{
		B: b,
		D: d,
		E: mainE,
	}
	h := NewH(a, mainE)
	g := h.G
	x := NewX(b, d, f, g)
	return x
}

var (
	_wireCValue = c
	_wireEValue = e
)