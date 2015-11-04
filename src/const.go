package main

import (
	"fmt"
)

//const 右边不能有未知的变量  除了go自带的 和 同样是const
const (
	e = len(a)
)
const (
	a = "A"
	b
	c = iota
	d
)

//iota 很像下标啊  每个const 组里面iota是一个集合 是单独的 封闭的  不信你闻闻
func main() {

	fmt.Println(a)

	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}
