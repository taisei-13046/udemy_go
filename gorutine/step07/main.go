package main

import (
	"fmt"
	"time"
)

func main() {
	// chanを返す
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)

	// loopに名前をつけたら抜けれる
OuterLoop:
	for {
		select {
		// 何も変数に入れないこともできる
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			break OuterLoop
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
	fmt.Println("#########")
}
