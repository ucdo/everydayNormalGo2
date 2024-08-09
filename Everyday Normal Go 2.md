# Golang

## 基础

### 变量
1. 定义:
    1. var x int = 1
    2. var x1  = 1
    3. x := 1
    4. := 只能在函数内使用

### 简单类型
1. 空值 nil
2. 整型 int(取决于操作系统), int8, int16, int32, int64, uint8, uint16...
3. 浮点 float32 float64
4. 字节 byte = uint8
5. 字符串 string
6. 布尔 ture false

### 编码
1. 默认是 UTF8编码
2. 英文或者ASCII码占一个字节
3. 中文占3个字节
4. 操作中文使用 for k,v := range rune(x){} 
5. len(x) 实际上是返回的是字节数
6. 判断是否包含UTF8字符：utf8.ValidString(x)
7. string 转 rune []rune(x)


### array
1. [数据的拷贝是值拷贝，不是地址](./Guru/array/array1.go)
2. 数组是一块连续内存，通过偏移量来获取后面的值
3. 数组的长度不能变

### slice
1. 底层还是数组，只不过是slice的struct的header指向了这个底层数组的地址
2. 用 [:] 截取数组，那么slice的底层指向原来这个数组**start**这个位置的地址。修改会影响原数组
3. [start:end] 实际上是 [start,end)
4. 切片上再切片，实际上底层数据的地址还是同一个,一个地址指向start的地方。修改同2.所描述
5. 扩容: 只要超过原来的cap，底层数组就一定会发生变化
6. 扩容策略：新cap < 1024,则扩容两倍；新cap > 1024,每次 * 1.25,直到满足需求
7. copy 深拷贝。 这里被接受复制的变量，分配了多少长度就会被分配 <= cap[src] 的变量
8. var s []int   copy(s,x)  只申请了变量，没分配空间，所以 打印s也是空的，地址是0x0
9. 没有delete，只有自己用append来删，比如删除index=1的  x = append(x[:1],x[2:]...)
10. 创建 slice的4种方式：
    1. []int{} 
    2. make([]int) 
    3. 数组切片 
    4. 切片再切片

### string
1. 字符串时不可变的。尝试修改字符串会导致编译错误
2. 通过复制一份string副本进行修改
    ```go
    package main
    import "fmt"
    func main(){
        s := "123123"
        b := []byte(s)
        b[0] = 'x'
        fmt.Println(s) // 123123
        fmt.Println(b) // x23123
    }
    ```

### for range
1. for range slice/map时,创建了一个新的变量来存储当前迭代的元素的副本

## 进阶

### reflect 反射
1. 获取类型 reflect.TypeOf(x) reflect.TypeOf(x).Kind()


### CGO

### unsafe