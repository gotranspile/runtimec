package libc

import "testing"

func TestStrCopy(t *testing.T) {
	cases := []struct {
		name   string
		old    bool
		limits bool
		src    []byte
		exp    [5]byte
		exps   *[5]byte
	}{
		{name: "must zero term (new)", old: false, src: nil, exp: [5]byte{0, 0, 0, 0, 0}},
		{name: "must zero term (old)", old: true, src: nil, exp: [5]byte{0, 1, 1, 1, 1}, exps: &[5]byte{0, 0, 0, 0, 0}},
		{name: "enough space (new)", old: false, src: []byte{2, 2}, exp: [5]byte{2, 2, 0, 0, 0}},
		{name: "enough space (old)", old: true, src: []byte{2, 2}, exp: [5]byte{2, 2, 0, 1, 1}, exps: &[5]byte{2, 2, 0, 0, 0}},
		{name: "no space (new)", old: false, limits: true, src: []byte{1, 2, 3, 4, 5, 6}, exp: [5]byte{1, 2, 3, 4, 0}},
		{name: "no space (old)", old: true, limits: true, src: []byte{1, 2, 3, 4, 5, 6}, exp: [5]byte{1, 2, 3, 4, 0}},
	}
	t.Run("ptr", func(t *testing.T) {
		for _, c := range cases {
			if c.limits {
				continue
			}
			c := c
			t.Run(c.name, func(t *testing.T) {
				var (
					src [6]byte
					dst [5]byte
				)
				if c.old {
					dst = [5]byte{1, 1, 1, 1, 1}
				}
				copy(src[:], c.src)
				StrCopy(&dst[0], &src[0])
				if dst != c.exp {
					t.Fatal(dst, c.exp)
				}
			})
		}
	})
	t.Run("slice", func(t *testing.T) {
		for _, c := range cases {
			if c.limits {
				continue
			}
			c := c
			t.Run(c.name, func(t *testing.T) {
				var dst [5]byte
				if c.old {
					dst = [5]byte{1, 1, 1, 1, 1}
				}
				StrCopySlice(&dst[0], c.src)
				if dst != c.exp {
					t.Fatal(dst, c.exp)
				}
			})
		}
	})
	t.Run("slice2", func(t *testing.T) {
		for _, c := range cases {
			c := c
			t.Run(c.name, func(t *testing.T) {
				var dst [5]byte
				if c.old {
					dst = [5]byte{1, 1, 1, 1, 1}
				}
				exp := c.exp
				if c.exps != nil {
					exp = *c.exps
				}
				StrCopyS(dst[:], c.src)
				if dst != exp {
					t.Fatal(dst, exp)
				}
			})
		}
	})
}

func TestStrCopyStr(t *testing.T) {
	cases := []struct {
		name   string
		old    bool
		limits bool
		src    string
		exp    [5]byte
		exps   *[5]byte
	}{
		{name: "must zero term (new)", old: false, src: "", exp: [5]byte{0, 0, 0, 0, 0}},
		{name: "must zero term (old)", old: true, src: "", exp: [5]byte{0, 1, 1, 1, 1}, exps: &[5]byte{0, 0, 0, 0, 0}},
		{name: "enough space (new)", old: false, src: "22", exp: [5]byte{'2', '2', 0, 0, 0}},
		{name: "enough space (old)", old: true, src: "22", exp: [5]byte{'2', '2', 0, 1, 1}, exps: &[5]byte{'2', '2', 0, 0, 0}},
		{name: "no space (new)", old: false, limits: true, src: "123456", exp: [5]byte{'1', '2', '3', '4', 0}},
		{name: "no space (old)", old: true, limits: true, src: "123456", exp: [5]byte{'1', '2', '3', '4', 0}},
	}
	t.Run("ptr", func(t *testing.T) {
		for _, c := range cases {
			if c.limits {
				continue
			}
			c := c
			t.Run(c.name, func(t *testing.T) {
				var dst [5]byte
				if c.old {
					dst = [5]byte{1, 1, 1, 1, 1}
				}
				StrCopyStr(&dst[0], c.src)
				if dst != c.exp {
					t.Fatal(dst, c.exp)
				}
			})
		}
	})
	t.Run("slice", func(t *testing.T) {
		for _, c := range cases {
			c := c
			t.Run(c.name, func(t *testing.T) {
				var dst [5]byte
				if c.old {
					dst = [5]byte{1, 1, 1, 1, 1}
				}
				exp := c.exp
				if c.exps != nil {
					exp = *c.exps
				}
				StrCopyStrS(dst[:], c.src)
				if dst != exp {
					t.Fatal(dst, exp)
				}
			})
		}
	})
	t.Run("ptr16", func(t *testing.T) {
		for _, c := range cases {
			if c.limits {
				continue
			}
			c := c
			t.Run(c.name, func(t *testing.T) {
				var (
					dst, exp [5]uint16
				)
				if c.old {
					dst = [5]uint16{1, 1, 1, 1, 1}
				}
				for i, v := range c.exp {
					exp[i] = uint16(v)
				}
				StrCopyStr16(&dst[0], c.src)
				if dst != exp {
					t.Fatal(dst, exp)
				}
			})
		}
	})
	t.Run("slice16", func(t *testing.T) {
		for _, c := range cases {
			c := c
			t.Run(c.name, func(t *testing.T) {
				var dst, exp16 [5]uint16
				if c.old {
					dst = [5]uint16{1, 1, 1, 1, 1}
				}
				exp := c.exp
				if c.exps != nil {
					exp = *c.exps
				}
				for i, v := range exp {
					exp16[i] = uint16(v)
				}
				StrCopyStr16S(dst[:], c.src)
				if dst != exp16 {
					t.Fatal(dst, exp16)
				}
			})
		}
	})
}
