package libc

import "unsafe"

// MemCopy is a generic implementation of C memcpy.
func MemCopy[T any](dst, src *T, sz int) {
	copy(unsafe.Slice(dst, sz), unsafe.Slice(src, sz))
}

// MemCopyP is similar to MemCopy, but accepts unsafe.Pointer.
func MemCopyP(dst, src unsafe.Pointer, sz int) {
	MemCopy((*byte)(dst), (*byte)(src), sz)
}
