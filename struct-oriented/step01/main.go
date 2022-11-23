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

type Vertex3D struct {
	// Embeddedで継承できる
	Vertex
	Z int
}

// func 型名 関数名 返り値
func (v Vertex3D) Area3D() int {
	return v.X * v.Y * v.Z
}

func (v *Vertex3D) Scale3D(i int) {
	v.X = v.X * i
	v.Y = v.Y * i
	v.Z = v.Z * i
}

// コンストラクタ
func New(x, y, z int) *Vertex3D {
	return &Vertex3D{Vertex{x, y}, z}
}

func main() {
	v := New(3, 4, 5)
	fmt.Println(v.Area3D())

	v.Scale3D(10)
	fmt.Println(v.Area3D())
}
