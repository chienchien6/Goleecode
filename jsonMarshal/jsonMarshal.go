package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	//JSON 序列化
	p := Person{Name: "Alice mei", Age: 18}
	jsonData, _ := json.Marshal(p)
	fmt.Println(string(jsonData))

	//JSON 反序列化
	var p2 Person
	json.Unmarshal(jsonData, &p2)
	fmt.Println(p2)

}
