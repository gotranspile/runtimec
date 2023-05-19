package libc_test

import (
	"testing"

	"github.com/gotranspile/runtimec/libc"
	"github.com/gotranspile/runtimec/libc/testalloc"
)

func TestGoAllocator(t *testing.T) {
	testalloc.RunTestAlloc(t, new(libc.GoAllocator))
}
