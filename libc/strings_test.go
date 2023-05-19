package libc

import "testing"

func TestStrLen(t *testing.T) {
	if StrLen((*byte)(nil)) != 0 {
		t.Fatal()
	}
	var buf [5]byte
	if StrLen(&buf[0]) != 0 {
		t.Fatal()
	}
	buf[0] = 'a'
	buf[1] = 'b'
	buf[3] = 'c'
	if StrLen(&buf[0]) != 2 {
		t.Fatal()
	}
	if StrLen(&buf[3]) != 1 {
		t.Fatal()
	}
}

func TestStrLenN(t *testing.T) {
	if StrLenN((*byte)(nil), 0) != 0 {
		t.Fatal()
	}
	if StrLenN((*byte)(nil), 5) != 0 {
		t.Fatal()
	}
	var buf [5]byte
	if StrLenN(&buf[0], 5) != 0 {
		t.Fatal()
	}
	buf[0] = 'a'
	buf[1] = 'b'
	buf[3] = 'c'
	if StrLenN(&buf[0], 5) != 2 {
		t.Fatal()
	}
	if StrLenN(&buf[0], 1) != 1 {
		t.Fatal()
	}
	if StrLenN(&buf[0], 2) != 2 {
		t.Fatal()
	}
	if StrLenN(&buf[3], 2) != 1 {
		t.Fatal()
	}
}

func TestStrLenS(t *testing.T) {
	if StrLenS(([]byte)(nil)) != 0 {
		t.Fatal()
	}
	if StrLenS([]byte{}) != 0 {
		t.Fatal()
	}
	var buf [5]byte
	if StrLenS(buf[:]) != 0 {
		t.Fatal()
	}
	buf[0] = 'a'
	buf[1] = 'b'
	buf[3] = 'c'
	if StrLenS(buf[:]) != 2 {
		t.Fatal()
	}
	if StrLenS(buf[:1]) != 1 {
		t.Fatal()
	}
	if StrLenS(buf[:2]) != 2 {
		t.Fatal()
	}
	if StrLenS(buf[3:]) != 1 {
		t.Fatal()
	}
}

func TestStrSlice(t *testing.T) {
	s := StrSlice((*byte)(nil))
	if s != nil {
		t.Fatal()
	}
	var buf [5]byte
	s = StrSlice(&buf[0])
	if s == nil || len(s) != 0 || cap(s) != 1 {
		t.Fatal()
	}
	buf[0] = 'a'
	buf[1] = 'b'
	s = StrSlice(&buf[0])
	if s == nil || len(s) != 2 || cap(s) != 3 {
		t.Fatal(len(s), cap(s))
	}
}

func TestStrSliceN(t *testing.T) {
	s := StrSliceN((*byte)(nil), 0)
	if s != nil {
		t.Fatal()
	}
	s = StrSliceN((*byte)(nil), 5)
	if s != nil {
		t.Fatal()
	}
	var buf [5]byte
	s = StrSliceN(&buf[0], 5)
	if s == nil || len(s) != 0 || cap(s) != 5 {
		t.Fatal()
	}
	buf[0] = 'a'
	buf[1] = 'b'
	s = StrSliceN(&buf[0], 5)
	if s == nil || len(s) != 2 || cap(s) != 5 {
		t.Fatal(len(s), cap(s))
	}
}

func TestStrCopyZeroS(t *testing.T) {
	var dst [5]byte

	dst = [5]byte{1, 1, 1, 1, 1}
	StrCopyZeroS(dst[:], nil)
	if dst != [5]byte{0, 1, 1, 1, 1} {
		t.Fatal()
	}

	dst = [5]byte{1, 1, 1, 1, 1}
	StrCopyZeroS(dst[:], []byte{})
	if dst != [5]byte{0, 1, 1, 1, 1} {
		t.Fatal()
	}

	dst = [5]byte{1, 1, 1, 1, 1}
	StrCopyZeroS(dst[:], []byte{2, 2})
	if dst != [5]byte{2, 2, 0, 1, 1} {
		t.Fatal()
	}

	dst = [5]byte{1, 1, 1, 1, 1}
	StrCopyZeroS(dst[:], []byte{2, 2, 2, 2, 2})
	if dst != [5]byte{2, 2, 2, 2, 0} {
		t.Fatal()
	}
}

func TestStrCopyFullS(t *testing.T) {
	var dst [5]byte

	dst = [5]byte{1, 1, 1, 1, 1}
	StrCopyFullS(dst[:], nil)
	if dst != [5]byte{0, 1, 1, 1, 1} {
		t.Fatal()
	}

	dst = [5]byte{1, 1, 1, 1, 1}
	StrCopyFullS(dst[:], []byte{})
	if dst != [5]byte{0, 1, 1, 1, 1} {
		t.Fatal()
	}

	dst = [5]byte{1, 1, 1, 1, 1}
	StrCopyFullS(dst[:], []byte{2, 2})
	if dst != [5]byte{2, 2, 0, 1, 1} {
		t.Fatal()
	}

	dst = [5]byte{1, 1, 1, 1, 1}
	StrCopyFullS(dst[:], []byte{2, 2, 2, 2, 2})
	if dst != [5]byte{2, 2, 2, 2, 2} {
		t.Fatal()
	}
}

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
