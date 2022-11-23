package main

import "fmt"

func gorutine1(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	// 合計したらchannelにsumを入れる
	c <- sum
}

func gorutine2(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	sum++
	// 合計したらchannelにsumを入れる
	c <- sum
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	c := make(chan int) // 15, 15
	go gorutine1(s, c)
	go gorutine2(s, c)
	// ブロッキングしてcが入るまで待っている
	x := <-c
	fmt.Println(x)
	y := <-c
	fmt.Println(y)
}
