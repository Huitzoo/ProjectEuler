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

func makeRangeInt(from, to int) ([]int, []int) {
	primes := make([]int, 0)
	nonPrimes := make([]int, 0)
	for i := from; i < to; i++ {
		if primeQuestion(float64(i)) {
			primes = append(primes, i)
		} else {
			nonPrimes = append(nonPrimes, i)
		}
		//numbers[i] = from + i
	}
	return primes, nonPrimes
}
func main() {

	primes, _ := makeRangeInt(1, 20)
	r := 1
	for i := 0; i < len(primes); i++ {
		a := int(math.Floor(math.Log(20) / math.Log(float64(primes[i]))))
		r = r * int(math.Pow(float64(primes[i]), float64(a)))
	}
	println(r)

}
