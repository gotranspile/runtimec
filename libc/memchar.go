package libc

import "unsafe"

// Reference implementation:
// https://github.com/bminor/musl/blob/master/src/string/memchr.c
// https://github.com/bminor/musl/blob/master/src/string/memrchr.c

// MemCharP is the same as MemChar, but accepts and returns unsafe.Pointer.
func MemCharP(p unsafe.Pointer, v byte, sz int) unsafe.Pointer {
	return unsafe.Pointer(MemChar((*byte)(p), v, sz))
}

// MemChar is a generic implementation of C memchr.
func MemChar[T comparable](p *T, v T, sz int) *T {
	i := MemIndex(p, v, sz)
	if i < 0 {
		return nil
	}
	s := unsafe.Slice(p, sz)
	return &s[i]
}

// MemRCharP is the same as MemChar, but accepts and returns unsafe.Pointer.
func MemRCharP(p unsafe.Pointer, v byte, sz int) unsafe.Pointer {
	return unsafe.Pointer(MemRChar((*byte)(p), v, sz))
}

// MemRChar is a generic implementation of C memchr, but searches the value from the end.
func MemRChar[T comparable](p *T, v T, sz int) *T {
	i := MemLastIndex(p, v, sz)
	if i < 0 {
		return nil
	}
	s := unsafe.Slice(p, sz)
	return &s[i]
}
