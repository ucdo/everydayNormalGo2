# GO

## 最基础
```
1. package: 包
2. base: 每个可执行程序有一个main函数
3. base: 优先级 init() > main
4. base: 首字母大写即外部可见
5. base: 首字母小写即包内可见
6. string: 字符串拼接 直接用  + 
7. base: 变量申明 var identifier type
8. base: 变量申明 identifier := value // 好像可以自动类型推断
9. base: 只定义变量不复制，各有各的初始值
10. const: 常量 ： const identity [type] = value
11. iota : 初始0，一直 ++，有值跳过，无值赋值; 嗯这里iota的值，怪得很
12. auto: 自动类型推断
13. base: 无三目运算符
14. 字符串是定长的
15. 数组是定长的
16. 数组申明方式 var a = []int{1,3,}  b := []int{1}
17. 数组申明长度但是不赋值，会以数组类型初始化
18. &a  表示取地址
19. 指针  var identity *type
20. 不能对指针随便赋值，只能放地址
21. 指针也是有地址的
22. 如果数组指定了类型或者长度，任意一项不匹配都不行
23. 闭包
24. base: 数组是定长的，切片是随意的
25. [size]type: 指针数组，存放指针的数组
26. pointer: 多重指针，即指向指针的指针。 var a **int // 二重指针
27. struct: s结构体指针  var c *Book = &book
28. pointer: 自动类型推断的指针 c := &book
29. func: 函数接收参数时，指定为 func c(c *Book){} // 这里要结构体类型
30. slice: 是数组的抽象
31. slice: 不指定大小的数组就是切片
32. slice: var inden []type
33. slice: var s []type = make([]type,len) 
34. slice: slice := make([]type,len)
35. make([]T, length, capacity) // 类型，长度，容量
36. slice: a := []int{1,2,3} s := a[start:end] // end > cap(a) 会报错
37. slice: s := []int{} var s []int ,分别用 s == nil 判断，前一个不为nil 
38. printf: %x 以16进制输出， %d以10进制输出
39. copy: copy(target, origin)（只针对slice），当target的cap小于origin，只copy target.cap个
40. range: for k,v = range f{} // f可以是array,slice,channel,map,
41. range: for _,v = range "你好"
42. map: map是引用类型，一个地方修改会影响到所有地方
43. map: 跟slice一样，会自动扩容
44. map: 始终建议用make来初始化
45. map: 杰哥： 一般不用多维，多维一般用struct
46. map: 多维数组看有没有初始化可以用 == nil来判断，为了程序的安全性
47. strconv: 类型转换
48. 类型转换的时候一定要处理error，不然我估计要崩溃
49. interface: 接口太自由了
50. interface: 任何其他类型实现了这些方法就是实现了这个接口。必须全部实现
51. goroutine: 所有goroutine共享一个地址空间
52. goroutine: 异步的，普通方法是同步的
53. goroutine: go func_name(...param)
53. channel： c := make(chan int)
54. channel： 在make的时候可以设置是否有缓冲区
55. channel： 如果无缓冲区，那么读写都是原子性的。意思会阻塞其他协程写
56. channel： 无缓冲就会导致竞争，这没问题吧？所以高并发怎么玩？
57. channel： 有缓冲区，如果没写满，就会一直写，不会竞争
58. channel： 有缓冲区，如果写满了，就要等有人来读，不然一直等，可能 deadlock
59. 逃逸： 逃逸是指生命周期从短生命周期提升到长生命周期。比如将func里面局部变量的地址赋给全局变量
60. 生命周期：生命周期长的变量是分配在堆上的；相反就是分配在栈上
61. 类型转换： 高精度转低精度会导致数值改变或者精度丢失 f := 3.14 i := int(f) //3
```

### 并发
#### 关键字 go chan
#### 怎么玩的？

```go
package main

func main() {
	// goroutine简单输出
	go test(1) // 1
	test(2)    // 2
}

func test(a int) {
	println(a)
}

```
这里需要注意的是，因为协程是异步的，如果 `2` 也用 `go test(2)`，就会导致main退出。然后就不执行

#### 无缓冲chan

#### 带缓冲channel死锁演示
```go
package main

// 演示： deadlock in chan with buffer
// fatal error: all goroutines are asleep - deadlock!
//
// goroutine 1 [chan send]:
// main.main()
//
//	F:/go/src/hello/chanDL.go:7 +0x47
//
// exit status 2

// 为什么？ 因为缓冲区满了，但是这时没人来读，就导致阻塞，就死锁了
func main() {
	c := make(chan int, 1)
	c <- 1
	c <- 2
	println(<-c)
	println(<-c)
}

```

## 关键字
```` 
break	    default         func        interface   select
case	    defer           go          map         struct
chan	    else            goto        package     switch
const	    fallthrough	    if	        range       type
continue    for	            import      return      var
````

## 预定义标识符
```
append  bool    byte    cap     close   complex     complex64   complex128  uint16
copy	false   float32 float64 imag    int         int8        int16       uint32
int32   int64	iota	len     make	new     nil     panic	uint64
print	println real	recover string  true	uint	uint8	uintptr
```

## 执行/编译
``` 
1. go run file.go
2. go build file.go
3. ./file.exe
4. go run .
```

## package

``` 
1. 类似于其他语言的模块或者库的概念
2. 封装、模块化、代码重用
3. 单独编译
4. 每个包都对应一个独立的名字空间
5. 首字母大写即可导出
6. 每个包在解决依赖的前提下，以导入声明的顺序初始化，每个包只会被初始化一次
```

## 作用域

``` 
1. 声明语句的作用域是指源代码中 可以有效使用这个变量名的范围
2. 不要把作用域和生命周期混为一谈
3. 它属于编译是的属性
4. 当然，有全局的，也有局部的
5. if 的作用域就在if语句块内，在外面就没了
6. tmp2.go 演示了if的作用域
7. 
```

## 生命周期

``` 
1. 一个变量的声明周期指的是程序在运行时变量存在的有效时间段
2. 在声明周期内可以被其他程序引用
3. 运行时概念

```

## 编译器

``` 
1. 当编译器遇到一个名字应用时，它会对其定义进行查找，查找的过程从最内层的词法域向全局的作用域进行。
2. 编译成linux包：
    如果要编译到linux上需要执行：
    SET CGO_ENABLE=0
    SET GOOS=linux
    SET GOARCH=amd64
    要在编译成windows上的可执行文件
    SET GOOS=windows
3. 
```

## 基础类型

1. 0开头的8进制
2. 0x或者0X开头的16进制

### 整型 int

```
1. int8 int16 int32 int64
2. uint8 uint16 uint32 uint64
3. byte 等价于 uint8 
4. rune 等价于 uint32
5. uintptr 这个是跟C交互会用到
6. int8 = 2^(8-1) ~ 2^(8-1)-1 
7. uint8 = 0 ~ 2^8 - 1 = 0 ~ 255
```

### 浮点型 float

``` 
1. float32 float64
2. 在math包里
```

### 复数 complex

### 布尔型 bool

### 字符串 string

``` 
1. 一个不可改变得字节序列
2. s[i] i >= len(s) : panic
3. 非ASCII字符的UTF8编码，会占用多个字节，比如 s := "国家" len(s) // 6
4. s1 := s[0,3] // 国  s2 := [1,4] //  ���
5. 4个重要包： bytes,strings,strconv,unicode
6. bytes.Buffer
7. strconv.Itoa // int to string
8. s[i:j] 字符串切片生成新的字符串
9. 使用 for range 遍历非ASCII码的字符串（比如中文）,底层会帮助处理；for i 则不会
10. 字符串操作可以转换成[]byte 或者 []rune 类型再操作
11. 10.示例代码
    // 省略...
    s:=  "xxx"
    sx := byte(s) // 这里其实是copy
    // 省略...
```

### 字节 []byte

``` 
1. 字节切片生成新的 []byte
```

### 常量 const 

``` 
1. 在编译期完成
2. iota 关键字： 在const里面使用
3. 每遇到一个const ， iota初始化为0
4. consts 里面，每新增一行变量申明，iota加1
5. const 如果下面的变量如果不赋值，默认和上一行一致
```

## 复合类型

### 数组

``` 
1. 定长
2. 如果只指定长度，不进行操作，那么默认为0值
3. 定义的是时候要指定长度： a := [2]type{} 或者 s := [...]type{1,2,3}
4. a := [3]type{1,2,3} 清空数组  a = [3]type{}
```

### slice

``` 
1. 类似数组，但是变长
2. 一个slice由指针，长度，以及容量构成
3. 指针指向第一个slice元素对应的底层数组元素的地址（但是不一定就是数组的第一个元素，比如我从3开始切片）
4. 切片操作 s[i:j] 切片将有用 j-i个元素 0 <= i <= j <= cap(s)
5. 容量： a := [13]int{} s := a[4,7] fmt.Println(len(s),cap(s)) // 3 9
6. 真神奇： 继续使用上面的s， fmt.Println(s[:9]) // 这里将输出 a[4:]里面对应的内容
7. 如果切片操作超出cap(s)的上限将导致一个panic异常
8. s[0] = 123 // 这里修改了，然后 数组 a 对应的值也会被修改，具体位置要看slice的第一个元素的指针位置
9. 向函数传递slice意味着，底层元素也可能会被修改
10. 不能像 array 一样直接比较， 好像有个bytes.Equal，但是只能比较 []byte 类型的，其他类型需要自己写
11. var s []int  s == nil // true 
12. s = []int(nil) s== nil // true
13. s = nil s == nil // true 
14. s = []int{}  s == nil // false
15. 11-14没什么用
16. 要判断slice是否为空，只需要 len(s) == 0 来判断
17. make([]T,len,cap) []T{} make的优点是可以
18. slice appendint1 //老版本 ： 必须先检测slice底层数组是否有足够的容量来保存新添加的元素
19. slice appentint1 版本： 如果有足够空间的话，直接扩展slice（依然在原有的底层数组之上），将新添加的y元素复制到新扩展的空间，并返回slice。因此，输入的x和输出的z共享相同的底层数组。
20. 如果没有足够的增长空间的话，appendInt函数则会先分配一个足够大的slice用于保存新的结果，先将输入的x复制到新的空间，然后添加y元素。结果z和输入的x引用的将是不同的底层数组。
21. 内置的append函数更加复杂；所以不清楚新的slice和原始slice是否引用的是相同的底层数组空间。
22. 同样也不能确定在原先的slice上操作会不会影响新的slice
23. ... 变长参数slice
24. 有坑！ 要是操作slice，并返回slice的切片的时候，一定要注意切片切的地方；要么就直接make一个，然后append进去再返回
```

### map

``` 
1. 无序kv集合，检索、更新、删除都是常数的时间复杂度
2. 哈希表的引用
3. 创建 
    1. age := make(map[string]int) 
    2. age := map[string]int{
            "name1":1,
            "name2":2,
        }
    3. age := map[string]int{}
4. 跟php的array类似？ 但是 age["keyyyy"] 如果key不存在，不会报错
5. 禁止对map取地址，取也会编译报错；因为map可能会随着元素增加而重新分配更大的存储空间，从而导致之前的地址失效
6. 遍历的顺序是随机的
7. 好像，好像不能直接对map进行排序
8. v,ok := age[key]; !ok{/*值不存在*/}
9. map 之间也不能进行比较，除非nil 或者你就要手动写代码去比较
10. map的key居然可以是结构体，吐了
11. map的key用struct也无所谓，因为他又不是引用类型，改了也无所谓
```

跟PHParray的区别

``` 
1. age["keyyyy"] 如果key不存在，不会报错
2. 每次遍历的顺序，php是一致的；而go的map不确定
3. 好像不能直接排序，要把map的key放在slice里面，用 sort.Strings(s)对slice里面的值进行排序
4. 当前结构体可以嵌套其他结构体
5. 两个匿名的结构体成员吧不能有同名的成员，不然会报错
6. 在包外部，不可导出的成员，也不能通过匿名方式访问和初始化
7. 
```

## 结构体 struct

``` 
1. 可以是复合类型
2. 里面可以引用struct本身
3. 初始化可以记住struct申明的顺序结构，然后挨个初始化
4. 初始化也可以通过申明的名字来初始化
5. 未初始化的成员，会被默认初始化成对应类型的零值
6. 不能初始化其他包里面可导出struct的不可导出成员
7. type s struct{
        a int
    }
    ...
    pp := &s{a:111} // 初始化并把地址赋予给pp
    // 下面是同样的操作
    pp := new(s)
    *pp = s{a:777}
    ...
8. 指针接受者：
    1. 需要修改接收者中的值
    2. 接收者拷贝代价比较大的对象
    3. 保持一致性： 如果某个方法使用了指针接收者，那么其他方法也应该使用指针接收者
9. 值接收者：

```

### json

``` 
1. 发送和接收结构化数据的解析，类似的还有xml，protobuf
2. 标准库 encoding/json
3. 只有可导出的才会被编码
4. type f struct{
        a string 
        B string `json:"BNickname"`
    } 
5. json.Marshal(T) // 默认格式
6. json.MarshalIndent(T,"","    ") // 格式化
7. 结构体成员的json里面人如果带了 omitempty ，当改成员为空时，则不导出
```

## 函数 func

``` 
1. 指针，slice(切片)、map、function、channel等类型，实参可能会由于函数的间接引用被修改。
2. 任何进行I/O操作的函数都会面临出现错误的可能，只有没有经验的程序员才会相信读写操作不会失败
3. 居然也可以像PHP一样，把方法名赋予变量，然后通过 v() 调用
4. 分享一个骚东西
    const INSERT = 1
    var gen = map[string]func(v ...interface{})(string,[]interface{})
    // ... 省略
    func init(){
        gen = make(map[string]func(v ...interface{})(string,[]interface{}))
        gen[INSERT] = _insert
    }

    // ...
    func _insert(v ...interface{})(string,[]interface{}){
        //...
    }

    // 调用
    s := []interface{}{1,2,3}
    _,_ = gen[INSERT](s...) // 这里就相当于调用了 _insert方法
```

## 匿名函数

``` 
1. 捕获迭代变量：这里很容易出问题
2. 看个1 的例子
    var rmdirs []func()
    for _, d := range tempDirs() {
        dir := d // NOTE: necessary!
        os.MkdirAll(dir, 0755) // creates parent directories too
        rmdirs = append(rmdirs, func() {
            os.RemoveAll(dir)
        })
    }
    // ...do some work…
    for _, rmdir := range rmdirs {
        rmdir() // clean up
    }

    代码的第三行为什么要这么操作？ 因为后面的匿名函数记录的是
    运行时变量的内存地址，如果直接赋值d的话，d就是最后一次迭代的值了，
    那么所有的删除操作都变成对 最后一个文件的删除操作了

3. 为了解决2上的问题，通常引入一个同名的局部变量
    for _, dir := range tempDirs() {
        dir := dir // declares inner dir, initialized to outer dir
        // ...
    }

    go programming language 写到
4. 不仅仅是for range有这个问题，fori，go，defer都有这个问题
5. go或者defer会等循环结束之后再执行，所以4
6. 函数内的匿名函数可以访问包括返回值在内的所有变量


```

## 可变参数 -> 参数的数量不确定

```
1. sum(val ...int) 这里可以传多个int类型的参数
2. 调用者会创建一个匿名数组来装接收值
3. 怎么向sum里面传 int类型的slice？
4. s := []int{1,2,3,4} sum(s...)
```

## defer

```
1. 好用的defer
2. 一个日志记录的例子
    func bigSlowOperation(){
        defer trace("some msg")()
    }

    func trace(msg string) func(){
        timeStart := time.Now()
        log.Printf("%s is start...\n",msg)
        return func(){
            log.Printf("%s is ending, cost time %s",msg,time.Since(timeStart))
        }
    }
3. defer语句中的函数会在return语句更新返回值变量后再执行
4. 特别注意： 在fori，forr中defer会在最后才执行
5. 如果程序中发生了panic，但是还是会执行defer的函数
6. 一个函数中的多个defer函数会逆序执行
```

## panic

```
1. 运行时错误会导致panic，比如数组越界，空指针引用
2. 发生panic时，会导致程序中断-》并输出堆栈跟踪信息-》
   通常，我们不需要再次运行程序去定位问题，日志信息已经提供了足够的诊断依据（一定什么事都要记录日志）
3. 由于panic会引起程序的崩溃，因此panic一般用于严重错误，如程序内部的逻辑不一致
4. 对于大部分漏洞，我们应该使用Go提供的错误机制，而不是panic，尽量避免程序的崩溃
5. 如果程序中发生了panic，但是还是会执行defer的函数
```

## recover -> 从panic中恢复正常

```
1. 不加区分的恢复所有的panic异常，不是可取的做法
2. 你不应该试图去恢复其他包引起的panic
3. 你也不应该恢复一个由他人开发的函数引起的panic
4. 安全的做法是有选择性的recover
5. 换句话说，只恢复应该被恢复的panic异常
6. 恢复的异常占比应该尽可能低
7. >< 自己测试了，如果recover会导致一些问题
   这里导致的问题就是如果在return前面panic了
   那么return的值不会按预期的返回，因为还没到return
8. recover之后，从panic开始，后面的代码都不会执行（只是发生panic的这个func）
    为什么？
    因为是在defer里面执行的panic
```

## 方法 （OOP）
``` 
1. 封装和组合
2. 
```

## 嵌入结构体扩展类型
``` 
1. 
    import "image/color"

    type Point struct{ X, Y float64 }
    
    type ColoredPoint struct {
        Point
        Color color.RGBA
    }
2. 看段代码举个例子
    package main

    import (
        "fmt"
        "image/color"
        "math"
    )
    
    type Point struct {
        x, y float64
    }
    
    type ColorPoint struct {
        Point
        Color color.RGBA
    }
    //上面的代码里，ColorPoint里面 嵌入得有Point，但是ColorPoint并不是Point的子类；而是组合
    //概括：内嵌字段会指导编译器去生成额外的包装方法来委托已经声明好的方法
    //举个例子：
    func (q Point) Distance(p Point) {
    }
    // 然后Point内嵌到了ColorPoint里面
    // 编译器生成额外的包装方法
    func (c ColorPoint) Distance(p Point){
        return c.Point.Distance(p)
    }
3. 还可以有指针类型的内嵌
    type Cp struct {
        *Point
        Color color.RGBA
    }
    // 这种写法的话，就相当于字段和方法都会被间接引入到当前类型（但是需要通过指针去获取）
    // 例子
    a := Cp{&Point{1,2},color.RGBA{0,0,0,0}}
    a.Distance(*a.Point)
```

## 基于指针的对象方法
```
1. 约定： 如果Point类有一个指针作为接收器的方法，那么所有的
    Point的方法都必须有一个指针接收器，即使是并不需要指针接收器的函数
2. p.x = 1 这里相当于对接收器里面的数据进行了修改
    这里其实是go隐式的对p这个接收器取了地址然后再操作
3. 无论method的receiver是*T还是T，都可以对其进行指针操作
4. 这里设置接收器是*T还是T需要考虑两件事：
    1. 这个对象的本身是不是特别大，如果声明为非指针类型，调用会产生一次拷贝
    2. 如果用指针类型作为receiver，那么该指针类型始终指向同一个内存地址，无论你是否对他进行拷贝
5. nil也算是合法的接收器，参考 struct.go 里面的intList
6. 5中，要指出nil的意义，比如上面的链表，nil代表没有后续
7. 
```

## interface 接口
```
1. 接口类型是对其他类型行为的抽象和概括
2. 接口类型不会何特定的实现细节绑定在一起
3. go接口的独特之处在于他是满足隐式实现的
4. 接口也可以内嵌
5. 一个类型如果拥有一个接口所需要的所有方法，那么这个类型就实现了这个接口
6. interface{} 被称为空接口类型，可以接收任何类型的值
7. 不能对空接口进行操作
8. 接口这玩意好像还是动态分配的，类型是动态的
9. 对接口的值进行比较时，只有非常确定能比较的值才进行比较，否则有panic的可能
10. 同样，interface的值作为switch操作数或者map的key的时候，也可能panic
11. net/http 包的试用 /book/http*.go
12. 类型断言 Type Assertion : 判断接口变量里的值是否属于某个具体的类型
13. 类型分支 : 像switch case一样，根据具体的类型做相应的操作；
    例如 https://golang-china.github.io/gopl-zh/ch7/ch7-14.html中Exec操作这里
15. 接口设计: ask only for what you need（只考虑你需要的东西）
16. 并不是任何事物都需要被当作一个对象，独立的函数或者未封装的数据类型也各有他们的好处
17. 在实现interface定义的方法是，参数、类型、返回值要保持一致才能算是实现
18. 接口是引用类型
19. 由类型和值构成
20. 类型断言 value,ok := x.(Type) if ok{} //断言成功则获取对应的值，否则是对应类型的零值
```

## goroutine 
```
1. 跟踪调试并发程序还是很困难
2. 线性程序中形成的直觉往往会使我们误入歧途
3. 可以把gorountine 类比线程（但是并不是）
4. 一个程序启动时，主函数运行在一个独立的gorountine里，称为main rountine
5. 新的gorountine会使用go语句创建
6. 主函数返回时，所有gorountine都会被直接打断，程序退出
7. 
```

## channel

``````
1. gorountine 的通讯机制
2. ch := make(chan T) // ch就只能发送T这种类型的数据
3. 跟map一样，都是对底层数据结构的引用（复制，传参都是引用）
4. channel的零值也是nil
5. 类型相同的channel可以通过  == 比较
6. 发送和接收都用 <- 
7. close(ch) 关闭通道；后续的基于该channel的任何发送操作都会导致panic
8. close(ch)之后，还是可以接收之前已经发送成功的数据
9. 如果channel中没有数据，将给零值
10. 无缓冲channel  make(chan int) // unbuffered channel
11. 无缓冲的channel： 发送者goroutine发送之后被阻塞，等到另一个goroutine从相同的channel里接收了数据，然后继续执行各自后续的语句；接收者先发生，那么接收者也会被阻塞，直到另一个goroutine向相同的channel里面发送数据
12. 无缓冲channel：发送和接收都会导致两个goroutine做一次同步操作做；所以也被称之为 同步channels
13. x事件在y事件之前发生 happends before
14. 串联的channels（Pipeline） 管道
15. x,ok := <-ch  !ok 代表通道已经关闭了并且没有值可以接收
16. 重复关闭一个channel、试图关闭为nil的channel，都会导致panic异常
17. 关闭channel会触发广播机制
18. 单向channel： chan<- T 只发送； <-chan T 只接收
19. 单项channel： 关闭操作只用于断言不再发送新的数据；所以发送方才会调用；接收方close会编译错误
20. chan int 可以转为 chan<- int 或者 <-chan int ；反之则不行
21. 带缓存的channel：
22. ch := make(chan T, cap)
23. 发送：从队尾写入；读取：从队头读取
24. 满了还发则阻塞；空的还读也阻塞
25. 这玩意儿还不能当成队列使用（slice可以当成队列使用）
26. 多个goroutine并发向同一个channel发送数据，或者从同一个channel接收数据都是常见的
27. 泄露的goroutine不会被gc
28. goroutine泄露：
    func query() string {
        ch := make(chan string) // channel without buffer
        go func(){ch <- request("xxx.com")}()
        go func(){ch <- request("xxx.com")}()
        go func(){ch <- request("xxx.com")}()
        return <-ch
    }
29. 纠正一个错误： 从关闭的channel里面读取数据，并不会panic，也不会报错，只会返回chan T对应的零值
30. channel当成同步机制的案例：
	回想一下 <并发非阻塞缓存>
	go e.call 和 go e.deliver

	func (e *entry) call(f Func, key string) {
        e.res.value, e.res.err = f(key)
        close(e.ready)
    }

    func (e *entry) deliver(resp chan<- result) {
        <-e.ready //  为什么这里会被阻塞？ 结合上面的代码来看，因为ready是个无buffer的channel。读空阻塞，call写入之后就可以了
        resp <- e.res
    }
``````



## 并发

``````
1. 什么是并发？不能确定x,y执行的先后顺序
2. 在并发的条件下，函数能够正确地工作，称之为并发安全
3. 一定要避免数据竞争
4. 不要使用共享数据来通信，而是使用通信来共享数据
5. go的锁不能重入
	举个例子：
    var mu Sync.Mutex
	func a(){
		mu.Lock()
		defer mu.Unlock()
	}
	
	func b(){
		mu.Lock()
		defer mu.Unlock()
		a()
	}
	// 直接死锁了
6. 互斥锁 xx := sync.Mutex  xx.Lock() defer xx.Unlock()
7. 读写锁 sync.RWMutex  写不加锁，读会加锁
	1. 读锁 RLock() defer RUnLock()
	2. 写锁 Lock() defer Unlock()
8. 所有并发问题都可以用一致的、简单的既定模式来规避
   1. 如果可能，将变量限定在goroutine内部
   2. 如果是多个goroutine都要访问的变量 ，使用互斥条件访问
9. 有数据竞争的时候记得要加锁；并且加合适的锁。
``````



## goroutine和线程

``````
1. 动态栈
	1. 每个os的线程都有一个固定大小的内存块（一般是2M）来当作栈
	2. 这个栈会来用存储正在被调用或者挂起（指在调用其他函数时）的函数内部变量
	3. goroutine会以初始2kb，并会根据需要动态扩容，最大1gb
	4. m:n
``````



## package包

``````
1. 匿名包也会被编译，因为有时候会用到他的副作用，比如init之类的
2. 尽可能让包的命名简短有描述而无歧义
3. 包名一般使用单数形式
4. go get -u 只能简单地保证每个包是最新版本；并不适合发布程序
5. 对于本地依赖地包版本更新需要谨慎并且可控
6. 
``````
## GOPATH
```
1. export GOPATH=$HOME/gobook // linux上
   go env -w GOPATH="E:\\gozoom" 或者 setx GOPATH "C:\path\to\dir"
2. src 用于存储源码
3. pkg用于保存编译后的包的目标文件
4. bin用于保存编译后的可执行程序
```

## 编译
```
1. go build 和 go install 都不会重新编译没发生变化的包
2. go install 会把编译结果安装到啊GOOS和GOARCH对应的目录
```

## 测试
```
1. _test.go 作为后缀名
2. 不会被编译
3. *_test.go 包含三类：测试函数，基准函数，示例函数
4. 测试函数：以 Test 为函数名前缀的函数；用于测试程序的一些逻辑行为是否正确；
    go test会调用这些函数并生成 PASS 或者FAIL   
5. 基准函数：以Benchmark为前缀的函数；用于衡量一些函数的性能；
    go test会多次调用并取平均值
6. 示例函数：以Example为函数名前缀的函数；提供一个由编译器保证正确性的示例文档
7. go test 会执行所有的 *_test
8. go test -v 会显示所有的测试用例以及测试用例的情况
9. go test -v -run="t1|t2" 只跑这两个指定的测试用例
10. 随机测试也很有必要
11. log.Fatal 和 os.Exit会导致代码提前退出；这类函数应该放在main函数里
12. 如果真的又panic，在测试代码里应该recover然后当成错误来记录和处理
13. 测试要避免不良的行为，例如更新生产数据库或者信用卡消费行为等
14. go test并不会并发地执行多个测试
15. 在包内的测试代码一般称为包内测试
16. 但是有好几个包一起测试的话，就需要用外部测试；也可以避免循环依赖
17. 技巧：在一个外部测试包里要使用一个内部测试里面的测试，但是没法访问非导出的方法
    这时，可以在 export_test.go 中 导出用于测试。  var IsSpace = isSpace 
18. go希望测试者完成发部分的工作
19. 好的测试不应该引发其他无关的错误信息。
20. 最理想的情况是不看代码就能根据错误信息定位问题
21. 好的测试不应该在遇到一点小错误就立刻退出测试。
22. 应该尝试和记录更多的相关错误信息。
23. 测试资源占用并生成热力图 
```

## 错误

``` 
1. 虽然Go有各种异常机制，但这些机制仅被使用在处理那些未被预料到的错误，
   即bug，而不是那些在健壮程序中应该被避免的程序错误
2. Go中大部分函数的代码结构几乎相同，首先是一系列的初始检查，
    防止错误发生，之后是函数的实际逻辑。   
3. 
```

## 类型转换

### 隐式转换

```
var f float64 = 3 + 0i
f = 2
f = 1e123
f = 'a'
上面的代码相当于
var f float64 = float64(3 + 0i)
f = float64(2)
f = float64(1e123)
f = float64('a')
```

## 格式化输出

```
1. %o %#o   分别以八进制以及带符号的八进制输出
2. %x %#x   十六进制以及带符号的十六进制
3. %c %q    分别输出单字符以及带引号的单字符
4. %b %#b   以二进制得形式输出
5. %T       打印类型
6. %v       暂时还不是很确定
7. %t       打印bool类型
8.%+v	    啥子类型都能打印
```

## 性能
```
1. 过早的优化是万恶之源
2. 过早的抽象也会导致代码复杂
3. 
```

## new
``` 
1. new
2. new返回的是一个指针
3. new会初始化该类型为各自的零值
4. new:对于 struct或者非引用类型
5. e.g.
    tpye A struct {
        a int
    }
    
    ...
    a := new(A)
    b := &A{a:0}
    这俩效果是一样的
6. 不过要是由多个参数都需要初始化成0值，则直接用new比较爽
7. 如果是部分需要初始化成特殊值，用 &A{}好一点 
```

## make 
``` 
1. make
2. make只能用来初始化slice,map,chan这三个数据类型
```

## reflect
```
1. 提供了在运行时检查、修改和构建变量的能力
2. reflect.TypeOf:返回以一个reflect.Type对象。该对象表示其参数的动态类型
    exp:
    var x int = 1
    fmt.Println("type:",reflect.TypeOf(x)) // type: int
    fmt.Println("value:", reflect.ValueOf(x)) // value: 1
3. reflect.Indirect
4. v := reflect.TypeOf(x) v.Name() v.Kind()
5. 对于复合类型 v.Name() 是空
6. reflect.TypeOf(x).Kind() 可以用switch case 来获取类型
```