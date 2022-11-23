package main

import "fmt"

// non-struct ttype
type MyInt int

func (i MyInt) Double() int {
	fmt.Printf("%T\n", i)
	return int(i * 2)
}

func main() {
	myInt := MyInt(10)
	fmt.Println(myInt.Double())
}
