##  并发

1. 并发和并行的区别
   1. 并发：和自己两个女朋友聊天（同一时间段同事做多个事情）
   2. 并行：两个人各自和自己的女朋友聊天（同一时刻同事做多个事情）


### 进程、线程、协程

### goroutine
1. go func(){}() // go funcNane()
2. wg sync.WaitGroup // 实际上是计数器
   1. wg.Add() 
   2. defer wg.Done() 
   3. wg.Wait() 
3. 可增长栈
   1. 默认2kb 可增长
4. GMP M个goroutine：N个OS线程 具体的协程还是要落实到具体的线程去完成 Google大佬写好的线程池
5. runtime.GOMAXPROCS 1.5版本之后默认是跑满所有CPU
6. 不需要跑满所有CPU的手动配置，比如日志库之类的
7. 通信用channel

### channel
1. channel是goroutine之间通信的连接
2. FIFO 队列
3. 引用类型


### sync