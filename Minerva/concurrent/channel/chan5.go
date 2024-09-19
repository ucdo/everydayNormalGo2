package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
)

// 写个生产者，消费者模型
// 使用goroutine和channel

// 生产者： 产生随机数
// 消费者： 计算各个数位的和

// 一个生产者，20个消费者

type pro struct {
	id    int64
	value int64
}

type result struct {
	pro *pro
	sum int64
}

// 生产者
func producer(c chan *pro) {
	// 生成随机数放在channel里
	i := 0
	for {
		x := rand.Int63()
		c <- &pro{id: int64(i), value: x}
		i++
	}
}

func consumer(c chan *pro, r chan *result) {
	for v := range c {
		sum := calc(v.value)
		temp := &result{pro: v, sum: sum}
		r <- temp
	}
}

func calc(num int64) int64 {
	sum := int64(0)
	for num > 0 {
		sum += num % 10
		num = num / 10
	}
	return sum
}

func startWork(n int, c chan *pro, r chan *result) {
	for i := 0; i < n; i++ {
		go consumer(c, r)
	}
}

func printfCalc(r chan *result) {
	for i := range r {
		fmt.Printf("calculate result: id:%d value:%d sum:%d\n", i.pro.id, i.pro.value, i.sum)
	}

}

var wg sync.WaitGroup

// TODO 实现完美优雅的退出
func main() {
	c := make(chan *pro, 20)
	r := make(chan *result, 20)
	e := make(chan string, 1)
	go func() {
		os.Stdin.Read(make([]byte, 1))
		e <- "xx"
	}()

	go producer(c)
	go startWork(20, c, r)
	go printfCalc(r)
	select {
	case <-e:
		break
	}
}
