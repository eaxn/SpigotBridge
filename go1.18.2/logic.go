package main

import "math"

func BitAt(n int) int {
	return int(math.Pow(2, float64(n)))
}
