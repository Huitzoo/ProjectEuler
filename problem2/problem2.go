package main

func main() {
	fibo := make([]int, 3)
	fibo[0] = 1
	fibo[1] = 2
	res := 0
	for {
		if fibo[1] > 4000000 {
			break
		}
		if fibo[1]%2 == 0 {
			res += fibo[1]
		}
		fibo[2] = fibo[0] + fibo[1]
		fibo[0] = fibo[1]
		fibo[1] = fibo[2]
	}
	println("The sum of even-number is",fibo[1])
}
