package main

type Array struct {
	H1 *Type
	H2 *Type
	H3 *Type
	H4 *Type
}

type Type struct {
	Name  string
	Value interface{}
}

func (A Array) Length() {
	// if A.H1 != nil // Check for another array in it.
}
