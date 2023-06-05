package libc

import "testing"

func TestCString(t *testing.T) {
	t.Run("ptr", testCStringPtr)
	t.Run("slice", testCStringSlice)
}

func testCStringPtr(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		c := CString("")
		if c == nil {
			t.Fatal()
		}
		t.Cleanup(func() {
			Free(c)
		})
		if s := GoString(c); s != "" {
			t.Fatal(s)
		}
	})
	t.Run("str", func(t *testing.T) {
		c := CString("abc")
		t.Cleanup(func() {
			Free(c)
		})
		if s := GoString(c); s != "abc" {
			t.Fatal(s)
		}
	})
	t.Run("zero", func(t *testing.T) {
		c := CString("ab\x00c")
		t.Cleanup(func() {
			Free(c)
		})
		if s := GoString(c); s != "ab" {
			t.Fatal(s)
		}
	})
}

func testCStringSlice(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		b := CStringS("")
		defer FreeS(b)
		if cap(b) != 1 || len(b) != 0 {
			t.Fatal()
		}
		if s := string(b[:cap(b)]); s != "\x00" {
			t.Fatal(s)
		}
		if s := GoStringS(b); s != "" {
			t.Fatal(s)
		}
	})
	t.Run("str", func(t *testing.T) {
		b := CStringS("abc")
		defer FreeS(b)
		if cap(b) != 4 || len(b) != 3 {
			t.Fatal()
		}
		if s := string(b[:cap(b)]); s != "abc\x00" {
			t.Fatal(s)
		}
		if s := GoStringS(b); s != "abc" {
			t.Fatal(s)
		}
	})
	t.Run("zero", func(t *testing.T) {
		b := CStringS("ab\x00c")
		defer FreeS(b)
		if cap(b) != 3 || len(b) != 2 {
			t.Fatal()
		}
		if s := string(b[:cap(b)]); s != "ab\x00" {
			t.Fatal(s)
		}
		if s := GoStringS(b); s != "ab" {
			t.Fatal(s)
		}
	})
}

func TestCString16(t *testing.T) {
	t.Run("ptr", testCString16Ptr)
	t.Run("slice", testCString16Slice)
}

func testCString16Ptr(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		c := CString16("")
		if c == nil {
			t.Fatal()
		}
		t.Cleanup(func() {
			Free(c)
		})
		if s := GoString16(c); s != "" {
			t.Fatal(s)
		}
	})
	t.Run("str", func(t *testing.T) {
		c := CString16("abc")
		t.Cleanup(func() {
			Free(c)
		})
		if s := GoString16(c); s != "abc" {
			t.Fatal(s)
		}
	})
	t.Run("zero", func(t *testing.T) {
		c := CString16("ab\x00c")
		t.Cleanup(func() {
			Free(c)
		})
		if s := GoString16(c); s != "ab" {
			t.Fatal(s)
		}
	})
}

func testCString16Slice(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		b := CStringS16("")
		defer FreeS(b)
		if cap(b) != 1 || len(b) != 0 {
			t.Fatal()
		}
		if s := (*[1]uint16)(b[:cap(b)]); *s != [1]uint16{0} {
			t.Fatal(s)
		}
		if s := GoStringS16(b); s != "" {
			t.Fatal(s)
		}
	})
	t.Run("str", func(t *testing.T) {
		b := CStringS16("abc")
		defer FreeS(b)
		if cap(b) != 4 || len(b) != 3 {
			t.Fatal()
		}
		if s := (*[4]uint16)(b[:cap(b)]); *s != [4]uint16{'a', 'b', 'c', 0} {
			t.Fatal(s)
		}
		if s := GoStringS16(b); s != "abc" {
			t.Fatal(s)
		}
	})
	t.Run("zero", func(t *testing.T) {
		b := CStringS16("ab\x00c")
		defer FreeS(b)
		if cap(b) != 3 || len(b) != 2 {
			t.Fatal()
		}
		if s := (*[3]uint16)(b[:cap(b)]); *s != [3]uint16{'a', 'b', 0} {
			t.Fatal(s)
		}
		if s := GoStringS16(b); s != "ab" {
			t.Fatal(s)
		}
	})
}
