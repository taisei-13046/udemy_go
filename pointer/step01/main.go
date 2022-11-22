package main

import "fmt"

// pointerを引数にすると参照渡しができる
func one(x *int) {
	*x = 1
}

func main() {
	var n int = 100
	one(&n)
	fmt.Println(n)

	//fmt.Println(n)

	//fmt.Println(&n)

	//var p *int = &n
	//fmt.Println(p)
	//fmt.Println(*p)
}
