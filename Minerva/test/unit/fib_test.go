package unit

import (
	"fmt"
	"testing"
)

func benchmarkFib(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		fib(n)
	}
}

func BenchmarkFib1(b *testing.B) { benchmarkFib(b, 1) }
func BenchmarkFib4(b *testing.B) { benchmarkFib(b, 5) }

//func BenchmarkFib2(b *testing.B) { benchmarkFib(b, 10) }
//func BenchmarkFib3(b *testing.B) { benchmarkFib(b, 20) }

func TestMain(m *testing.M) {
	fmt.Println("hello world")
	retCode := m.Run()
	fmt.Println(retCode)
}
