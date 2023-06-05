package libc

import "unsafe"

// MemSet is a generic implementation of C memset.
func MemSet[T any](p *T, v T, sz int) {
	MemSetS(unsafe.Slice(p, sz), v)
}

// MemSetS is a generic implementation of C memset that accepts a slice.
func MemSetS[T any](dst []T, v T) {
	for i := range dst {
		dst[i] = v
	}
}
