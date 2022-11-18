package main

import "fmt"

// constは型を指定していないためoverflowするような値を宣言可能
const Pi = 3.14
const (
	Username = "test_user"
	Password = "test_pass"
)

func main() {
	fmt.Println(Pi, Username, Password)
}
