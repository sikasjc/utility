package time

import (
	"sync/atomic"
	"time"
	"unsafe"
)

var current *time.Time

func Init() {
	refresh(time.Now())
	go refreshTask()
}

func Now() time.Time {
	return *(*time.Time)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&current))))
}

func refreshTask() {
	t := time.NewTicker(time.Millisecond * 1)
	defer t.Stop()
	for {
		cur := <-t.C
		refresh(cur)
	}
}

func refresh(t time.Time) {
	atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&current)), unsafe.Pointer(&t))
}
