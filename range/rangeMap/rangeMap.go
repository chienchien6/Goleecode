package main

import "fmt"

func main() {

	cities := map[string]int{"北京": 200000, "上海": 800000, "广州": 30000, "深圳": 50000}

	for city, population := range cities {
		fmt.Printf("%s的人口是%d\n", city, population)
	}
}
