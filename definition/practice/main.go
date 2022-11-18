package main

import "fmt"

func main() {
	f := 1.11
	fmt.Println(int(f))

	m := map[string]int{
		"Mike":  20,
		"Nancy": 24,
		"Messi": 30,
	}
	fmt.Printf("%T %v", m, m)
}
