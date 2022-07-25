package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.RWMutex
	a := 1

	go func() {
		wg.
		a = 2
		wg.Unlock()
	}()

	wg.Lock()
	a = 3
	wg.Unlock()
	fmt.Println("a ===", a)
}
