package main


import (
"fmt"
"math"
)

func Sqrt(x float64) float64 {
	z := 2.5
	var before, next float64
	for {
		before = z
		z = z- (z*z-x)/(2*z)
		next = math.Nextafter(z, before)
		if next == before {
			break
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(6))
	fmt.Println(math.Sqrt(6))
}

