package main

import "fmt"

func main() {
	a := [1]int{1}
	b := a
	b[0] = 0
	fmt.Println(a) // 1
	fmt.Println(b) // 0
}
