package libc

import "testing"

func TestStrLen(t *testing.T) {
	if StrLen((*byte)(nil)) != 0 {
		t.Fatal()
	}
	var buf [5]byte

	buf = [5]byte{}
	if StrLen(&buf[0]) != 0 {
		t.Fatal()
	}

	buf = [5]byte{'a', 'b', 0, 'c', 0}
	if StrLen(&buf[0]) != 2 {
		t.Fatal()
	}
	if StrLen(&buf[3]) != 1 {
		t.Fatal()
	}
}

func TestStrLenN(t *testing.T) {
	if StrNLen((*byte)(nil), 0) != 0 {
		t.Fatal()
	}
	if StrNLen((*byte)(nil), 5) != 0 {
		t.Fatal()
	}
	var buf [5]byte

	buf = [5]byte{}
	if StrNLen(&buf[0], 5) != 0 {
		t.Fatal()
	}

	buf = [5]byte{'a', 'b', 0, 'c', 0}
	if StrNLen(&buf[0], 5) != 2 {
		t.Fatal()
	}
	if StrNLen(&buf[0], 1) != 1 {
		t.Fatal()
	}
	if StrNLen(&buf[0], 2) != 2 {
		t.Fatal()
	}
	if StrNLen(&buf[3], 2) != 1 {
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
	s := StrNSlice((*byte)(nil), 0)
	if s != nil {
		t.Fatal()
	}
	s = StrNSlice((*byte)(nil), 5)
	if s != nil {
		t.Fatal()
	}
	var buf [5]byte
	s = StrNSlice(&buf[0], 5)
	if s == nil || len(s) != 0 || cap(s) != 5 {
		t.Fatal()
	}
	buf[0] = 'a'
	buf[1] = 'b'
	s = StrNSlice(&buf[0], 5)
	if s == nil || len(s) != 2 || cap(s) != 5 {
		t.Fatal(len(s), cap(s))
	}
}
