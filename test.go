package main

import "fmt"

func main() {
	//test_map()
	//test_func()
	test_slice()
}

func test_map() {
	var s map[string]int // 这样声明的map是 nil 不能直接赋值
	s = map[string]int{}
	s["xxx"] = 1
	fmt.Println(s)
}

func test_func() {
	s := [3]int{1, 2, 3}
	func(array [3]int) {
		array[0] = 7
		fmt.Println(array)
	}(s)
	fmt.Println(s)
}

func test_slice() {
	s := []int{1, 2, 3}
	func(s []int) {
		s = append(s, s[:]...)
		s[0] = 7
		fmt.Println(s)
	}(s)

	fmt.Println(s)
}

func appendSlice(slice []int) []int {
	// 向切片添加一个元素，注意这会改变切片的长度
	slice = append(slice, 4)
	return slice
}

func modifySliceElement(slice []int) {
	// 修改切片中的一个元素
	if len(slice) > 0 {
		slice[0] = 10
	}
}
