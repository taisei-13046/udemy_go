package main

import "fmt"

// 型の定義
type Human interface {
	Say() string
}

type Person struct {
	Name string
}

func (p *Person) Say() string {
	p.Name = "Mr." + p.Name
	return p.Name
}

func DriveCar(human Human) {
	if human.Say() == "Mr.Mike" {
		fmt.Println("Run")
	} else {
		fmt.Println("Get out")
	}
}

func main() {
	// pointer型を使う場合はアドレスを渡す
	var mike Human = &Person{"Mike"}

	DriveCar(mike)
}
