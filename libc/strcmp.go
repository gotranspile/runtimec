package libc

import (
	"strings"
	"unsafe"

	"golang.org/x/exp/constraints"
)

// Reference implementation:
// https://github.com/bminor/musl/blob/master/src/string/strcmp.c
// https://github.com/bminor/musl/blob/master/src/string/strncmp.c

// StrCmp is a generic version of C strcmp.
//
// Deprecated: Unsafe, use StrCmpS.
func StrCmp[T constraints.Ordered](p1, p2 *T) int {
	if p1 == nil && p2 == nil {
		return 0
	} else if p1 == nil {
		return -1
	} else if p2 == nil {
		return +1
	}
	s1 := StrSlice(p1)
	s2 := StrSlice(p2)
	return StrCmpS(s1, s2)
}

// StrCmpS is the same as StrCmp, but accepts slices. It still looks for a zero value as a terminator.
func StrCmpS[T constraints.Ordered](p1, p2 []T) int {
	if cap(p1) == 0 && cap(p2) == 0 {
		return 0
	} else if cap(p1) == 0 {
		return -1
	} else if cap(p2) == 0 {
		return +1
	}
	var zero T
	for {
		if len(p1) == 0 && len(p2) == 0 {
			return 0
		} else if len(p1) == 0 {
			return -1
		} else if len(p2) == 0 {
			return +1
		}
		v1, v2 := p1[0], p2[0]
		if v1 == zero && v2 == zero {
			return 0
		} else if v1 == zero {
			return -1
		} else if v2 == zero {
			return +1
		}
		if v1 < v2 {
			return -1
		} else if v1 > v2 {
			return +1
		}
		p1 = p1[1:]
		p2 = p2[1:]
	}
}

// StrCmpN is the same as StrCmp, but accepts length for underlying arrays. It still looks for a zero value as a terminator.
func StrCmpN[T constraints.Ordered](p1, p2 *T, max int) int {
	if p1 == nil && p2 == nil {
		return 0
	} else if p1 == nil {
		return -1
	} else if p2 == nil {
		return +1
	}
	return StrCmpS(unsafe.Slice(p1, max), unsafe.Slice(p2, max))
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
