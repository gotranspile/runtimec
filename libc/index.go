package libc

import "unsafe"

// IndexS returns first index of value v in slice s. It returns -1 if v is not in s.
func IndexS[T comparable](s []T, v T) int {
	for i, c := range s {
		if c == v {
			return i
		}
	}
	return -1
}

// LastIndexS returns last index of value v in slice s. It returns -1 if v is not in s.
func LastIndexS[T comparable](s []T, v T) int {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == v {
			return i
		}
	}
	return -1
}

// MemIndex is similar to IndexS, but accepts a pointer and a size.
func MemIndex[T comparable](p *T, v T, sz int) int {
	// https://github.com/bminor/musl/blob/master/src/string/memchr.c
	if p == nil || sz == 0 {
		return -1
	}
	return IndexS(unsafe.Slice(p, sz), v)
}

// MemLastIndex is similar to LastIndexS, but accepts a pointer and a size.
func MemLastIndex[T comparable](p *T, v T, sz int) int {
	// https://github.com/bminor/musl/blob/master/src/string/memrchr.c
	if p == nil || sz == 0 {
		return -1
	}
	return LastIndexS(unsafe.Slice(p, sz), v)
}
