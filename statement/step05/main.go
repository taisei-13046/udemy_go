package main

import (
	"fmt"
	"os"
)

func foo() {
	defer fmt.Println("world foo")

	fmt.Println("hello foo")
}

func main() {
	foo()
	// main関数実行後に処理
	defer fmt.Println("world")

	fmt.Println("hello")

	// 初めに入れたものが最後に実行される
	fmt.Println("run")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("success")

	file, _ := os.Open("./main.go")
	// 最後に実行を保証してくれる
	defer file.Close()
	data := make([]byte, 100)
	file.Read(data)
	fmt.Println(string(data))
}
