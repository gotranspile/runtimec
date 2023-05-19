package libc

import (
	"strings"
	"unsafe"
)

// GoString is similar to C.GoString, but doesn't use C package.
//
// Deprecated: Unsafe, use GoStringS.
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

// StrCopy copies string into p.
//
// Deprecated: Unsafe, use StrCopyZero or StrCopyFull.
func StrCopy(p *byte, src string) {
	dst := StrSliceN(p, len(src)+1)
	n := copy(dst, src)
	dst[n] = 0
}

// StrCaseCmp is similar to C strcasecmp.
//
// Deprecated: Unsafe, use StrCaseCmpS or StrCaseCmpN.
func StrCaseCmp(p1, p2 *byte) int {
	if p1 == nil && p2 == nil {
		return 0
	} else if p1 == nil {
		return -1
	} else if p2 == nil {
		return +1
	}
	s1 := StrSlice(p1)
	s2 := StrSlice(p2)
	return StrCaseCmpS(s1, s2)
}

// StrCaseCmpS is safe version of StrCaseCmp.
func StrCaseCmpS(p1, p2 []byte) int {
	if cap(p1) == 0 && cap(p2) == 0 {
		return 0
	} else if cap(p1) == 0 {
		return -1
	} else if cap(p2) == 0 {
		return +1
	}
	s1 := GoStringS(p1)
	s2 := GoStringS(p2)
	s1 = strings.ToLower(s1)
	s2 = strings.ToLower(s2)
	return strings.Compare(s1, s2)
}

// StrCaseCmpN is safe version of StrCaseCmp which requires length.
func StrCaseCmpN(p1, p2 *byte, max int) int {
	if p1 == nil && p2 == nil {
		return 0
	} else if p1 == nil {
		return -1
	} else if p2 == nil {
		return +1
	}
	s1 := unsafe.Slice(p1, max)
	s2 := unsafe.Slice(p2, max)
	return StrCaseCmpS(s1, s2)
}
