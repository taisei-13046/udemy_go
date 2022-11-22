package main

import "fmt"

type Vertex struct {
	// typeは全て大文字
	X int
	Y int
}

func changeVertex(v Vertex) {
	v.X = 1000
}

func changeVertex2(v *Vertex) {
	v.X = 1000
}

func main() {
	value := Vertex{1, 2}
	changeVertex(value)
	fmt.Println(value)

	changeVertex2(&value)
	fmt.Println(value)

	v := Vertex{X: 1, Y: 2}
	fmt.Println(v)
	fmt.Println(v.X, v.Y)

	v.X = 100
	fmt.Println(v)

	// 何もないと初期値が入る
	v2 := Vertex{X: 1}
	fmt.Println(v2)

	// *main.Vertex
	v6 := new(Vertex)
	fmt.Printf("%T\n", v6)

	// こっちをよく使う
	v7 := &Vertex{}
	fmt.Printf("%T", v7)
}
