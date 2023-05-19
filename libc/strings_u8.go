package libc

import "unsafe"

// GoString is similar to C.GoString, but doesn't use C package.
func GoString(p *byte) string {
	n := StrLen(p)
	if n == 0 {
		return ""
	}
	return string(unsafe.Slice(p, n))
}

// GoStringS is similar to C.GoString, but doesn't use C package. It accepts slice.
func GoStringS(p []byte) string {
	n := StrLenS(p)
	if n == 0 {
		return ""
	}
	return string(p[:n])
}

// CString is similar to C.CString, but uses the allocator provided by this package.
//
// Caller must free the string with Free.
func CString(s string) *byte {
	p := Make[byte](len(s) + 1)
	copy(p, s)
	p[len(s)] = 0
	return &p[0]
}

// CStringS is similar to C.CString, but uses the allocator provided by this package. It returns slice.
//
// Caller must free the string with Free.
func CStringS(s string) []byte {
	p := Make[byte](len(s) + 1)
	copy(p, s)
	p[len(s)] = 0
	return p[:len(s)]
}

// StrCopyZero copies string to p. It always adds a trailing zero byte, potentially trimming the string.
func StrCopyZero(p []byte, s string) {
	n := copy(p, s)
	if n < len(p) {
		p[n] = 0
	} else {
		p[n-1] = 0
	}
}

// StrCopyFull copies string to p. It adds a trailing zero byte if string is smaller than p.
func StrCopyFull(p []byte, s string) {
	n := copy(p, s)
	if n < len(p) {
		p[n] = 0
	}
}
