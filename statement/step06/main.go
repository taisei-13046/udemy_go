package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	save()
	fmt.Println("OK")

	file, err := os.Open("./main.go")
	if err != nil {
		log.Fatalln("Exit", err)
	}
	defer file.Close()
	data := make([]byte, 100)
	// errは上書きで使いまわせる
	count, err := file.Read(data)
	if err != nil {
		log.Fatalln("error")
	}
	fmt.Println(count, string(data))

	log.Println("logging")
	log.Printf("%T", "test")

	// ここでプログラムが終了する
	log.Fatalln("error")

}

func thirdPartyConnectDB() {
	// あまり推奨されていない
	panic("unable")
}

func save() {
	// panicをキャッチしてrecoverしている
	defer func() {
		s := recover()
		fmt.Println(s)
	}()
	thirdPartyConnectDB()
}
