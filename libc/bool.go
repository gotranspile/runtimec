package libc

// BoolToInt is a helper that return 1 for true and 0 for false boolean values.
func BoolToInt(v bool) int {
	if v {
		return 1
	}
	return 0
}
