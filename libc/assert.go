package libc

// Assert panics when the condition is false.
func Assert(cond bool) {
	if !cond {
		panic("assert failed")
	}
}
