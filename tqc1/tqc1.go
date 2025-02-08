package main

import "fmt"

func main() {

	var quantity int
	const appleJuicePrice float64 = 23.34

	fmt.Print("請輸入果汁數量:")
	fmt.Scanln(&quantity)

	totalCost := float64(quantity) * appleJuicePrice
	fmt.Printf("總工要花%#v", totalCost)

}
