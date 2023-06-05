package libc

import (
	"unsafe"
)

var alloc Allocator = new(GoAllocator)

// SetAllocator sets the global allocator. By default, GoAllocator is used. It returns the old allocator.
func SetAllocator(a Allocator) Allocator {
	old := alloc
	alloc = a
	return old
}

// Allocator is an interface that provides C-like allocation mechanisms.
type Allocator interface {
	// Realloc allocates a zero-initialized array of cnt element, each of size sz.
	//
	// If the old pointer is given, all data is copied to the new array (possibly trimming the content),
	// and the old array is automatically freed. Zero-initialization is best-effort in this case.
	Realloc(ptr unsafe.Pointer, cnt, sz uintptr) unsafe.Pointer
	// Free previously allocated memory.
	Free(ptr unsafe.Pointer)
}

// New is similar to builtin new, but uses the allocator. Caller must free the result with Free.
func New[T any]() *T {
	var zero T
	p := alloc.Realloc(nil, 1, unsafe.Sizeof(zero))
	return (*T)(p)
}

// Make is similar to builtin make, but uses the allocator. Caller must free the result with FreeS.
func Make[T any](sz int) []T {
	var zero T
	p := alloc.Realloc(nil, uintptr(sz), unsafe.Sizeof(zero))
	return unsafe.Slice((*T)(p), sz)
}

// Clone the object, using memory provided by the allocator. Caller must free the result with Free.
func Clone[T any](p *T) *T {
	p2 := New[T]()
	*p2 = *p
	return p2
}

// CloneS the slice, using memory provided by the allocator. Caller must free the result with FreeS.
func CloneS[T any](p []T) []T {
	p2 := Make[T](len(p))
	copy(p2, p)
	return p2
}

// Remake is similar to builtin make and C realloc. Caller must free the result with FreeS.
func Remake[T any](p []T, sz int) []T {
	if cap(p) == 0 {
		return Make[T](sz)
	}
	p = p[:1]
	p2 := Realloc(&p[0], sz)
	return unsafe.Slice(p2, sz)
}

// Malloc allocates memory using the allocator. It is similar malloc in C, but always zero-initializes memory.
func Malloc(sz uintptr) unsafe.Pointer {
	return alloc.Realloc(nil, 1, sz)
}

// Calloc allocates memory using the allocator. It is similar calloc in C.
func Calloc(cnt, sz uintptr) unsafe.Pointer {
	return alloc.Realloc(nil, cnt, sz)
}

// Realloc is a generic implementation of C realloc.
func Realloc[T any](p *T, sz int) *T {
	var zero T
	p2 := alloc.Realloc(unsafe.Pointer(p), uintptr(sz), unsafe.Sizeof(zero))
	return (*T)(p2)
}

// ReallocP is similar to C realloc.
func ReallocP(p unsafe.Pointer, cnt, sz uintptr) unsafe.Pointer {
	return alloc.Realloc(p, cnt, sz)
}

// Free allocated memory.
func Free[T any](p *T) {
	FreeP(unsafe.Pointer(p))
}

// FreeP is similar to Free, but accepts unsafe.Pointer.
func FreeP(p unsafe.Pointer) {
	alloc.Free(p)
}

// FreeS is similar to Free, but accepts a slice.
func FreeS[T any](p []T) {
	if cap(p) == 0 {
		panic("nil pointer in free")
	}
	p = p[:1]
	FreeP(unsafe.Pointer(&p[0]))
}
