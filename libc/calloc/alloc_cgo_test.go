//go:build cgo

package calloc

import (
	"testing"

	"github.com/gotranspile/runtimec/libc/testalloc"
)

func TestCAllocator(t *testing.T) {
	testalloc.RunTestAlloc(t, new(CAllocator))
}
