package main

import "github.com/google/wire"

var Set = wire.NewSet(
	NewA,
	wire.Bind(new(I1), new(A)),
	NewB,
	wire.InterfaceValue(new(I2), c),
	NewD,
	wire.Value(e),
	wire.Struct(new(F), "*"),
	NewH,
	wire.FieldsOf(new(H), "G"),
	NewX,
)

type A struct{}

func NewA() A {
	return A{}
}

type B struct {
	I1
}

func NewB(i I1) B {
	return B{i}
}

type C struct{}

var c C

type D struct {
	I2
}

func NewD(i I2) D {
	return D{i}
}

type E struct{}

var e E

type F struct {
	B
	D
	E
}

type G struct {
	I1
}

type H struct {
	C
	G
}

func NewH(i I1, e E) H {
	g := G{i}
	return H{c, g}
}

type X struct{}

func NewX(B, D, F, G) X {
	return X{}
}

type I1 interface{}

type I2 interface{}
