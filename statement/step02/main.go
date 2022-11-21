package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 3 {
			// スキップする
			continue
		}

		if i > 5 {
			// forを終了する
			break
		}
		fmt.Println(i)
	}

	// whileっぽくも描ける
	sum := 1
	for sum < 10 {
		sum++
		fmt.Println(sum)
	}

	// 無限ループも簡単に作れる
	//for {
	//	fmt.Println("hello")
	//}
}
