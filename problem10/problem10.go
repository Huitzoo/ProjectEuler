package main

import (
	"math"
)

func primeQuestion(i float64) bool {
	factor := int(math.Sqrt(i))
	if int(i)%factor == 0 && factor != 1 {
		return false
	}
	for p := 2; p < factor; p++ {
		if int(i)%p == 0 {
			return false
		}
	}
	return true
}

func main() {
	num := float64(2000000)
	factor1 := float64(2)
	primeResult := 0.0
	i := 0.0
	for i < num {
		if primeQuestion(factor1) {
			primeResult += factor1
		}
		i++
		factor1++
	}
	println(int(primeResult))
}
