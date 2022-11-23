package main

import "fmt"

// interface{} どんなものでも引数に渡せる
func do(i interface{}) {
	// 一度typeアサーションをしてから代入している
	//ii := i.(int)
	//ii *= 2
	//fmt.Println(ii)

	//ss := i.(string)
	//fmt.Println(ss + "mike")

	// switchとtypeはセット
	switch v := i.(type) {
	case int:
		fmt.Println(v * 2)
	case string:
		fmt.Println(v + "!")
	default:
		fmt.Println("I don't know")
	}
}

func main() {
	do(10)
	do("mike")
	do(true)
}
