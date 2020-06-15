// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 359.

// Package equal provides a deep equivalence relation for arbitrary values.
package equal

import (
	"reflect"
	"unsafe"
)

const BillionDiff float64 = 1.0 / 1000000000

//!+
func equal(x, y reflect.Value, seen map[comparison]bool) bool {
	if !x.IsValid() || !y.IsValid() {
		return x.IsValid() == y.IsValid()
	}
	if x.Type() != y.Type() {
		return false
	}

	// ...cycle check omitted (shown later)...

	//!-
	//!+cyclecheck
	// cycle check
	if x.CanAddr() && y.CanAddr() {
		xptr := unsafe.Pointer(x.UnsafeAddr())
		yptr := unsafe.Pointer(y.UnsafeAddr())
		if xptr == yptr {
			return true // identical references
		}
		c := comparison{xptr, yptr, x.Type()}
		if seen[c] {
			return true // already seen
		}
		seen[c] = true
	}
	//!-cyclecheck
	//!+
	switch x.Kind() {
	case reflect.Bool:
		return x.Bool() == y.Bool()

	case reflect.String:
		return x.String() == y.String()

	// ...numeric cases omitted for brevity...

	//!-
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
		reflect.Int64:
		return x.Int() == y.Int() || x.Int() < y.Int() && float64(y.Int()-x.Int()) < BillionDiff || x.Int() > y.Int() && float64(x.Int()-y.Int()) < BillionDiff

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32,
		reflect.Uint64, reflect.Uintptr:
		return x.Uint() == y.Uint() || x.Uint() < y.Uint() && float64(y.Uint()-x.Uint()) < BillionDiff || x.Uint() > y.Uint() && float64(x.Uint()-y.Uint()) < BillionDiff

	case reflect.Float32, reflect.Float64:
		return x.Float() == y.Float() || x.Float() < y.Float() && float64(y.Float()-x.Float()) < BillionDiff || x.Float() > y.Float() && float64(x.Float()-y.Float()) < BillionDiff

	case reflect.Complex64, reflect.Complex128:
		if x.Complex() == y.Complex() {
			return true
		}
		matchReal := real(x.Complex()) == real(y.Complex()) || real(x.Complex()) < real(y.Complex()) && real(y.Complex())-real(x.Complex()) < BillionDiff || real(x.Complex()) > real(y.Complex()) && real(x.Complex())-real(y.Complex()) < BillionDiff
		matchImag := imag(x.Complex()) == imag(y.Complex()) || imag(x.Complex()) < imag(y.Complex()) && imag(y.Complex())-imag(x.Complex()) < BillionDiff || imag(x.Complex()) > imag(y.Complex()) && imag(x.Complex())-imag(y.Complex()) < BillionDiff
		return matchReal && matchImag
	//!+
	case reflect.Chan, reflect.UnsafePointer, reflect.Func:
		return x.Pointer() == y.Pointer()

	case reflect.Ptr, reflect.Interface:
		return equal(x.Elem(), y.Elem(), seen)

	case reflect.Array, reflect.Slice:
		if x.Len() != y.Len() {
			return false
		}
		for i := 0; i < x.Len(); i++ {
			if !equal(x.Index(i), y.Index(i), seen) {
				return false
			}
		}
		return true

	// ...struct and map cases omitted for brevity...
	//!-
	case reflect.Struct:
		for i, n := 0, x.NumField(); i < n; i++ {
			if !equal(x.Field(i), y.Field(i), seen) {
				return false
			}
		}
		return true

	case reflect.Map:
		if x.Len() != y.Len() {
			return false
		}
		for _, k := range x.MapKeys() {
			if !equal(x.MapIndex(k), y.MapIndex(k), seen) {
				return false
			}
		}
		return true
		//!+
	}
	panic("unreachable")
}

//!-

//!+comparison
// Equal reports whether x and y are deeply equal.
//!-comparison
//
// Map keys are always compared with ==, not deeply.
// (This matters for keys containing pointers or interfaces.)
//!+comparison
func Equal(x, y interface{}) bool {
	seen := make(map[comparison]bool)
	return equal(reflect.ValueOf(x), reflect.ValueOf(y), seen)
}

type comparison struct {
	x, y unsafe.Pointer
	t    reflect.Type
}

//!-comparison
