package main

import (
	"fmt"
)

func main() {
	// メモリ領域を確保したい
	var p *int = new(int)
	fmt.Println(p)
	*p++
	fmt.Println(*p)

	// nilになる
	var p2 *int
	fmt.Println(p2)

	s := make([]int, 0)
	fmt.Printf("%T\n", s)

	m := make(map[string]int)
	fmt.Printf("%T\n", m)

	// newとmakeの違いは、newはpointerを返す
	var p3 *int = new(int)
	fmt.Printf("%T\n", p3)
}
