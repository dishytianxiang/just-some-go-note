package main

import (
	"fmt"
)

type human struct {
	Sex int
}

type teacher struct {
	human
	Name string
	Age  int
}
type student struct {
	human
	Name string
	Age  int
}

func main() {
	a := teacher{Name: "dishy1", Age: 19, human: human{Sex: 1}}
	b := student{Name: "dishy2", Age: 20, human: human{Sex: 2}}
	a.Sex = 2
	b.Sex = 10
	fmt.Println(a, b.Sex)
}
