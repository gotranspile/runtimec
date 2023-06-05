package libc

import "testing"

func TestStrCmp(t *testing.T) {
	cases := []struct {
		name string
		a, b []byte
		exp  int
	}{
		{name: "nils", a: nil, b: nil, exp: 0},
		{name: "nil and empty", a: nil, b: []byte{}, exp: 0},
		{name: "empty", a: []byte{}, b: []byte{}, exp: 0},
		{name: "nil and non-empty", a: nil, b: []byte{1}, exp: -1},
		{name: "same", a: []byte{1, 2, 3}, b: []byte{1, 2, 3}, exp: 0},
		{name: "less", a: []byte{1, 2, 1}, b: []byte{1, 2, 3}, exp: -1},
		{name: "less 2", a: []byte{1, 1, 3}, b: []byte{1, 2, 3}, exp: -1},
		{name: "greater", a: []byte{1, 2, 9}, b: []byte{1, 2, 3}, exp: +1},
		{name: "greater 2", a: []byte{1, 9, 3}, b: []byte{1, 2, 3}, exp: +1},
		{name: "shorter", a: []byte{1, 2}, b: []byte{1, 2, 3}, exp: -1},
		{name: "longer", a: []byte{1, 2, 3, 1}, b: []byte{1, 2, 3}, exp: +1},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Run("slices", func(t *testing.T) {
				got := StrCmpS(c.a, c.b)
				if got != c.exp {
					t.Fatal(got)
				}
				got = StrCmpS(c.b, c.a)
				if got != -c.exp {
					t.Fatal(got)
				}
			})
			t.Run("ptrs", func(t *testing.T) {
				var pa, pb *byte
				if len(c.a) != 0 {
					dst := make([]byte, len(c.a)+1)
					copy(dst, c.a)
					pa = &dst[0]
				}
				if len(c.b) != 0 {
					dst := make([]byte, len(c.b)+1)
					copy(dst, c.b)
					pb = &dst[0]
				}
				got := StrCmp(pa, pb)
				if got != c.exp {
					t.Fatal(got)
				}
				got = StrCmp(pb, pa)
				if got != -c.exp {
					t.Fatal(got)
				}
			})
			t.Run("ptrs max", func(t *testing.T) {
				var pa, pb *byte
				if len(c.a) != 0 {
					dst := make([]byte, 10)
					copy(dst, c.a)
					pa = &dst[0]
				}
				if len(c.b) != 0 {
					dst := make([]byte, 10)
					copy(dst, c.b)
					pb = &dst[0]
				}
				got := StrCmpN(pa, pb, 10)
				if got != c.exp {
					t.Fatal(got)
				}
				got = StrCmpN(pb, pa, 10)
				if got != -c.exp {
					t.Fatal(got)
				}
			})
		})
	}
}
