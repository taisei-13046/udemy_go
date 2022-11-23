package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *Counter) Inc(key string) {
	c.mux.Lock()
	c.v[key]++
	c.mux.Unlock()
}

func (c *Counter) Value(key string) int {
	c.mux.Lock()
	c.mux.Unlock()
	return c.v[key]
}

func main() {
	//c := make(map[string]int)
	c := Counter{v: make(map[string]int)}
	go func() {
		for i := 0; i < 10; i++ {
			c.Inc("key")
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c.Inc("key")
		}
	}()

	time.Sleep(1 * time.Second)
	fmt.Println(c.Value("key"))
}
