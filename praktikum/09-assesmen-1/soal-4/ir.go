package main

import "fmt"

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func isValid(n int) bool {
	sum := 0
	lastDigit := 10 

	for n > 0 {
		digit := n % 10
		if digit > lastDigit {
			return false 
		}
		sum += digit
		lastDigit = digit
		n /= 10
	}

	return isPrime(sum)
}

func main() {
	var L, R int
	fmt.Scan(&L, &R)

	count := 0
	for i := L; i <= R; i++ {
		if isValid(i) {
			count++
		}
	}

	fmt.Println(count)
}