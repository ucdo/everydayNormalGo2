package main

import "fmt"

func main() {
	var t = 100
	var p = &t
	*p = 9
	fmt.Println(t)
}
