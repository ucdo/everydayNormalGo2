## 看看差距
1. 不会阻塞代码。以前是写了日志才能执行后续的操作
2. 他是直接再mylog 里面加了一个channel来保存
3. logDataChan chan *struct
4. 定义个结构体
5. 带容量的channel如果满了，会阻塞。（金融系统才对日志要求非常高）
6. ```go
   select {
   case ch <- msg: // 如果能写
   default: // 不能写就走这个default
   }
```