package libc

import (
	"unsafe"

	"golang.org/x/exp/constraints"
)

type number = interface {
	constraints.Integer | constraints.Float
}

// BoolToInt is a helper that return 1 for true and 0 for false boolean values.
func BoolToInt(v bool) int {
	if v {
		return 1
	}
	return 0
}

// PtrAdd is a generic version of unsafe.Add. It can be used in places where C pointer arithmetic was used.
func PtrAdd[T any](p *T, off int) *T {
	var zero T
	return (*T)(unsafe.Add(unsafe.Pointer(p), off*int(unsafe.Sizeof(zero))))
}

// If is similar to a ternary operator in C. It ALWAYS evaluates side effects.
// For correctly handling side effects, use IfFunc.
func If[T any](cond bool, then, els T) T {
	if cond {
		return then
	}
	return els
}

// IfFunc is similar to a ternary operator in C. It preserves C semantic in terms of expression evaluation.
func IfFunc[T any](cond bool, then, els func() T) T {
	if cond {
		return then()
	}
	return els()
}

// Assign is similar to x = v expression in C, but accepts a pointer to a variable.
func Assign[T any](p *T, v T) T {
	*p = v
	return v
}

// AddAssign is similar to x += v expression in C, but accepts a pointer to a variable.
func AddAssign[T number](p *T, v T) T {
	v = *p + v
	*p = v
	return v
}

// SubAssign is similar to x -= v expression in C, but accepts a pointer to a variable.
func SubAssign[T number](p *T, v T) T {
	v = *p - v
	*p = v
	return v
}

// MulAssign is similar to x *= v expression in C, but accepts a pointer to a variable.
func MulAssign[T number](p *T, v T) T {
	v = *p * v
	*p = v
	return v
}

// DivAssign is similar to x /= v expression in C, but accepts a pointer to a variable.
func DivAssign[T number](p *T, v T) T {
	v = *p / v
	*p = v
	return v
}

// ModAssign is similar to x %= v expression in C, but accepts a pointer to a variable.
func ModAssign[T constraints.Integer](p *T, v T) T {
	v = *p % v
	*p = v
	return v
}

// AndAssign is similar to x &= v expression in C, but accepts a pointer to a variable.
func AndAssign[T constraints.Integer](p *T, v T) T {
	v = *p & v
	*p = v
	return v
}

// OrAssign is similar to x |= v expression in C, but accepts a pointer to a variable.
func OrAssign[T constraints.Integer](p *T, v T) T {
	v = *p | v
	*p = v
	return v
}

// XorAssign is similar to x ^= v expression in C, but accepts a pointer to a variable.
func XorAssign[T constraints.Integer](p *T, v T) T {
	v = *p ^ v
	*p = v
	return v
}

// LshAssign is similar to x <<= v expression in C, but accepts a pointer to a variable.
func LshAssign[T constraints.Integer](p *T, v T) T {
	v = *p << v
	*p = v
	return v
}

// RshAssign is similar to x >>= v expression in C, but accepts a pointer to a variable.
func RshAssign[T constraints.Integer](p *T, v T) T {
	v = *p >> v
	*p = v
	return v
}

// PreInc is similar to ++x expression in C, but accepts a pointer to a variable.
func PreInc[T number](p *T) T {
	*p++
	return *p
}

// PostInc is similar to x++ expression in C, but accepts a pointer to a variable.
func PostInc[T number](p *T) T {
	v := *p
	*p++
	return v
}

// PreDec is similar to --x expression in C, but accepts a pointer to a variable.
func PreDec[T number](p *T) T {
	*p--
	return *p
}

// PostDec is similar to x-- expression in C, but accepts a pointer to a variable.
func PostDec[T number](p *T) T {
	v := *p
	*p--
	return v
}
