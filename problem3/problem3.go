package main

import "math"

var prime = float64(2)

func primeQuestion(i float64) bool {
	factor := int(math.Sqrt(i))
	for p := 2; p < factor; p++ {
		if int(i)%p == 0 {
			return false
		}
	}
	return true
}

func main() {

	num := float64(600851475143)
	factor1 := float64(2)

	for {
		if primeQuestion(factor1) {
			if int(math.Mod(num, factor1)) == 0 {
				num = num / factor1
				if num == 1 {
					break
				}
			}
		}
		factor1++
	}
	println("The biggest prime factor is: ", int(factor1))
}
