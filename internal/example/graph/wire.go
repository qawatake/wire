package main

import (
	"context"

	"github.com/google/wire"
)

var Set = wire.NewSet(
	NewA,
	wire.Bind(new(I1), new(A)),
	NewB,
	wire.InterfaceValue(new(I2), C{}),
	NewD,
	wire.Value(E{}),
	wire.Struct(new(F), "*"),
	NewH,
	wire.FieldsOf(new(H), "G"),
	NewX,
)

type A struct{}

func NewA(context.Context) A {
	return A{}
}

type B struct {
	I1
}

func NewB(i I1) B {
	return B{i}
}

type C struct{}

type D struct {
	I2
}

func NewD(i I2) D {
	return D{i}
}

type E struct{}

type F struct {
	B
	D
	E
}

type G struct {
	A
}

type H struct {
	E
	G
}

func NewH(a A, e E) H {
	g := G{a}
	return H{e, g}
}

type X struct{}

func NewX(B, D, F, G) X {
	return X{}
}

type I1 interface{}

type I2 interface{}
