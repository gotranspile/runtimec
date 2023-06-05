package libc

import "testing"

func TestStrNCopy(t *testing.T) {
	cases := []struct {
		name  string
		old   bool
		limit int
		src   []byte
		exp   [5]byte
	}{
		{name: "must pad with zeros (new)", old: false, src: nil, exp: [5]byte{0, 0, 0, 0, 0}},
		{name: "must pad with zeros (old)", old: true, src: nil, exp: [5]byte{0, 0, 0, 0, 0}},
		{name: "enough space (new)", old: false, src: []byte{2, 2}, exp: [5]byte{2, 2, 0, 0, 0}},
		{name: "enough space (old)", old: true, src: []byte{2, 2}, exp: [5]byte{2, 2, 0, 0, 0}},
		{name: "enough space (respect limit)", old: true, limit: 2, src: []byte{2, 2}, exp: [5]byte{2, 2, 0, 1, 1}},
		{name: "no space for zero (new)", old: false, src: []byte{2, 2, 2, 2, 2}, exp: [5]byte{2, 2, 2, 2, 2}},
		{name: "no space for zero (old)", old: true, src: []byte{2, 2, 2, 2, 2}, exp: [5]byte{2, 2, 2, 2, 2}},
		{name: "no space for zero (new, respect limit)", old: false, limit: 1, src: []byte{2, 2, 2, 2}, exp: [5]byte{2, 2, 2, 2, 0}},
		{name: "no space for zero (old, respect limit)", old: true, limit: 1, src: []byte{2, 2, 2, 2}, exp: [5]byte{2, 2, 2, 2, 1}},
	}
	t.Run("ptr", func(t *testing.T) {
		for _, c := range cases {
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
				StrNCopy(&dst[0], &src[0], len(dst)-c.limit)
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
				StrNCopySlice(&dst[0], c.src, len(dst)-c.limit)
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
				StrNCopyS(dst[:len(dst)-c.limit], c.src)
				if dst != c.exp {
					t.Fatal(dst, c.exp)
				}
			})
		}
	})
}

func TestStrNCopyStr(t *testing.T) {
	cases := []struct {
		name  string
		old   bool
		limit int
		src   string
		exp   [5]byte
	}{
		{name: "must pad with zeros (new)", old: false, src: "", exp: [5]byte{0, 0, 0, 0, 0}},
		{name: "must pad with zeros (old)", old: true, src: "", exp: [5]byte{0, 0, 0, 0, 0}},
		{name: "enough space (new)", old: false, src: "22", exp: [5]byte{'2', '2', 0, 0, 0}},
		{name: "enough space (old)", old: true, src: "22", exp: [5]byte{'2', '2', 0, 0, 0}},
		{name: "enough space (respect limit)", old: true, limit: 2, src: "22", exp: [5]byte{'2', '2', 0, 1, 1}},
		{name: "no space for zero (new)", old: false, src: "22222", exp: [5]byte{'2', '2', '2', '2', '2'}},
		{name: "no space for zero (old)", old: true, src: "22222", exp: [5]byte{'2', '2', '2', '2', '2'}},
		{name: "no space for zero (new, respect limit)", old: false, limit: 1, src: "2222", exp: [5]byte{'2', '2', '2', '2', 0}},
		{name: "no space for zero (old, respect limit)", old: true, limit: 1, src: "2222", exp: [5]byte{'2', '2', '2', '2', 1}},
	}
	t.Run("ptr", func(t *testing.T) {
		for _, c := range cases {
			c := c
			t.Run(c.name, func(t *testing.T) {
				var dst [5]byte
				if c.old {
					dst = [5]byte{1, 1, 1, 1, 1}
				}
				StrNCopyStr(&dst[0], c.src, len(dst)-c.limit)
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
				StrNCopyStrS(dst[:len(dst)-c.limit], c.src)
				if dst != c.exp {
					t.Fatal(dst, c.exp)
				}
			})
		}
	})
	t.Run("ptr16", func(t *testing.T) {
		for _, c := range cases {
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
				StrNCopyStr16(&dst[0], c.src, len(dst)-c.limit)
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
				var dst, exp [5]uint16
				if c.old {
					dst = [5]uint16{1, 1, 1, 1, 1}
				}
				for i, v := range c.exp {
					exp[i] = uint16(v)
				}
				StrNCopyStr16S(dst[:len(dst)-c.limit], c.src)
				if dst != exp {
					t.Fatal(dst, exp)
				}
			})
		}
	})
}
