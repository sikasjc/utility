package time

import (
	"fmt"
	"testing"
	"time"
)

func TestNow(t *testing.T) {
	Init()
	time.Sleep(2 * time.Second)
	fmt.Println("Now")
	for i := 0; i < 5; i++ {
		time.Sleep(2 * time.Millisecond)
		fmt.Println(Now())
	}
	fmt.Println("time.Now")
	for i := 0; i < 5; i++ {
		time.Sleep(2 * time.Millisecond)
		fmt.Println(time.Now())
	}
}

func BenchmarkNow(b *testing.B) {
	Init()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Now()
	}
}

func BenchmarkTimeNow(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		time.Now()
	}
}

func BenchmarkNow2(b *testing.B) {
	Init()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Now()
		}
	})
}
func BenchmarkTimeNow2(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			time.Now()
		}
	})
}
