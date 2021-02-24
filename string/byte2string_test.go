package string

import (
	"bytes"
	"strings"
	"testing"
)

func Benchmark_StringToBytes(b *testing.B) {
	s := "Golang"
	for i := 0; i < b.N; i++ {
		StringToBytes(s)
	}
}

func Benchmark_toBytes(b *testing.B) {
	s := "Golang"
	for i := 0; i < b.N; i++ {
		_ = []byte(s)
	}
}

func Benchmark_BytesToString(b *testing.B) {
	s := []byte("Golang")
	for i := 0; i < b.N; i++ {
		BytesToString(s)
	}
}

func Benchmark_toString(b *testing.B) {
	s := []byte("Golang")
	for i := 0; i < b.N; i++ {
		_ = string(s)
	}
}

func Benchmark_StringToBytesLarge(b *testing.B) {
	s := "Golang"
	s = strings.Repeat(s, 100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringToBytes(s)
	}
}

func Benchmark_toBytesLarge(b *testing.B) {
	s := "Golang"
	s = strings.Repeat(s, 100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = []byte(s)
	}
}

func Benchmark_BytesToStringLarge(b *testing.B) {
	s := []byte("Golang")
	s = bytes.Repeat(s, 100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BytesToString(s)
	}
}

func Benchmark_toStringLarge(b *testing.B) {
	s := []byte("Golang")
	s = bytes.Repeat(s, 100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = string(s)
	}
}
