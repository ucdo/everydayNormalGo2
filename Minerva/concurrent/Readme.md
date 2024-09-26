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
4. GMP *M:N* M个goroutine：N个OS线程 具体的协程还是要落实到具体的线程去完成 Google大佬写好的线程池
5. runtime.GOMAXPROCS 1.5版本之后默认是跑满所有CPU
6. 不需要跑满所有CPU的手动配置，比如日志库之类的
7. 通信用channel
8. Goroutine的特点
   1. 一个goroutine对应一个函数 
   2. main函数也是goroutine
   3. 当goroutine执行的函数return了，goroutine就结束了
   4. main函数结束，由它启动的goroutine也结束咯


### channel
1. channel是goroutine之间通信的连接
2. FIFO 队列
3. 引用类型
   1. var x chan Type
   2. 跟map，slice一样，需要用make之后才能使用。 ch = make(chan Type,[cap])
4. [channel操作](./channel/chan1.go)
   1. 发送：ch<-
   2. 接收：x := <-ch
   3. 关闭：close(ch) 
      1. 关闭后仍可取值，取完之后返回对应类型零值
      2. 不能重复关闭
      3. 关闭后不能再发送
5. 有无缓冲区
   1. 无缓冲区：同步，必须有接收才能发送，否则阻塞
   2. 有缓冲区：异步，超过容量就阻塞
      1. 怎么避免超过容量还一直往里面写（当写的内容不重要时：）
      ```
         // 这样写：对的         
         select {
         case l.logChan <- logData:
         default:
         }
         // 这样写：错的
         //l.logChan <- logData
      ```
6. 优雅取值
   1. v,ok := <- ch 能判断通道是否关闭
   2. for v := range ch{}
7. 单向通道
   1. 只能对通道进行接收操作：ch <-chan int
   2. 只能对通道进行发送操作：ch chan<- int
8. [select 多路复用](./channel/chan6_select.go)
   1. 同一时刻可以对多个通道做发送和接收操作
9. channel是线程安全的


### [sync](./sync) 
1. 需要goroutine之间协同时
2. 多个goroutine操作同一个全局变量时，存在数据竞争
3. 互斥锁
   1. 场景: 数据竞争
   2. var mtx sync.Mutex
   3. defer mtx.Unlock() mtx.Lock()
4. 读写锁
   1. 场景: 读多写少
   2. var rwLock sync.RWLock
   3. 读锁: defer rwLock.RUnlock() rwLock.RLock()
   4. 写锁: defer rwLock.Unlock() rwLock.Lock()
5. sync.Map{}
   1. 高并发下直接用map会有问题，除非你自己加锁
   2. sync.Map{} 是开箱即用，不用make，线程安全的
6. sync.Once
   1. var once sync.Once
   2. once.Do(fn)
   3. [如果fn是带参的](./sync/sync3_closure.go),就这样玩
7. atomic