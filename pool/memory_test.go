package pool

import (
	"fmt"
	"testing"
)

func TestMemory(t *testing.T) {
	m := NewMemory(16)

	buf := m.Alloc(1000, 2048)
	if len(buf) != 1000 || cap(buf) != 2048 {
		t.Errorf("len(buf) want 1000, got %d, cap(buf) want 2048, get %d", len(buf), cap(buf))
	}
	m.Free(buf)

	buf = m.Alloc(1000, 953)
	if len(buf) != 1000 || cap(buf) != 1024 {
		t.Errorf("len(buf) want 1000, got %d, cap(buf) want 1024, get %d", len(buf), cap(buf))
	}
	m.Free(buf)

	buf = m.Alloc(0, 0)
	if len(buf) != 0 || cap(buf) != 1 {
		t.Errorf("len(buf) want 0, got %d, cap(buf) want 1, get %d", len(buf), cap(buf))
	}
	m.Free(buf)

	buf = make([]byte, 0, 1000)
	m.Free(buf)
}

func TestNewMemory(t *testing.T) {
	m := NewMemory(16)
	for i := 0; i < len(m.p); i++ {
		fmt.Println(i, cap(m.p[i].Get().([]byte)))
	}
}

func Test_calcIndex(t *testing.T) {
	type args struct {
		size uint
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1024",
			args: args{size: 1024},
			want: 10,
		},
		{
			name: "1024 * 512",
			args: args{size: 1024 * 512},
			want: 19,
		},
		{
			name: "1024 + 512",
			args: args{size: 1024 + 512},
			want: 11,
		},
		{
			name: "1024 + 512",
			args: args{size: 1024 + 512},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcIndex(tt.args.size); got != tt.want {
				t.Errorf("calcIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
