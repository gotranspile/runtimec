//go:build cgo

package calloc

/*
#include <stdlib.h>
*/
import "C"
import (
	"unsafe"

	"github.com/gotranspile/runtimec/internal/syncu"
)

// CAllocator uses cgo to allocate memory. Zero value is safe to use.
type CAllocator struct {
	// callocSize stores size information for all C allocations. This is done to support Realloc natively.
	sizes syncu.Map[unsafe.Pointer, uintptr]
}

func (a *CAllocator) calloc(cnt, sz uintptr) unsafe.Pointer {
	p := C.calloc(C.size_t(cnt), C.size_t(sz))
	a.sizes.Store(unsafe.Pointer(p), cnt*sz)
	return p
}

func (a *CAllocator) Realloc(p unsafe.Pointer, cnt, sz uintptr) unsafe.Pointer {
	if sz*cnt == 0 {
		panic("zero alloc")
	}
	if p == nil {
		// we want zero-initialized memory
		return a.calloc(cnt, sz)
	}
	osz, ok := a.sizes.LoadAndDelete(p)
	if !ok {
		panic("realloc of memory not owned by the allocator")
	}
	defer C.free(p)
	// sub-optimal, but zero initialized
	p2 := a.calloc(cnt, sz)
	src := unsafe.Slice((*byte)(p), osz)
	dst := unsafe.Slice((*byte)(p2), cnt*sz)
	copy(dst, src)
	return p2
}

func (a *CAllocator) Free(p unsafe.Pointer) {
	if p == nil {
		panic("nil pointer in free")
	}
	_, ok := a.sizes.LoadAndDelete(p)
	if !ok {
		panic("free of memory not owned by the allocator")
	}
	C.free(p)
}
