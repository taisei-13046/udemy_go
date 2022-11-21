package main

import (
	"fmt"
	"time"
)

func getOsName() string {
	return "mac"
}

func main() {
	switch os := getOsName(); os {
	case "mac":
		fmt.Println("mac")
	case "windows":
		fmt.Println("windows")
	default:
		fmt.Println("default")
	}

	t := time.Now()
	fmt.Println(t.Hour())

	// switch文に条件式がなくてもいい
	switch {
	case t.Hour() < 12:
		fmt.Println("morning")
	}
}
