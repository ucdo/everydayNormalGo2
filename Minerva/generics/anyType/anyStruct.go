package anytype

import (
	"fmt"
	"reflect"
)

// Vector 可以存任意类型了
// 如果不是泛型的话，要定义各自类型的slice
// 比如 []int{} []string{}
// 这里的泛型就解决了这个问题
type Vector[T any] struct {
	Inner []T
}

func (v *Vector[T]) Len() int {
	return len(v.Inner)
}

func (v *Vector[T]) Get(key int) T {
	return v.Inner[key]
}

type Array[T any] [8]T // 泛型数组

type Slice[T any] []T // 泛型切片

type Map[K comparable, V any] map[K]V // 泛型哈希表

type Chan[T any] chan T // 泛型通道

type Iterator[T any] interface { // 泛型接口
	Next() bool
	Value() T
}

type ListIter[T any] struct {
	index int
	inner []T
}

func (l *ListIter[T]) Next() bool {
	if l.index < len(l.inner)-1 {
		l.index++
		return true
	}

	return false
}

func (l *ListIter[T]) Value() T {
	if l.index >= 0 && l.index < len(l.inner) {
		return l.inner[l.index]
	}

	var empty T
	return empty
}

// Fn 其他的
type Fn[A, B any] func(A) B

type MyInt[T fmt.Stringer] int

// 约束
// 近似约束
type it interface {
	~int
}

// 实际约束
type it2 interface {
	int
}

// Test 这里的T可以是int，也可以是int的别名
func Test[T it](b T) {
	fmt.Println(b)
}

// Test2 这里的T只能是int类型，不能是任何形式的别名
func Test2[T it2](b T) {
	fmt.Println(b)
}

// 联合约束类型。当然没有写完。下同
type unionT interface {
	int | int8 | int16 | int32 | int64 | uint
}

// 联合近似类型。
type unionT2 interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint
}

// 匿名约束
func closure[T interface{ int | ~int64 }]() {}

// 上面的匿名约束，可以简化成如下
func closure2[T int | ~int64]() {}

// 类型集合和代码
type typeCollect interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

func AnyMax[T typeCollect](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// JoinSlice T 关联的类型集合为所有实现了 fmt.Stringer 接口的类型
func JoinSlice[T fmt.Stringer](elements []T, separator string) string {
	var ret string
	for k, v := range elements {
		ret += v.String()
		if k < len(elements)-1 {
			ret += separator
		}
	}

	return ret
}

// Float 测试泛型类型断言
type Float interface {
	~float32 | ~float64
}

// NewtonSqrt 这里的类型断言是断言的动态类型，而非底层类型
func NewtonSqrt[T Float](v T) {
	var iterations int
	switch (interface{})(v).(type) {
	case float32:
		iterations = 4
	case float64:
		iterations = 5
	default:
		panic(fmt.Sprintf("unexpected type %T", v))
	}
	// Code omitted.
	fmt.Println(iterations)
}

// NewtonSqrt2 使用反射来判断类型，无疑更合理，因为这里接受的是类似基础类型（即基础类型别名）
func NewtonSqrt2[T Float](v T) {
	var iterations int
	// switch (interface{})(v).(type) {
	switch reflect.TypeOf(v).Kind() {
	case reflect.Float32:
		iterations = 4
	case reflect.Float64:
		iterations = 5
	default:
		panic(fmt.Sprintf("unexpected type %T", v))
	}
	// Code omitted.
	fmt.Println(iterations)
}
