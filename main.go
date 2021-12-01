package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	mu sync.Mutex
<<<<<<< HEAD
	v map[string]int
}

func (c *SafeCounter) Inc(key string){
=======
	v  map[string]int
}

func (c *SafeCounter) Inc(key string) {
>>>>>>> 3261a91 (push)
	c.mu.Lock()
	c.v[key]++
	c.mu.Unlock()
}

func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v[key]
}

<<<<<<< HEAD

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++{
=======
func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
>>>>>>> 3261a91 (push)
		go c.Inc("somekey")
	}
	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
