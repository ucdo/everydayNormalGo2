package main

import (
	anytype "Minerva/generics/anyType"
	"Minerva/generics/funct"
	"fmt"
)

type myint int

type ft float64

func main() {
	s := funct.A[int](1, 2, 3)
	fmt.Println(s)

	// any type of struct
	_ = anytype.Vector[int]{Inner: []int{1, 2, 3, 4}}
	_ = anytype.Vector[string]{Inner: []string{"1", "xx"}}

	anytype.Test[myint](1)

	anytype.AnyMax(1, 2)
	anytype.AnyMax("1", "2")
	// anytype.AnyMax([]byte("1"), []byte("2")) 错误的。不满足泛型约束
	var x ft = 1e-55
	anytype.NewtonSqrt(x)  // panic 因为断言的是动态类型
	anytype.NewtonSqrt2(x) // 通过，因为用反射获取的基础类型
}
