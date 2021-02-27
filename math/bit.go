package math

// IsPowerOfTwo ...
// Note that 0 is incorrectly considered a power of 2 here.
func IsPowerOfTwo(x uint) bool {
	return x&(x-1) == 0
}
