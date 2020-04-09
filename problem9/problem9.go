package main

import (
	"fmt"
	"math"
)

var number float64

func pythagorean(sqrtNumber, res float64) (bool, int) {

	for a := 1; a < int(number)/2; a++ {
		aFloat := float64(a)

		b := int((number * (number - 2*aFloat)) / (2 * (number - aFloat)))

		if a > b {
			continue
		}
		
		c := int(math.Sqrt(float64(a*a + b*b)))
		
		if a < b && b < c {
			if a*a+b*b == c*c {
				fmt.Println(a,b,c)
				return true,(a*b*c)
			}
		}
	}
	return false, 0
}

func main() {

	number = 1000
	sqrtNumber := math.Sqrt(float64(number / 2))
	res := sqrtNumber - float64(int(sqrtNumber))
	fmt.Println(pythagorean(sqrtNumber, res))
	
}
