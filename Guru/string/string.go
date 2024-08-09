package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//update()
	var x = "xx中国"
	includeUtf8("x中国")
	fmt.Println([]rune(x))
}

func update() {
	s := "12341234."
	b := []byte(s)
	b[0] = 'x'
	fmt.Println(string(b))
	fmt.Println(s)
}

func includeUtf8(a string) {
	boolx := utf8.ValidString(a)
	fmt.Println(boolx)
}
