# 雅典娜
## 泛型(多态)
1. 语法： func Foo\[T any,...\](parameter) returnType
2. 泛型 实验包 golang.org/x/exp/constraints
3. 比如下面的多态代码
   ```
   func Add[constraints.Integer](a,b T) T{
        return a + b
   }
   ```
   调用：
   s := Add[int]\(1,2)  
   这里的调用即是实例化了这个泛型，确定了**T**的类型是**int**

   [泛型结构体](./generics/anyType/anyStruct.go#L11)
4. ~底层类型， 这里的泛型是指 包含所有当前底层类型以及底层类型的别名  
   举个例子：  
   type myInt int
   type it interface{
        ~int // 这里就包含了上面定义的 myInt
   }

   type it2 interface{
        int // 这里就只能匹配到定义的时候是int。但是别名是int的不行
   }
5. 文章作者说： [**注意如果接口中添加了上文提到的三种元素之一，该接口就只能作为约束使用，不能作为一般类型来使用了。**](./generics/anyType/anyStruct.go#L83)
6. [*调用接口方法*](./generics/anyType/anyStruct.go#L114)
7. 在编译期就会确定泛型的类型。所以反射也是只能查到对应的底层类型
8. 要查询底层类型，只能用反射才准确
9. ***需要在实践中多多练习***
10. [泛型思维导图](../LuXun/泛型.xmind)
## 反射