package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 100
	ch <- 200

	// 一度取り出せば容量が空く
	x := <-ch
	fmt.Println(x)

	ch <- 100

	close(ch)
	for c := range ch {
		fmt.Println(c)
	}
}
