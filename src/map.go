package main

import (
	"fmt"
)

func main() {
	m := make(map[int]string)
	m[1] = "OK"
	delete(m, 1)
	a := m[1]
	fmt.Println(a)

	var n map[int]map[int]string
	n = make(map[int]map[int]string)
	n[1] = make(map[int]string)
	n[1][1] = "OK"
	fmt.Println(n)
	b, ok := n[4][3] //this can check n[4] if it exists
	if !ok {
		n[4] = make(map[int]string)
		n[4][3] = "OKOK"
	}
	b, ok = n[4][3]
	fmt.Println(b)
	fmt.Println(n)
	sm := make([]map[int]string, 5)
	for i, _ := range sm {
		sm[i] = make(map[int]string, 1)
		sm[i][1] = "OK"
	}
	fmt.Println(sm)
	sd := map[int]string{1: "ha1", 2: "ha2", 3: "ha3"}
	sdt := make([]int, len(sd))
	fmt.Println(sd)
	fmt.Println(sdt)
}
