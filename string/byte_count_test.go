// Package byte has some utility functions convert a size in bytes
// to human-readable string in either SI(decimal) or IEC(binary) format.
//
// Input         ByteCountSI        ByteCountIEC
//  999            "999 B"            "999 B"
//  1000           "1.0 kB"           "1000 B"
//  1023           "1.0 kB"           "1023 B"
//  1024           "1.0 kB"           "1.0 KiB"
//  987,654,321	   "987.7 MB"         "941.9 MiB"
//  math.MaxInt64  "9.2 EB"           "8.0 EiB"
//
// Source:
// https://yourbasic.org/golang/formatting-byte-size-to-human-readable-format/

package string

import (
	"math"
	"testing"
)

func TestSI(t *testing.T) {
	type args struct {
		b uint64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "999",
			args: args{999},
			want: "999 B",
		},
		{
			name: "1000",
			args: args{1000},
			want: "1.0 kB",
		},
		{
			name: "1023",
			args: args{1023},
			want: "1.0 kB",
		},
		{
			name: "1024",
			args: args{1024},
			want: "1.0 kB",
		},
		{
			name: "987654321",
			args: args{987654321},
			want: "987.7 MB",
		},
		{
			name: "math.MaxInt64",
			args: args{math.MaxInt64},
			want: "9.2 EB",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SI(tt.args.b); got != tt.want {
				t.Errorf("SI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIEC(t *testing.T) {
	type args struct {
		b uint64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "999",
			args: args{999},
			want: "999 B",
		},
		{
			name: "1000",
			args: args{1000},
			want: "1000 B",
		},
		{
			name: "1023",
			args: args{1023},
			want: "1023 B",
		},
		{
			name: "1024",
			args: args{1024},
			want: "1.0 KiB",
		},
		{
			name: "987654321",
			args: args{987654321},
			want: "941.9 MiB",
		},
		{
			name: "math.MaxInt64",
			args: args{math.MaxInt64},
			want: "8.0 EiB",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IEC(tt.args.b); got != tt.want {
				t.Errorf("IEC() = %v, want %v", got, tt.want)
			}
		})
	}
}
