package bool

import "sync/atomic"

type AtomicBool int32

func New(ok bool) *AtomicBool {
	b := new(AtomicBool)
	if ok {
		b.Set()
	}
	return b
}

func (ab *AtomicBool) Set() {
	atomic.StoreInt32((*int32)(ab), 1)
}

func (ab *AtomicBool) SetTo(ok bool) {
	if ok {
		ab.Set()
	} else {
		ab.UnSet()
	}
}

func (ab *AtomicBool) UnSet() {
	atomic.StoreInt32((*int32)(ab), 0)
}

func (ab *AtomicBool) IsSet() bool {
	return atomic.LoadInt32((*int32)(ab)) == 1
}

func (ab *AtomicBool) IsUnSet() bool {
	return atomic.LoadInt32((*int32)(ab)) != 1
}

func (ab *AtomicBool) SetToIf(old, new bool) bool {
	var o, n int32
	if old {
		o = 1
	}
	if new {
		n = 1
	}
	return atomic.CompareAndSwapInt32((*int32)(ab), o, n)
}
