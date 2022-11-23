package main

import (
	"fmt"
	"time"
)

func gorutine1(ch chan string) {
	for {
		ch <- "packet from 1"
		time.Sleep(time.Second * 1)
	}
}

func gorutine2(ch chan string) {
	for {
		ch <- "packet from 2"
		time.Sleep(time.Second * 1)
	}
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go gorutine1(c1)
	go gorutine2(c2)

	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}
