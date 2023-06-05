package libc

import "unsafe"

// Reference implementation:
// https://github.com/bminor/musl/blob/master/src/string/strchrnul.c
// https://github.com/bminor/musl/blob/master/src/string/strrchr.c
// https://github.com/bminor/musl/blob/master/src/string/memrchr.c

// StrIndex is similar to C strchr, but returns an index.
//
// Deprecated: Unsafe, use StrIndexS.
func StrIndex[T comparable](p *T, v T) int {
	return StrIndexS(StrSlice(p), v)
}

// StrIndexS is a safe version of StrIndex.
func StrIndexS[T comparable](p []T, v T) int {
	n := StrLenS(p)
	var zero T
	if v == zero {
		if n == cap(p) {
			return -1
		}
		return n
	}
	return IndexS(p[:n], v)
}

// StrChar returns a pointer to the first occurrence of value in the C array/string p.
//
// The terminating null-value is considered part of the C array.
// Therefore, it can also be located in order to retrieve a pointer to the end of array/string.
//
// Deprecated: Unsafe, use StrCharS.
func StrChar[T comparable](p *T, v T) *T {
	i := StrIndexS(StrSlice(p), v)
	if i < 0 {
		return nil
	}
	s := unsafe.Slice(p, i+1)
	return &s[i]
}

// StrCharS is a safe version of StrChar.
func StrCharS[T comparable](p []T, v T) []T {
	i := StrIndexS(p, v)
	if i < 0 {
		return nil
	}
	if i >= len(p) {
		p = p[:cap(p)]
		return p[i : i+1]
	}
	return p[i:]
}

// StrLastIndex is similar to C strrchr, but returns an index.
//
// Deprecated: Unsafe, use StrLastIndexS.
func StrLastIndex[T comparable](p *T, v T) int {
	return StrLastIndexS(StrSlice(p), v)
}

// StrLastIndexS is a safe version of StrLastIndex.
func StrLastIndexS[T comparable](p []T, v T) int {
	n := StrLenS(p)
	var zero T
	if v == zero {
		if n == cap(p) {
			return -1
		}
		return n
	}
	return LastIndexS(p[:n], v)
}

// StrRChar returns a pointer to the last occurrence of value in the C array/string p.
//
// The terminating null-value is considered part of the C array.
// Therefore, it can also be located to retrieve a pointer to the end of array/string.
//
// Deprecated: Unsafe, use StrRCharS.
func StrRChar[T comparable](p *T, v T) *T {
	i := StrLastIndexS(StrSlice(p), v)
	if i < 0 {
		return nil
	}
	s := unsafe.Slice(p, i+1)
	return &s[i]
}

// StrRCharS is a safe version of StrRChar.
func StrRCharS[T comparable](p []T, v T) []T {
	i := StrLastIndexS(p, v)
	if i < 0 {
		return nil
	}
	if i >= len(p) {
		p = p[:cap(p)]
		return p[i : i+1]
	}
	return p[i:]
}
