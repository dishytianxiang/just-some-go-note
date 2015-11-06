package main

import (
	"fmt"
)

func main() {
	a := 0
	for {
		if a >= 3 { //if not have this expression, for loop will be infinite
			fmt.Println("--- ", a)
			break
		}
		a++
		fmt.Println("+++ ", a)
	}
LABLE1:
	for {
		fmt.Println("for lable1")
		for i := 0; i < 10; i++ {
			if i >= 9 {
				fmt.Println("i = ", i)
				break LABLE1 //jump the first for loop
			}
		}

	}
	for {
		for j := 0; j < 10; j++ {
			if j >= 9 {
				fmt.Println("j = ", j)
				goto LABLE2 //jump the first for loop
			}
		}
		fmt.Println("for lable2")
	}
LABLE2:
	fmt.Println("lable2")
	switch {
	case a >= 2:
		fmt.Println("switch a = ", a)
		fallthrough //judge the next case
	case a >= 1:
		fmt.Println("switch a >=1 ", a)
	case a >= 0:
		fmt.Println("switch a >=0 ", a)
	default:
		fmt.Println("switch default!")
	}

}
