package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4}
	s := a[1:4]
	s[1] = 0
	x := s
	x[0] = 1
	//fmt.Println(a)
	//fmt.Println(s)
	//fmt.Println(x)
	//fmt.Println(len(s), cap(s))
	//for k := range a {
	//	fmt.Printf("%p \n", &a[k])
	//}
	//fmt.Printf("%p\n", &a)

	//kuorong()
	//copy2()
	del()
}

func kuorong() {
	s := []int{1, 2}
	fmt.Printf("扩容前：%d ,%d ,%p \n", len(s), cap(s), s)

	for i := 0; i < 100000; i++ {
		s = append(s, 1)
	}
	fmt.Printf("扩容后：%d ,%d ,%p \n", len(s), cap(s), s)

	x := [2]int{12, 1}
	xx := x[:]
	fmt.Printf("扩容前：%d ,%d ,%p \n", len(xx), cap(xx), xx)

	for i := 0; i < 1000000; i++ {
		xx = append(xx, 1)
	}

	fmt.Printf("扩容前后：%d ,%d ,%p \n", len(xx), cap(xx), xx)
}

func copy2() {
	x := []int{1, 2, 3}
	//var s []int
	s := make([]int, len(x), cap(x))
	copy(s, x)
	fmt.Printf("%p \n", x)
	fmt.Printf("%p \n", s)
	fmt.Println(s)
}

func del() {
	x := []int{1, 2, 3}
	x = append(x[:1], x[2:]...)
	fmt.Println(x)
}
