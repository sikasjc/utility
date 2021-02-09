// Package string has some utility functions processing string
package string

import (
	"fmt"
	"math/rand"
	"time"
)

var digits = "0123456789"
var specials = "~=+%^*/()[]{}!@#$?|"
var all = "ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
	"abcdefghijklmnopqrstuvwxyz" +
	digits + specials

// RandomString generates a random ASCII string with at least `digitNum` digit(s) and `specNum` special character(s).
func RandomString(length, digitNum, specNum int) (string, error) {
	rand.Seed(time.Now().UnixNano())
	if length < digitNum+specNum {
		return "", fmt.Errorf("the number of digits and special characters must be less than or equal to the total length")
	}
	buf := make([]byte, length)
	for i := 0; i < digitNum; i++ {
		buf[i] = digits[rand.Intn(len(digits))]
	}
	for i := digitNum; i < digitNum+specNum; i++ {
		buf[i] = specials[rand.Intn(len(specials))]
	}
	for i := digitNum + specNum; i < length; i++ {
		buf[i] = all[rand.Intn(len(all))]
	}
	rand.Shuffle(len(buf), func(i, j int) {
		buf[i], buf[j] = buf[j], buf[i]
	})
	return string(buf), nil
}
