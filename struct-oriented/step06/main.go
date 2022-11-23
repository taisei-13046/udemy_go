package main

import "fmt"

type UserNotFound struct {
	UserName string
}

// errorの際はpointerで渡してあげる
func (e *UserNotFound) Error() string {
	return fmt.Sprintf("User not found %v", e.UserName)
}

func myFunc() error {
	ok := false
	if ok {
		return nil
	}
	return &UserNotFound{UserName: "Mike"}
}

func main() {
	if err := myFunc(); err != nil {
		fmt.Println(err)
	}
}
