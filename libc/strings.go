package libc

import (
	"unicode/utf16"
)

// GoString is similar to C.GoString, but doesn't use C package.
//
// Deprecated: Unsafe, use GoStringS.
func GoString(p *byte) string {
	return string(StrSlice(p))
}

// GoStringS is similar to C.GoString, but doesn't use C package. It accepts slice.
func GoStringS(p []byte) string {
	return string(StrSliceS(p))
}

// GoString16 is similar to GoString, but operates on UTF-16 string.
//
// Deprecated: Unsafe, use GoStringS16.
func GoString16(p *uint16) string {
	return string(utf16.Decode(StrSlice(p)))
}

// GoStringS16 is similar to GoString, but operates on UTF-16 string. It accepts slice.
func GoStringS16(p []uint16) string {
	return string(utf16.Decode(StrSliceS(p)))
}

// CString is similar to C.CString, but uses the allocator provided by this package.
//
// Caller must free the string with Free.
func CString(s string) *byte {
	s = StrSliceStr(s)
	p := Make[byte](len(s) + 1)
	copy(p, s)
	p[len(s)] = 0
	return &p[0]
}

// CStringS is similar to C.CString, but uses the allocator provided by this package. It returns slice.
//
// Caller must free the string with Free.
func CStringS(s string) []byte {
	s = StrSliceStr(s)
	p := Make[byte](len(s) + 1)
	copy(p, s)
	p[len(s)] = 0
	return p[:len(s)]
}

// CString16 is similar to CString, but produces a UTF-16 string.
func CString16(s string) *uint16 {
	s = StrSliceStr(s)
	s16 := utf16.Encode([]rune(s))
	p := Make[uint16](len(s16) + 1)
	copy(p, s16)
	p[len(s16)] = 0
	return &p[0]
}

// CStringS16 is similar to CString, but produces a UTF-16 string. It returns slice.
func CStringS16(s string) []uint16 {
	s = StrSliceStr(s)
	s16 := utf16.Encode([]rune(s))
	p := Make[uint16](len(s16) + 1)
	copy(p, s16)
	p[len(s16)] = 0
	return p[:len(s16)]
}
