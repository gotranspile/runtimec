package libc

import (
	"unicode/utf16"
	"unsafe"
)

// Reference implementation:
// https://github.com/bminor/musl/blob/master/src/string/stpcpy.c

// strTerm always puts a zero terminator into p. Even if it trims legit data.
func strTerm[T comparable](p []T, n int) {
	var zero T
	if n >= len(p) {
		p[len(p)-1] = zero
		return
	}
	for i := n; i < len(p); i++ {
		p[i] = zero
	}
}

// StrCopy copies zero-terminated string s into a buffer pointed by p, including zero terminator.
//
// Size of buffer p must be large enough for the string and zero terminator.
//
// Deprecated: Unsafe, use StrCopyS.
func StrCopy[T comparable](p *T, s *T) {
	str := StrSlice(s)
	dst := unsafe.Slice(p, len(str)+1)
	n := copy(dst, str)
	strTerm(dst, n)
}

// StrCopySlice copies zero-terminated string s into a buffer pointed by p, including zero terminator.
//
// Size of buffer p must be large enough for the string and zero terminator.
//
// Deprecated: Unsafe, use StrCopyS.
func StrCopySlice[T comparable](p *T, s []T) {
	s = StrSliceS(s)
	dst := unsafe.Slice(p, len(s)+1)
	n := copy(dst, s)
	strTerm(dst, n)
}

// StrCopyStr copies string s into a buffer pointed by p, including zero terminator.
//
// Size of buffer p must be large enough for the string and zero terminator.
//
// Deprecated: Unsafe, use StrCopyStrS.
func StrCopyStr(p *byte, s string) {
	s = StrSliceStr(s)
	dst := unsafe.Slice(p, len(s)+1)
	n := copy(dst, s)
	strTerm(dst, n)
}

// StrCopyS copies zero-terminated string s into a buffer p, including zero terminator.
//
// If size of the buffer is too small, the string will be trimmed.
//
// There is subtle difference between StrCopyS and StrNCopyS. Both accept slices and will stop writing if buffer is too small.
// However, StrNCopyS is allowed to omit zero terminator in this case, while StrCopyS will always add it.
// Thus, effective capacity of buffer for StrCopyS is len(p)-1, while for StrNCopyS it's len(p).
func StrCopyS[T comparable](p []T, s []T) {
	s = StrSliceS(s)
	n := copy(p, s)
	strTerm(p, n)
}

// StrCopyStrS copies string s into a buffer p, including zero terminator.
//
// If size of the buffer is too small, the string will be trimmed.
//
// For difference between StrCopyStrS and StrNCopyStrS see docs for StrCopyS or StrNCopyS.
func StrCopyStrS(p []byte, s string) {
	s = StrSliceStr(s)
	n := copy(p, s)
	strTerm(p, n)
}

// StrCopyStr16 copies string s into a UTF-16 buffer pointed by p, including zero terminator.
//
// Size of buffer p must be large enough for the string and zero terminator.
//
// Deprecated: Unsafe, use StrCopyStr16S.
func StrCopyStr16(p *uint16, s string) {
	s = StrSliceStr(s)
	s16 := utf16.Encode([]rune(s))
	dst := unsafe.Slice(p, len(s16)+1)
	n := copy(dst, s16)
	strTerm(dst, n)
}

// StrCopyStr16S copies string s into a UTF-16 buffer p, including zero terminator.
//
// If size of the buffer is too small, the string will be trimmed.
//
// For difference between StrCopyStr16S and StrNCopyStr16S see docs for StrCopyS or StrNCopyS.
func StrCopyStr16S(p []uint16, s string) {
	s = StrSliceStr(s)
	s16 := utf16.Encode([]rune(s))
	n := copy(p, s16)
	strTerm(p, n)
}
