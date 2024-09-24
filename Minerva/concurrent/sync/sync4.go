package main

// 使用闭包传参
import (
	"fmt"
)

func f(i int) {
	fmt.Println(i)
}

func closer(x int) func() {
	return func() {
		f(x)
	}
}

func main() {
	c := closer(1)
	one.Do(c)
}
