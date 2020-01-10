package main

import "strconv"

func isPalindrome(num int) bool {
	r := strconv.Itoa(num)
	size := len(r) - 1
	for i, dig := range r {
		if i == size {
			break
		}
		if string(dig) != string(r[size]) {
			return false
		}
		size--
	}
	return true
}

func main() {

	for i := 999; i >= 900; i-- {
		for j := 999; j >= 900; j-- {
			res := i * j
			if isPalindrome(res) {
				println(res)
				break
			}
		}
	}
}
