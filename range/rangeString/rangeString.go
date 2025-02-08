package main

import (
	"fmt"
)

func main() {
	names := []string{"John", "Peter", "Tom", "Mary", "Jane"}
	for _, name := range names {
		fmt.Println(name, "the length of name is", len(name))
	}
}
