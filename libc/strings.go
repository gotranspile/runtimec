package libc

import (
	"unsafe"

	"golang.org/x/exp/constraints"
)

// StrLen is a generic version of C strlen.
//
// It accepts a pointer to the first array element and advances it until it reaches a zero value.
// Position of this zero terminator is returned, which equals to a length of the array.
//
// Additionally, function checks for a nil pointer and returns 0 length in this case.
//
// The behavior is undefined if p is not a pointer to a null-terminated array.
func StrLen[T comparable](p *T) int {
	if p == nil {
		return 0
	}
	var zero T
	elem := unsafe.Sizeof(zero)
	sz := 0
	for *p != zero {
		p = (*T)(unsafe.Add(unsafe.Pointer(p), elem))
		sz++
	}
	return sz
}

// StrLenN is a generic version of C strlen. It's the same as StrLen, but accepts a length of the buffer.
//
// It is safer than using StrLenS and unsafe.Slice directly, because it checks for a zero pointer before converting to a slice.
func StrLenN[T comparable](p *T, sz uintptr) int {
	if p == nil || sz == 0 {
		return 0
	}
	return StrLenS(unsafe.Slice(p, sz))
}

// StrLenS is similar to StrLen, but accepts a slice. It still looks for a zero value as a terminator in the slice.
func StrLenS[T comparable](p []T) int {
	if len(p) == 0 {
		return 0
	}
	var zero T
	for i, v := range p {
		if v == zero {
			return i
		}
	}
	return len(p)
}

// StrSlice uses StrLen to determine the length of a zero-terminated string, and returns it as a slice.
//
// This method is only suitable for reading from a string. For writing to a string buffer, see StrSliceN.
func StrSlice[T comparable](p *T) []T {
	if p == nil {
		return nil
	}
	n := StrLen(p)
	return unsafe.Slice(p, n+1)[:n]
}

// StrSliceN uses StrLen to determine the length of a zero-terminated string, and returns it as a slice.
//
// It uses max as a capacity of the returned slice, making it suitable for writing to a string buffer.
func StrSliceN[T comparable](p *T, max int) []T {
	if p == nil || max == 0 {
		return nil
	}
	s := unsafe.Slice(p, max)
	n := StrLenS(s)
	return s[:n]
}

// StrCopyZeroS copies slice into p, looking for a zero value terminator in s.
//
// It always adds a trailing zero value to p, potentially trimming the string.
func StrCopyZeroS[T comparable](p []T, s []T) {
	n := StrLenS(s)
	n = copy(p, s[:n])
	var zero T
	if n < len(p) {
		p[n] = zero
	} else {
		p[n-1] = zero
	}
}

// StrCopyFullS copies slice into p, looking for a zero value terminator in s.
//
// It omits trailing zero value if string length is equal to p.
func StrCopyFullS[T comparable](p []T, s []T) {
	n := StrLenS(s)
	n = copy(p, s[:n])
	var zero T
	if n < len(p) {
		p[n] = zero
	}
}

// StrCmp is a generic version of C strcmp.
func StrCmp[T constraints.Ordered](p1, p2 *T) int {
	if p1 == nil && p2 == nil {
		return 0
	} else if p1 == nil {
		return -1
	} else if p2 == nil {
		return +1
	}
	var zero T
	elem := unsafe.Sizeof(zero)
	for {
		v1, v2 := *p1, *p2
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
		p1 = (*T)(unsafe.Add(unsafe.Pointer(p1), elem))
		p2 = (*T)(unsafe.Add(unsafe.Pointer(p2), elem))
	}
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
