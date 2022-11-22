package main

import "fmt"

func main2() {
	var i int = 100
	var j int = 200
	var p1 *int
	var p2 *int
	p1 = &i
	p2 = &j

	i = *p1 + *p2
	p2 = p1
	j = *p2 + i
	fmt.Println(j)
}

func main() {
	var i int = 10
	var p *int
	p = &i
	var j int
	j = *p
	fmt.Println(j)

	main2()
}
