# 测试驱动开发

## 单元测试
1. 文件以*_test.go
2. 方法以 TestXxx(t *testing.T){} 开头
3. go test -v //跑所有文件
4. go test -run="." -v //也是跑所有
5. go test -run="split" -v // 包含这个单词的方法
6. 加-v才会显示详细信息
7. go test -run /2 -v // 跑所有用例下叫2的 [TestSplit测试示例](./unit/split_test.go)

### 覆盖率测试  *下面的命令用GoLand运行会有问题*
1. go test -cover
2. go test cover -coverprofile=c.out
3. 通过go tool生成html查看覆盖了哪些代码 go tool cover -html=c.out

## 基准测试 *benchmark*
1. BenchmarkXXX(b *testing.B){}
2. go test -bench=. \[benchmem\] // 内存占用和内存分配测试
3. 可以通过基准测试来看和优化性能
4. [性能测试比较](./unit/fib_test.go)
5. 

## Setup与TearDown

## TestMain(m *testing.M){}
1. 类似main函数

## 并行测试
1. 在上面的方法后加上 Parallel



## 示例函数

