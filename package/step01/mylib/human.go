package mylib

import "fmt"

// 大文字にしないと他のpackageから参照できない
type Person struct {
	Name string
	Age  int
}

func Say() {
	fmt.Println("Human")
}
