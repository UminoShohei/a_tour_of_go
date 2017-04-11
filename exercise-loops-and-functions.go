package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	preZ := float64(0)
	for {
		z = z - (z*z-x)/(2*z)
		if math.Abs(preZ-z) < 1e-10 {
			return z
		}
		preZ = z
	}
}

func main() {
	fmt.Println(Sqrt(4))
}
