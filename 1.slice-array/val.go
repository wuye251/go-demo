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

	//新增 追加使起扩容  再进行修改 验证是否实参s会修改
	s = append(s, 11)
	for k := range s {
		s[k] += 1
	}
	fmt.Printf("------------ append ptr %p val %v\n", &s, s)
}
