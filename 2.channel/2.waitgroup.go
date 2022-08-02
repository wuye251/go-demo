package main

import (
	"fmt"
	"sync"
)

func A(mutex *sync.WaitGroup) {
	defer mutex.Done()
	fmt.Println("A func run")
}

func B(mutex *sync.WaitGroup) {
	defer mutex.Done()
	fmt.Println("B func run")
}
 
func F() {
	fmt.Println("end--F func run")
}
 
func main() {
	var mutex sync.WaitGroup

	mutex.Add(2)

	go A(&mutex)
	go B(&mutex)

	mutex.Wait()
	F()
}
