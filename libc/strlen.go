package libc

import (
	"strings"
	"unsafe"
)

// Reference implementation:
// https://github.com/bminor/musl/blob/master/src/string/strlen.c
// https://github.com/bminor/musl/blob/master/src/string/strnlen.c

// StrLen is a generic version of C strlen.
//
// It accepts a pointer to the first array element and advances it until it reaches a zero value.
// Position of this zero terminator is returned, which equals to a length of the array/string.
//
// Additionally, function checks for a nil pointer and returns 0 length in this case.
//
// The behavior is undefined if p is not a pointer to a null-terminated array.
//
// Deprecated: Unsafe, use StrNLen or StrLenS.
func StrLen[T comparable](p *T) int {
	if p == nil {
		return 0
	}
	var zero T
	// TODO: expose MemLen() (uintptr, bool) from Allocator to get alloc length in best-effort manner
	sz := 0
	for *p != zero {
		p = PtrAdd(p, 1)
		sz++
	}
	return sz
}

// StrNLen is a generic version of C strnlen. It's the same as StrLen, but accepts a length of the buffer.
//
// It is safer than using StrLenS and unsafe.Slice directly, because it checks for a zero pointer before converting to a slice.
func StrNLen[T comparable](p *T, max int) int {
	if p == nil || max == 0 {
		return 0
	}
	return StrLenS(unsafe.Slice(p, max))
}

// StrLenS is similar to StrLen, but accepts a slice. It still looks for a zero value as a terminator in the slice.
func StrLenS[T comparable](p []T) int {
	var zero T
	i := IndexS(p, zero)
	if i < 0 {
		return len(p)
	}
	return i
}

// StrSlice uses StrLen to determine the length of a zero-terminated string, and returns it as a slice.
//
// This method is only suitable for reading from a string. For writing to a string buffer, see StrSliceS or StrNSlice.
//
// Deprecated: Unsafe, use StrSliceS or StrNSlice.
func StrSlice[T comparable](p *T) []T {
	if p == nil {
		return nil
	}
	n := StrLen(p)
	return unsafe.Slice(p, n+1)[:n]
}

// StrSliceS uses StrLenS to determine the length of a zero-terminated string in a slice, and returns a new slice.
func StrSliceS[T comparable](p []T) []T {
	if cap(p) == 0 {
		return nil
	}
	n := StrLenS(p)
	return p[:n]
}

// StrNSlice uses StrLen to determine the length of a zero-terminated string, and returns it as a slice.
//
// It uses max as a capacity of the returned slice, making it suitable for writing to a string buffer.
func StrNSlice[T comparable](p *T, max int) []T {
	if p == nil || max == 0 {
		return nil
	}
	s := unsafe.Slice(p, max)
	n := StrLenS(s)
	return s[:n]
}

// StrSliceStr is similar to StrSlice, but accept a string.
func StrSliceStr(s string) string {
	if i := strings.IndexByte(s, 0); i >= 0 {
		s = s[:i]
	}
	return s
}
