package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounterNew struct {
	lookup map[string]int
	lock   sync.Mutex
}

func (s SafeCounterNew) Incre(key string) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.lookup[key]++
}

func (s SafeCounterNew) Read(key string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	fmt.Println("read:", s.lookup[key])
}

func main() {
	s := SafeCounterNew{lookup: make(map[string]int)}
	for i := 0; i < 10; i++ {
		go s.Incre("somekey")
		go s.Read("somekey")
	}
	time.Sleep(time.Second)

	//s.Read("somekey")

}

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	defer c.mux.Unlock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	//c.mux.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mux.Unlock()
	return c.v[key]
}

func main2() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
