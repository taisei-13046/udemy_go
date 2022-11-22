package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

// func 型名 関数名 返り値
func (v Vertex) Area() int {
	return v.X * v.Y
}

func (v *Vertex) Scale(i int) {
	v.X = v.X * i
	v.Y = v.Y * i
}

// コンストラクタ
func New(x, y int) *Vertex {
	return &Vertex{x, y}
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Area())

	v.Scale(10)
	fmt.Println(v)
}
