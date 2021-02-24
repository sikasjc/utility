// Package time has some utility functions for timing.
package time

import (
	"fmt"
	"time"
)

// example:
// func foo() {
//     defer func(){
//         fmt.Println(Duration(Track("foo")))
//     }
//     // Code to measure
// }

// Track the start time
func Track(msg string) (string, time.Time) {
	return msg, time.Now()
}

// Duration computes the duration from start time
func Duration(msg string, start time.Time) string {
	return fmt.Sprintf("%v: %v", msg, time.Since(start))
}
