package libc

// Reference implementation:
// https://github.com/bminor/musl/blob/master/src/string/strdup.c

// StrDup makes a copy of the array or string.
//
// Deprecated: Unsafe, use StrDupS.
func StrDup[T comparable](s *T) *T {
	src := StrSlice(s)
	dst := Make[T](len(src) + 1)
	copy(dst, src)
	var zero T
	dst[len(src)] = zero
	return &dst[0]
}

// StrDupS makes a copy of the array or string. It accepts and returns slices.
func StrDupS[T comparable](s []T) []T {
	src := StrSliceS(s)
	dst := Make[T](len(src) + 1)
	copy(dst, src)
	var zero T
	dst[len(src)] = zero
	return dst[:len(src)]
}
