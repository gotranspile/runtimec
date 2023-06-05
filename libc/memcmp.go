package libc

import (
	"bytes"
	"unsafe"
)

// MemCmp is a generic version of C memcmp.
func MemCmp[T comparable](p1, p2 *T, sz int) int {
	if p1 == nil && p2 == nil {
		return 0
	} else if p1 == nil {
		return -1
	} else if p2 == nil {
		return +1
	}
	var zero T
	return MemCmpP(unsafe.Pointer(p1), unsafe.Pointer(p2), sz*int(unsafe.Sizeof(zero)))
}

// MemCmpP is similar to MemCmp, but accepts unsafe.Pointer.
func MemCmpP(p1, p2 unsafe.Pointer, sz int) int {
	if p1 == nil && p2 == nil {
		return 0
	} else if p1 == nil {
		return -1
	} else if p2 == nil {
		return +1
	}
	s1, s2 := unsafe.Slice((*byte)(p1), sz), unsafe.Slice((*byte)(p2), sz)
	return bytes.Compare(s1, s2)
}

// MemCmpS is similar to MemCmp, but accepts slices.
func MemCmpS[T comparable](p1, p2 []T) int {
	if len(p1) < len(p2) {
		return -1
	} else if len(p1) > len(p2) {
		return +1
	}
	if len(p1) == 0 {
		return 0
	}
	return MemCmp(&p1[0], &p2[0], len(p1))
}
