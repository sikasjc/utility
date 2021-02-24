package string

import "unsafe"

// Mutual conversion of []byte and string should be avoided as much as possible,
// because string in Go is immutable type and the mutual conversion are copied
// by value in standard library.

/*
// slice (src/runtime/slice.go)
type slice struct {
    array unsafe.Pointer
    len   int
    cap   int
}

// string (src/runtime/string.go)
type stringStruct struct {
    str unsafe.Pointer
    len int
}

// string
func gostringnocopy(str *byte) string {
   ss := stringStruct{str: unsafe.Pointer(str), len: findnull(str)}
   s := *(*string)(unsafe.Pointer(&ss))
   return s
}
*/

// 1. The array of the []byte object points to a byte array, and the str in string is converted from *byte, so the string's str points to a byte array,
//    which is the same as []byte.
// 2. The only difference between []byte and string is that there is an additional cap attribute.
// 3. The str pointer to string is not modifiable, so every time you change the value of string, you have to reallocate memory.

// StringToBytes converts string to []byte (read only).
// Modify the []byte will be panic.
func StringToBytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	b := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&b))
}

// BytesToString converts []byte to string.
// If the original []byte is modified, the value of string is also changed.
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
