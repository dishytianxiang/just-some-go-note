package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go Go(&wg, i)
	}
	wg.Wait()
}
func Go(wg *sync.WaitGroup, index int) {
	a := 1
	for i := 0; i < 100000000; i++ {
		a += i

	}
	fmt.Println(index, a)

	wg.Done()

}
