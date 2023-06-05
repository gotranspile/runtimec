package libc

import "testing"

func TestStrChar(t *testing.T) {
	var s [6]byte

	// Search should be able to find zero terminator.
	s = [6]byte{'a', 'b', 'b', 0, 'c', 0}
	i := StrIndexS(s[:], 0)
	if i != 3 {
		t.Fatal(i)
	}
	i = StrIndex(&s[0], 0)
	if i != 3 {
		t.Fatal(i)
	}
	r := StrCharS(s[:], 0)
	if len(r) != 3 || &r[0] != &s[3] {
		t.Fatal(r)
	}
	p := StrChar(&s[0], 0)
	if p != &s[3] {
		t.Fatal(p)
	}

	// Otherwise, search should stop on zero terminator.
	s = [6]byte{'a', 'b', 'b', 0, 'c', 0}
	i = StrIndexS(s[:], 'c')
	if i != -1 {
		t.Fatal(i)
	}
	i = StrIndex(&s[0], 'c')
	if i != -1 {
		t.Fatal(i)
	}
	r = StrCharS(s[:], 'c')
	if r != nil {
		t.Fatal(r)
	}
	p = StrChar(&s[0], 'c')
	if p != nil {
		t.Fatal()
	}

	// Search should return first encountered value.
	s = [6]byte{'a', 'b', 'b', 0, 'b', 0}
	i = StrIndexS(s[:], 'b')
	if i != 1 {
		t.Fatal(i)
	}
	i = StrIndex(&s[0], 'b')
	if i != 1 {
		t.Fatal(i)
	}
	r = StrCharS(s[:], 'b')
	if len(r) != 5 || &r[0] != &s[1] {
		t.Fatal(r)
	}
	p = StrChar(&s[0], 'b')
	if p != &s[1] {
		t.Fatal()
	}
}

func TestStrRChar(t *testing.T) {
	var s [6]byte

	// Search should be able to find zero terminator.
	s = [6]byte{'a', 'b', 'b', 0, 'c', 0}
	i := StrLastIndexS(s[:], 0)
	if i != 3 {
		t.Fatal(i)
	}
	i = StrLastIndex(&s[0], 0)
	if i != 3 {
		t.Fatal(i)
	}
	r := StrRCharS(s[:], 0)
	if len(r) != 3 || &r[0] != &s[3] {
		t.Fatal(r)
	}
	p := StrRChar(&s[0], 0)
	if p != &s[3] {
		t.Fatal(p)
	}

	// Otherwise, search should stop on zero terminator.
	s = [6]byte{'a', 'b', 'b', 0, 'c', 0}
	i = StrLastIndexS(s[:], 'c')
	if i != -1 {
		t.Fatal(i)
	}
	i = StrLastIndex(&s[0], 'c')
	if i != -1 {
		t.Fatal(i)
	}
	r = StrRCharS(s[:], 'c')
	if r != nil {
		t.Fatal(r)
	}
	p = StrRChar(&s[0], 'c')
	if p != nil {
		t.Fatal()
	}

	// Search should return last encountered value.
	s = [6]byte{'a', 'b', 'b', 0, 'b', 0}
	i = StrLastIndexS(s[:], 'b')
	if i != 2 {
		t.Fatal(i)
	}
	i = StrLastIndex(&s[0], 'b')
	if i != 2 {
		t.Fatal(i)
	}
	r = StrRCharS(s[:], 'b')
	if len(r) != 4 || &r[0] != &s[2] {
		t.Fatal(r)
	}
	p = StrRChar(&s[0], 'b')
	if p != &s[2] {
		t.Fatal()
	}
}
