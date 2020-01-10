package main

import (
	"fmt"
	"math"
)

func main() {
	a := 100.0
	r := a * (a + 1) * (2*a + 1) / 6
	r1 := (a * (a + 1)) / 2

	fmt.Println(int(math.Pow(r1, 2) - r))
	//25164150
	//25164150
}
