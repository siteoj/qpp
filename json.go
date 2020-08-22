package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Person struct {
	Name  string `json:"name"`
	Hobby string `json:"hobby"`
}

func main() {
	//1.准备一段json
	b := []byte(`{"Name":"luhan"
				,"Hobby":"1"}`)
	//2.声明结构体
	var p Person
	//3.解析
	err := json.Unmarshal(b, &p)
	if err != nil {
		fmt.Println("json err:", err)
	}
	int, err := strconv.Atoi(p.Hobby)
	println(int)
	fmt.Println(p)
}