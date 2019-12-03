// Package string has some utility functions processing string

package string

import (
	"strings"
	"testing"
)

type void struct{}

func count(str string, chars string) int {
	res := 0
	for _, ch := range chars {
		res += strings.Count(str, string(ch))
	}
	return res
}

func TestRandomString(t *testing.T) {
	type args struct {
		length   int
		digitNum int
		specNum  int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "8, 1, 1",
			args:    args{8, 1, 1},
			wantErr: false,
		},
		{
			name:    "8, 0, 0",
			args:    args{8, 0, 0},
			wantErr: false,
		},
		{
			name:    "8, 4, 4",
			args:    args{8, 4, 4},
			wantErr: false,
		},
		{
			name:    "8, 4, 5",
			args:    args{8, 4, 5},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RandomString(tt.args.length, tt.args.digitNum, tt.args.specNum)
			if (err != nil) != tt.wantErr {
				t.Errorf("RandomString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}
			digit := count(got, digits)
			special := count(got, specials)
			if special < tt.args.specNum {
				t.Errorf("RandomString() has not enough specical characters, want >= %v, got %v", tt.args.specNum, special)
				return
			}
			if digit < tt.args.digitNum {
				t.Errorf("RandomString() has not enough digit, want >= %v, got %v", tt.args.specNum, digit)
				return
			}
		})
	}
}
