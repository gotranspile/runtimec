package libc

import (
	"strings"
	"unicode/utf16"
	"unsafe"
)

// GoString16 is similar to GoString, but operates on UTF-16 string.
func GoString16(p *uint16) string {
	n := StrLen(p)
	if n == 0 {
		return ""
	}
	return string(utf16.Decode(unsafe.Slice(p, n)))
}

// GoStringS16 is similar to GoString, but operates on UTF-16 string. It accepts slice.
func GoStringS16(p []uint16) string {
	n := StrLenS(p)
	if n == 0 {
		return ""
	}
	return string(utf16.Decode(p[:n]))
}

// CString16 is similar to CString, but produces a UTF-16 string.
func CString16(s string) *uint16 {
	s16 := utf16.Encode([]rune(s))
	p := Make[uint16](len(s16) + 1)
	copy(p, s16)
	p[len(s16)] = 0
	return &p[0]
}

// CStringS16 is similar to CString, but produces a UTF-16 string. It returns slice.
func CStringS16(s string) []uint16 {
	s16 := utf16.Encode([]rune(s))
	p := Make[uint16](len(s16) + 1)
	copy(p, s16)
	p[len(s16)] = 0
	return p[:len(s16)]
}

// StrCopyZero16 copies string to UTF-16 buffer p. It always adds a trailing zero byte, potentially trimming the string.
func StrCopyZero16(p []uint16, s string) {
	s16 := utf16.Encode([]rune(s))
	n := copy(p, s16)
	if n < len(p) {
		p[n] = 0
	} else {
		p[n-1] = 0
	}
}

// StrCopyFull16 copies string to UTF-16 buffer p. It adds a trailing zero byte if string is smaller than p.
func StrCopyFull16(p []uint16, s string) {
	s16 := utf16.Encode([]rune(s))
	n := copy(p, s16)
	if n < len(p) {
		p[n] = 0
	}
}

// StrCopy16 copies UTF-16 string into p.
//
// Deprecated: Unsafe, use StrCopyZero16 or StrCopyFull16.
func StrCopy16(p *uint16, s string) {
	s16 := utf16.Encode([]rune(s))
	dst := StrSliceN(p, len(s16)+1)
	n := copy(dst, s16)
	dst[n] = 0
}

// StrCaseCmp16 is similar to C strcasecmp for UTF-16.
//
// Deprecated: Unsafe, use StrCaseCmpS16 or StrCaseCmpN16.
func StrCaseCmp16(p1, p2 *uint16) int {
	if p1 == nil && p2 == nil {
		return 0
	} else if p1 == nil {
		return -1
	} else if p2 == nil {
		return +1
	}
	s1 := StrSlice(p1)
	s2 := StrSlice(p2)
	return StrCaseCmpS16(s1, s2)
}

// StrCaseCmpS16 is safe version of StrCaseCmp16.
func StrCaseCmpS16(p1, p2 []uint16) int {
	if cap(p1) == 0 && cap(p2) == 0 {
		return 0
	} else if cap(p1) == 0 {
		return -1
	} else if cap(p2) == 0 {
		return +1
	}
	s1 := GoStringS16(p1)
	s2 := GoStringS16(p2)
	s1 = strings.ToLower(s1)
	s2 = strings.ToLower(s2)
	return strings.Compare(s1, s2)
}

// StrCaseCmpN16 is safe version of StrCaseCmp16 which requires length.
func StrCaseCmpN16(p1, p2 *uint16, max int) int {
	if p1 == nil && p2 == nil {
		return 0
	} else if p1 == nil {
		return -1
	} else if p2 == nil {
		return +1
	}
	s1 := unsafe.Slice(p1, max)
	s2 := unsafe.Slice(p2, max)
	return StrCaseCmpS16(s1, s2)
}
