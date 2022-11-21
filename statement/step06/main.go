package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("aaaaa")
	if err != nil {
		log.Fatalln("Exit", err)
	}

	log.Println("logging")
	log.Printf("%T", "test")

	// ここでプログラムが終了する
	log.Fatalln("error")
}
