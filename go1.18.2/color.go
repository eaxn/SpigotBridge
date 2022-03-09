package main

import "fmt"

type RGB [3]uint8
type RGBA [4]uint8

func (R RGB) GetR() uint8 {
	return R[0]
}
func (R RGB) GetG() uint8 {
	return R[1]
}
func (R RGB) GetB() uint8 {
	return R[2]
}
func (R RGB) ToInt() int32 {
	a := []bool{
		R.GetR()&uint8(BitAt(0)) == uint8(BitAt(0)),
		R.GetR()&uint8(BitAt(1)) == uint8(BitAt(1)),
		R.GetR()&uint8(BitAt(2)) == uint8(BitAt(2)),
		R.GetR()&uint8(BitAt(3)) == uint8(BitAt(3)),
		R.GetR()&uint8(BitAt(4)) == uint8(BitAt(4)),
		R.GetR()&uint8(BitAt(5)) == uint8(BitAt(5)),
		R.GetR()&uint8(BitAt(6)) == uint8(BitAt(6)),
		R.GetR()&uint8(BitAt(7)) == uint8(BitAt(7)),
		R.GetG()&uint8(BitAt(0)) == uint8(BitAt(0)),
		R.GetG()&uint8(BitAt(1)) == uint8(BitAt(1)),
		R.GetG()&uint8(BitAt(2)) == uint8(BitAt(2)),
		R.GetG()&uint8(BitAt(3)) == uint8(BitAt(3)),
		R.GetG()&uint8(BitAt(4)) == uint8(BitAt(4)),
		R.GetG()&uint8(BitAt(5)) == uint8(BitAt(5)),
		R.GetG()&uint8(BitAt(6)) == uint8(BitAt(6)),
		R.GetG()&uint8(BitAt(7)) == uint8(BitAt(7)),
		R.GetB()&uint8(BitAt(0)) == uint8(BitAt(0)),
		R.GetB()&uint8(BitAt(1)) == uint8(BitAt(1)),
		R.GetB()&uint8(BitAt(2)) == uint8(BitAt(2)),
		R.GetB()&uint8(BitAt(3)) == uint8(BitAt(3)),
		R.GetB()&uint8(BitAt(4)) == uint8(BitAt(4)),
		R.GetB()&uint8(BitAt(5)) == uint8(BitAt(5)),
		R.GetB()&uint8(BitAt(6)) == uint8(BitAt(6)),
		R.GetB()&uint8(BitAt(7)) == uint8(BitAt(7)),
	}
	r := int32(0)
	for x := 0; x < len(a); x++ {
		if a[x] {
			r += int32(BitAt(x))
		}
	}
	return r
}

func (R RGB) MinecraftFormat() string {
	RGBHex := fmt.Sprintf("%x", R.ToInt())
	return "§#" + RGBHex
}
