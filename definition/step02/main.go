package main

import (
	"fmt"
	// packageのimportで階層がしたの場合、/で区切る
	"os/user"
	"time"
)

func main() {
	fmt.Println(time.Now())
	fmt.Println(user.Current())
}
