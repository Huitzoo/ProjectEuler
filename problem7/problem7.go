package main

import "math"

//10001

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
	i := 2
	p := 4
	last := 0
	for {
		p++
		if i == 10001 {
			break
		}
		if primeQuestion(float64(p)) {
			last = p
			i++
		}
	}
	println(last)
}
