package main

import "fmt"

func DeferFunc1(i int) (t int) {
	t = i
	defer func() {
		t += 3
	}()
	return t
} //4

func DeferFunc2(i int) int {
	t := i
	defer func() {
		t += 3
	}()
	return t
} //4

func DeferFunc3(i int) (t int) {
	defer func() {
		t += i
	}()
	return 2
} //3

func DeferFunc4() (t int) {
	defer func(i int) {
		fmt.Println(i)
		fmt.Println(t)
	}(t)
	t = 1
	return 2
} //0,2

func main() {
	fmt.Println(DeferFunc1(1))
	fmt.Println(DeferFunc2(1))
	fmt.Println(DeferFunc3(1))
	DeferFunc4()
}
