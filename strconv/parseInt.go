package main

import (
	"fmt"
	"strconv"
)

func parseInt(s string) (int, error) {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return n, nil
}

func main() {
	s := "2025"
	n, err := parseInt(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
}
