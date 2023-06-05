package libc

import (
	"unicode/utf16"
	"unsafe"
)

// Reference implementation:
// https://github.com/bminor/musl/blob/master/src/string/stpncpy.c

// strTermN puts zero terminator only if p has space for it.
func strTermN[T comparable](p []T, n int) {
	var zero T
	for i := n; i < len(p); i++ {
		p[i] = zero
	}
}

// StrNCopy copies zero-terminated string s into a buffer pointed by p, including zero terminator.
//
// Size of buffer p must be large enough for the string. It may omit zero terminator.
//
// Deprecated: Unsafe, use StrNCopyS.
func StrNCopy[T comparable](p *T, s *T, sz int) {
	str := StrSlice(s)
	dst := unsafe.Slice(p, sz)
	StrNCopyS(dst, str)
}

// StrNCopySlice copies zero-terminated string s into a buffer pointed by p, including zero terminator.
//
// Size of buffer p must be large enough for the string. It may omit zero terminator.
//
// Deprecated: Unsafe, use StrNCopyS.
func StrNCopySlice[T comparable](p *T, s []T, sz int) {
	dst := unsafe.Slice(p, sz)
	StrNCopyS(dst, s)
}

// StrNCopyStr copies string s into a buffer pointed by p, including zero terminator.
//
// Size of buffer p must be large enough for the string. It may omit zero terminator.
//
// Deprecated: Unsafe, use StrNCopyStrS.
func StrNCopyStr(p *byte, s string, sz int) {
	dst := unsafe.Slice(p, sz)
	StrNCopyStrS(dst, s)
}

// StrNCopyS copies zero-terminated string s into a buffer p, including zero terminator.
//
// If size of the buffer is too small, the string will be trimmed. It may omit zero terminator.
//
// There is subtle difference between StrCopyS and StrNCopyS. Both accept slices and will stop writing if buffer is too small.
// However, StrNCopyS is allowed to omit zero terminator in this case, while StrCopyS will always add it.
// Thus, effective capacity of buffer for StrCopyS is len(p)-1, while for StrNCopyS it's len(p).
func StrNCopyS[T comparable](p []T, s []T) {
	s = StrSliceS(s)
	n := copy(p, s)
	strTermN(p, n)
}

// StrNCopyStrS copies string s into a buffer p, including zero terminator.
//
// If size of the buffer is too small, the string will be trimmed. It may omit zero terminator.
//
// For difference between StrCopyStrS and StrNCopyStrS see docs for StrCopyS or StrNCopyS.
func StrNCopyStrS(p []byte, s string) {
	s = StrSliceStr(s)
	n := copy(p, s)
	strTermN(p, n)
}

// StrNCopyStr16 copies string s into a UTF-16 buffer pointed by p, including zero terminator.
//
// Size of buffer p must be large enough for the string. It may omit zero terminator.
//
// Deprecated: Unsafe, use StrNCopyStr16S.
func StrNCopyStr16(p *uint16, s string, sz int) {
	dst := unsafe.Slice(p, sz)
	StrNCopyStr16S(dst, s)
}

// StrNCopyStr16S copies string s into a UTF-16 buffer p, including zero terminator.
//
// If size of the buffer is too small, the string will be trimmed. It may omit zero terminator.
//
// For difference between StrCopyStr16S and StrNCopyStr16S see docs for StrCopyS or StrNCopyS.
func StrNCopyStr16S(p []uint16, s string) {
	s = StrSliceStr(s)
	s16 := utf16.Encode([]rune(s))
	n := copy(p, s16)
	strTermN(p, n)
}
