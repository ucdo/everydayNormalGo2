package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//update()
	includeUtf8("中国")
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
