package main

import "fmt"

// 一番先に呼ばれる
// 初期設定で使われる
func init() {
	fmt.Println("init")
}

func bazz() {
	fmt.Println("buzz")
}

func main() {
	fmt.Println("Hello world")
	bazz()
}
