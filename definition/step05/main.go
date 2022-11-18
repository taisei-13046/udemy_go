package main

import "fmt"

func main() {
	var a [2]int
	a[0] = 100
	a[1] = 200
	fmt.Println(a)

	// 配列はサイズを変更できない
	var b [2]int = [2]int{100, 200}
	//b = append(b, 300)
	fmt.Println(b)

	// スライスはリサイズできる
	n := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(n)
	fmt.Println(n[2])

	var board = [][]int{
		{1, 2, 3},
		{3, 4, 5},
		{6, 7, 8},
	}
	fmt.Println(board)

	m := make([]int, 3, 5)
	fmt.Println(len(m), cap(m), m)
	m = append(m, 1, 1)
	fmt.Println(len(m), cap(m), m)

	//c := make([]int, 5)
	c := make([]int, 0, 5)

	for i := 0; i < 5; i++ {
		c = append(c, i)
		fmt.Println(c)
	}
	fmt.Println(c)

	// x -> 0のスライスをメモリに確保
	// y -> nil
	x := make([]int, 0)
	var y []int
	fmt.Println(x, y)

	// pythonの辞書型
	h := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(h)

	fmt.Println(h["apple"])

	m2 := make(map[string]int)
	m2["pc"] = 5000
	fmt.Println(m2)

	b2 := []byte{72, 73}
	fmt.Println(b2)
	fmt.Println(string(b2))
}
