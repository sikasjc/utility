package pool

import (
	"math/bits"
	"sync"

	"github.com/sikasjc/utility/math"
)

type Memory struct {
	p []sync.Pool
}

func NewMemory(maxSize uint) *Memory {
	m := &Memory{p: make([]sync.Pool, maxSize)}
	for i := uint(0); i < maxSize; i++ {
		size := 1 << i
		m.p[i].New = func() interface{} {
			return make([]byte, 0, size)
		}
	}
	return m
}

func calcIndex(size uint) int {
	if size == 0 {
		return 0
	}
	if math.IsPowerOfTwo(size) {
		return bits.Len(size) - 1
	}
	return bits.Len(size)
}

// Alloc ...
func (m *Memory) Alloc(size uint, cap uint) []byte {
	c := size
	if cap > size {
		c = cap
	}
	res := m.p[calcIndex(c)].Get().([]byte)
	res = res[:size]
	return res
}

// Free ...
func (m *Memory) Free(buf []byte) {
	size := uint(cap(buf))
	if !math.IsPowerOfTwo(size) {
		return
	}
	buf = buf[:0]
	m.p[bits.Len(size)].Put(buf)
}
