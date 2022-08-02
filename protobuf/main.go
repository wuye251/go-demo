package main

import (
	"fmt"
	v1 "go-demo/protobuf/api/v1"
)

func main() {
	person := &v1.Person{
		Name:  "张三",
		Id:    7,
		Email: "boomboom@boom.com",
	}

	fmt.Println("person", person)
}
