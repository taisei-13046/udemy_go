package main

import "fmt"

func main() {
	l := []string{"python", "go", "java"}

	for i := 0; i < len(l); i++ {
		fmt.Println(i, l[i])
	}

	// indexとvalueを取得できる
	for i, v := range l {
		fmt.Println(i, v)
	}
	m := map[string]int{"apple": 100, "banana": 200}

	// mapも展開できる
	for k, v := range m {
		fmt.Println(k, v)
	}
}
