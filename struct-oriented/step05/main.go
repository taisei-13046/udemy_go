package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// Stringメソッドを定義すると出力方法を変えられる
func (p Person) String() string {
	return fmt.Sprintf("My name is %v", p.Name)
}

func main() {
	mike := Person{"Mike", 20}
	fmt.Println(mike)
}
