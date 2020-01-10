package main

import "math"

import "fmt"

var adjustA = 5
var adjustB = 10

func pythagorean(a, b float64) (bool, float64) {
	c := math.Pow(a, 2.0) + math.Pow(b, 2.0)
	if int(math.Mod(c, math.Sqrt(c))) == 0 {
		if a+b+c == 1000 {
			return true, (a * b * c)
		}
	}
	return false, 0.0
}

func main() {
	//a + b +c = 1000
	//a²+b²=c²
	//a < b <c
	//a2 + b2 = c2
	// a < c
	//r = a * b * c
	a := 3.0
	b := 4.0
	for {
		is, product := pythagorean(a, b)
		if is {
			fmt.Println(int(product))
			break
		}
		a
	}
}
