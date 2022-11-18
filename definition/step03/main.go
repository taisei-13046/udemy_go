package main

import "fmt"

// まとめて宣言できる
// 関数外でも宣言できる
var (
	i    int     = 1
	f64  float64 = 1.2
	s    string  = "test"
	t, f bool    = true, false
)

func main() {
	fmt.Println(i, f64, s, t, f)

	// 関数内でしかできない
	xi := 1
	xf64 := 1.2

	fmt.Println(xi, xf64)
	// Printfは改行されない
	fmt.Printf("%T\n", xf64)
	fmt.Printf("%T", xi)
}
