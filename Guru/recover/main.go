package main

import "fmt"

func main() {
	fmt.Println(get(5))
}

func get(index int) (ret int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("something went wrong:", r)
			// 需要注意的是，这里如果要对返回值进行赋值，要写成 变量名 类型 ,然后对变量名进行赋值
			ret = -1
		}
	}()

	a := [2]int{1, 2}
	ret = a[index]
	return
}
