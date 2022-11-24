package main

import (
	"fmt"

	// 一番親のディレクトリにgo.modを作成してそこを起点にimportする
	"github.com/taisei-13046/udemy_go/package/step01/mylib"
)

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(mylib.Average(s))
	mylib.Say()
	person := mylib.Person{Name: "Mike", Age: 20}
	fmt.Println(person)
}
