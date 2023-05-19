package testalloc

import (
	"bytes"
	"testing"

	"github.com/gotranspile/runtimec/libc"
)

func RunTestAlloc(t *testing.T, a libc.Allocator) {
	old := libc.SetAllocator(a)
	t.Cleanup(func() {
		libc.SetAllocator(old)
	})
	t.Run("zero init", func(t *testing.T) {
		const sz = 10
		s := libc.Make[byte](sz)
		defer libc.FreeS(s)
		if !bytes.Equal(s, make([]byte, 10)) {
			t.Fatal()
		}
	})
	t.Run("clone", func(t *testing.T) {
		const sz = 10
		s := libc.Make[byte](sz)
		defer libc.FreeS(s)
		for i := range s {
			s[i] = byte(i)
		}
		s2 := libc.Clone(s)
		defer libc.FreeS(s2)
		if !bytes.Equal(s, s2) {
			t.Fatal()
		}
		if &s[0] == &s2[0] {
			t.Fatal()
		}
	})
	t.Run("remake zero init", func(t *testing.T) {
		const sz = 10
		s := libc.Remake[byte](nil, sz)
		defer libc.FreeS(s)
		if !bytes.Equal(s, make([]byte, 10)) {
			t.Fatal()
		}
	})
	t.Run("remake", func(t *testing.T) {
		const sz = 10
		s := libc.Make[byte](sz)
		exp := make([]byte, 2*sz)
		for i := range s {
			s[i] = byte(i)
			exp[i] = byte(i)
		}
		s2 := libc.Remake(s, 2*sz)
		defer libc.FreeS(s2)

		if !bytes.Equal(s2, exp) {
			t.Fatal(s2)
		}
		if &s[0] == &s2[0] {
			t.Fatal()
		}
	})
}
