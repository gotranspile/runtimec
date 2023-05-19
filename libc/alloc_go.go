package libc

import (
	"unsafe"

	"github.com/gotranspile/runtimec/internal/syncu"
)

// GoAllocator is a simple allocator that uses mechanism similar to make to allocate memory.
// It keeps pointers of all allocations, so that GC will never reclaim it until Free.
// Zero value is safe for use.
type GoAllocator struct {
	allocs syncu.Map[*byte, []byte]
}

func (a *GoAllocator) calloc(cnt, sz uintptr) []byte {
	b := make([]byte, sz*cnt)
	a.allocs.Store(&b[0], b)
	return b
}

func (a *GoAllocator) Realloc(p unsafe.Pointer, cnt, sz uintptr) unsafe.Pointer {
	if sz*cnt == 0 {
		panic("zero alloc")
	}
	if p == nil {
		b := a.calloc(cnt, sz)
		return unsafe.Pointer(&b[0])
	}
	old, ok := a.allocs.LoadAndDelete((*byte)(p))
	if !ok {
		panic("realloc of memory not owned by the allocator")
	}
	b := make([]byte, sz*cnt)
	a.allocs.Store(&b[0], b)
	copy(b, old)
	return unsafe.Pointer(&b[0])
}

func (a *GoAllocator) Free(p unsafe.Pointer) {
	if p == nil {
		panic("nil pointer in free")
	}
	_, ok := a.allocs.LoadAndDelete((*byte)(p))
	if !ok {
		panic("free of memory not owned by the allocator")
	}
}
