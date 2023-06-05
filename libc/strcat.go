package libc

import "unsafe"

func StrCat[T comparable](dst *T, src *T) *T {
	s := StrSlice(src)
	n := StrLen(dst)
	d := unsafe.Slice(dst, n+len(s)+1)
	d = d[n:]
	copy(d, s)
	var zero T
	d[len(d)-1] = zero
	return dst
}

func StrCatS[T comparable](dst []T, src []T) []T {
	dst = StrSliceS(dst)
	src = StrSliceS(src)
	n := len(dst)
	dst = dst[:n+len(src)+1]
	copy(dst[n:], src)
	var zero T
	dst[len(dst)-1] = zero
	return dst[:len(dst)-1]
}

func StrCatStr(dst *byte, src string) *byte {
	n := StrLen(dst)
	d := unsafe.Slice(dst, n+len(src)+1)
	d = d[n:]
	copy(d, src)
	d[len(d)-1] = 0
	return dst
}

func StrCatStrS(dst []byte, src string) []byte {
	dst = StrSliceS(dst)
	n := len(dst)
	dst = dst[:n+len(src)+1]
	copy(dst[n:], src)
	dst[len(dst)-1] = 0
	return dst[:len(dst)-1]
}
