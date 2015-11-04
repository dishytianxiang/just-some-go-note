package main

import (
	"fmt"
	"math"
)

type (
	rune int32 //Set alias
)

func main() {
	var b rune
	b = 34
	var _, h, c, _ int = 1, 2, 3, 4

	fmt.Println(c)

	fmt.Println(h, b)
	fmt.Println(math.MaxInt8)
	fmt.Println(math.MinInt8)
	var k int = 65
	j := string(k)
	fmt.Println(j)
}
