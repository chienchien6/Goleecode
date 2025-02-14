package main

import "fmt"

func continuedFraction(a, b, n int) float64 {

	if n == 0 {
		return 0
	}
	return float64(a) + 1.0/(float64(b)+continuedFraction(a, b, n-1))
}

func main() {
	var a, b, n int
	fmt.Print("請輸入 a, b, n（用空格分隔）: ")
	fmt.Scan(&a, &b, &n)

	if n <= 0 {
		fmt.Println("請輸入正整數 n")
		return
	}

	fmt.Printf("Continued Fraction C(%d) = %.6f\n", n, continuedFraction(a, b, n))
}
