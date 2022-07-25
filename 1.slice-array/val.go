package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	fmt.Printf("init ptr %p val %v\n", &s, s)

	changeSlice(s)
	fmt.Printf("ret  ptr %p val %v\n", &s, s)

}

func changeSlice(s []int) {
	fmt.Printf("------ changeSlice ptr %p val %v\n", &s, s)
	for k := range s {
		s[k] += 1
	}
	fmt.Printf("------ changeSlice ptr %p val %v\n", &s, s)
}
