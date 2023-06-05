package libc

import "testing"

func TestBoolToInt(t *testing.T) {
	if BoolToInt(false) != 0 || BoolToInt(true) != 1 {
		t.Fatal()
	}
}

func TestIf(t *testing.T) {
	x := If(true, 1, 2)
	if x != 1 {
		t.FailNow()
	}
	x = If(false, 1, 2)
	if x != 2 {
		t.FailNow()
	}
	y := 0
	x = If(true, 1, PostInc(&y))
	if x != 1 {
		t.FailNow()
	}
	// Side effect is applied, even though the branch is not taken.
	if y != 1 {
		t.FailNow()
	}
}

func TestIfFunc(t *testing.T) {
	y := 0

	x := IfFunc(true, func() int {
		return 1
	}, func() int {
		return PostDec(&y)
	})
	if x != 1 {
		t.FailNow()
	}
	// Side effect not applied.
	if y != 0 {
		t.FailNow()
	}

	x = IfFunc(false, func() int {
		return PostDec(&y)
	}, func() int {
		return 1
	})
	if x != 1 {
		t.FailNow()
	}
	// Side effect not applied.
	if y != 0 {
		t.FailNow()
	}
}

func TestAssign(t *testing.T) {
	cases := []struct {
		name string
		fnc  func(p *int, v int) int
		a, b int
		exp  int
	}{
		{name: "=", fnc: Assign[int], a: -1, b: 1, exp: 1},
		{name: "+=", fnc: AddAssign[int], a: -1, b: 2, exp: 1},
		{name: "-=", fnc: SubAssign[int], a: -1, b: 1, exp: -2},
		{name: "*=", fnc: MulAssign[int], a: -1, b: 2, exp: -2},
		{name: "/=", fnc: DivAssign[int], a: -2, b: 2, exp: -1},
		{name: "%=", fnc: ModAssign[int], a: 1, b: 2, exp: 1},
		{name: "&=", fnc: AndAssign[int], a: 1, b: 3, exp: 1},
		{name: "|=", fnc: OrAssign[int], a: 1, b: 3, exp: 3},
		{name: "^=", fnc: XorAssign[int], a: 1, b: 3, exp: 2},
		{name: "<<=", fnc: LshAssign[int], a: 1, b: 3, exp: 8},
		{name: ">>=", fnc: RshAssign[int], a: 8, b: 3, exp: 1},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			v := c.a
			x := c.fnc(&v, c.b)
			if c.exp != x {
				t.Fatal(x)
			}
		})
	}
}

func TestInc(t *testing.T) {
	cases := []struct {
		name string
		fnc  func(p *int) int
		init int
		exr  int
		val  int
	}{
		{name: "++v", fnc: PreInc[int], init: 0, exr: 1, val: 1},
		{name: "v++", fnc: PostInc[int], init: 0, exr: 0, val: 1},
		{name: "--v", fnc: PreDec[int], init: 0, exr: -1, val: -1},
		{name: "v--", fnc: PostDec[int], init: 0, exr: 0, val: -1},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			v := c.init
			x := c.fnc(&v)
			if c.exr != x {
				t.Fatal(x)
			}
			if c.val != v {
				t.Fatal(v)
			}
		})
	}
}
