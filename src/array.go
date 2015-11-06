package main

import (
	"fmt"
)

func main() {
	var c = [2]int{1, 2}
	var d = [2]int{3, 5}
	c = d //you can  do c = d,but you can do this when d [2]int,beacuse go thank the are different type,d[1] wrong
	fmt.Println(c)
	fmt.Println(d)
	e := [...]int{19: 22}
	f := [20]int{10: 2, 16: 13} //set value according to index.
	fmt.Println(e)
	fmt.Println(f)
	a := [...]int{5, 2, 8, 1, 9}
	num := len(a)
	fmt.Println(a)
	for i := 0; i < num-1; i++ {
		for j := i + 1; j < num; j++ {
			if a[i] < a[j] {
				temp := a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}
	}
	fmt.Println(a)

	k := [10]int{}
	s1 := k[6:]
	fmt.Println(s1)
	s2 := make([]int, 3, 10)
	fmt.Println("len = ", len(s2), " cap = ", cap(s2))
	var a1 = []int{1, 2, 3, 4, 5, 4, 3, 2, 4, 5, 67, 8, 9, 90, 23}
	b1 := a1[5:10]
	b2 := b1[1]
	fmt.Println("b1 = ", b1)
	fmt.Println("b1[7] = ", b2)
	s3 := make([]int, 3, 6)
	s3 = append(s1, 3, 4, 1, 5, 3, 2)
	fmt.Println(s3)
	s4 := []int{4, 5, 3, 7, 0}
	s5 := []int{55, 66}
	copy(s4, s5) //s5 copy to s4,only silce can do this
	fmt.Println(s4)
}
