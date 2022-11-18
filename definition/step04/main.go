package main

import (
	"fmt"
	"strconv"
	"strings"
)

// constは型を指定していないためoverflowするような値を宣言可能
const Pi = 3.14
const (
	Username = "test_user"
	Password = "test_pass"
)

func main() {
	fmt.Println(Pi, Username, Password)

	// num
	var (
		u8 uint8 = 255
		i8 int8  = 127
	)
	fmt.Println(u8, i8)

	fmt.Println("1 + 1 = ", 1+1)

	// string
	// ascコードで出力される 104
	fmt.Println(("hello" + "world")[0])
	fmt.Println(string(("hello" + "world")[0]))

	s := "hello world"

	// 破壊的な変更はしていない
	fmt.Println(strings.Replace(s, "h", "X", 1))

	// 改行通りに実行したい場合は``でかこう
	fmt.Println(`
		Print
		Print
				Print
	`)

	var x int = 1

	xx := float64(x)
	fmt.Printf("%T %v %f", xx, xx, xx)

	// string -> intはできない
	var str string = "14"
	//z = int(s)

	// _を使うと代入を避けられる
	i, _ := strconv.Atoi(str)
	fmt.Println(i)

	h := "Hello World"
	fmt.Println(string(h[0]))
}
