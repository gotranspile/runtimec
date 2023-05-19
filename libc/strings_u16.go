package libc

import (
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
