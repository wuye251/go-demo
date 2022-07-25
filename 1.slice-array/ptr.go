package main

import "fmt"

func main() {
	slice := make([]int, 2, 3)
	for i := 0; i < len(slice); i++ {
		slice[i] = i
	}
	fmt.Printf("init slice val %v  ptr %p  len %v  cap %v \n", slice, &slice, len(slice), cap(slice)) // [0,1]  容量是3
	ret := changeSlice(&slice)
	//slice = [10, 1, 3](但是因为len() cap()修改的拷贝的副本的值 slice实参为修改 len=2 cap=3)
	//  ret = [10, 100, 3, 4]
	ret[1] = 111
	//ret = [10, 111, 3, 4]

	//结果：slice = [10, 1]  ret = [10, 111, 3, 4]
	fmt.Printf("ret  slice val %v  ptr %p  len %v  cap %v \n", slice, &slice, len(slice), cap(slice))
	fmt.Printf("ret  ret   val %v ptr %p  len %v  cap %v \n", ret, &ret, len(ret), cap(ret))

}

func changeSlice(s *[]int) []int {
	(*s)[0] = 10 //这里还没有触发扩容  所以修改实参   slice = [10, 1]
	fmt.Printf("----     s val %v     ptr %p  len %v  cap %v \n", *s, s, len(*s), cap(*s))
	*s = append(*s, 3) //未触发扩容 slice = [10, 1, 3]
	fmt.Printf("----     s val %v   ptr %p  len %v  cap %v \n", *s, s, len(*s), cap(*s))
	*s = append(*s, 4) //触发扩容  所以移动到了新地址  slice = [10, 1, 3]  new = [10, 1, 3, 4]
	fmt.Printf("----     s val %v ptr %p  len %v  cap %v \n", *s, s, len(*s), cap(*s))
	(*s)[1] = 100 //修改新地址数据 new = [10, 100, 3, 4] slice = [10, 1, 3]

	return *s
}
