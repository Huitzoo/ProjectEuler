package main

import "fmt"

func main() {
	res := 0
	for i := 1; i < 10; i++ {
		if i%5 == 0 || i%3 == 0 {
			res += i
		}
	}
	println(res)
}
