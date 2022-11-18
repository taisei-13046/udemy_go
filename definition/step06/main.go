package main

import "fmt"

func add(x, y int) (int, int) {
	fmt.Println("add")

	return x + y, x - y
}

// 返り値で宣言
// 返り値が多い場合に推測できるように
func cal(price, item int) (result int) {
	result = price * item
	return
}

func main() {
	r1, r2 := add(10, 20)
	fmt.Println(r1, r2)

	r3 := cal(100, 2)
	fmt.Println(r3)

	f := func() {
		fmt.Println("inner")
	}
	f()

	// 即時に実行することもできる
	func() {
		fmt.Println("inner")
	}()

	// クロージャー
	x := 0
	inc := func() int {
		x++
		return x
	}
	fmt.Println(inc())

	counter := incGen()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())

	c1 := circleArea(3.14)
	fmt.Println(c1(2))

	foo(100, 200)
	foo(100, 200, 300)

	s := []int{1, 2, 4}
	// 可変長に入れる場合
	foo(s...)
}

func foo(params ...int) {
	fmt.Println(len(params), params)
}

func circleArea(pi float64) func(radius float64) float64 {
	return func(radius float64) float64 {
		return pi * radius * radius
	}
}

// クロージャ関数
func incGen() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
