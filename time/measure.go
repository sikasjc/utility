// Package time has some utility functions for timing.
package time

import (
	"log"
	"time"
)

// You can track the execution time of a complete function call with this one-liner,
// which logs the result to the standard error stream.
// example:
// func foo() {
//     defer Duration(Track("foo"))
//     // Code to measure
// }

// Track the start time
func Track(msg string) (string, time.Time) {
	return msg, time.Now()
}

// Duration computes the duration from start time
func Duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}
