package main

import "fmt"

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)

}

func fractionFactorialSum(n int) float64 {
	if n == 0 {
		return 0
	}
	return 1.0/float64(factorial(n)) + fractionFactorialSum(n-1)
}

func main() {
	var n int
	fmt.Print("請輸入正整數 n: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("請輸入正整數")
		return
	}

	fmt.Printf("Fraction Factorial Sum F(%d) = %.6f\n", n, fractionFactorialSum(n))
}
