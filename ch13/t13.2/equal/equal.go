package equal

import (
	"reflect"
	"unsafe"
)

//!+
func cyclecheck(x reflect.Value, seen map[comparison]bool) bool {
	// ...cycle check omitted (shown later)...

	//!-
	//!+cyclecheck
	// cycle check
	if x.CanAddr() {
		xptr := unsafe.Pointer(x.UnsafeAddr())
		c := comparison{xptr, x.Type()}
		if seen[c] {
			return true // already seen
		}
		seen[c] = true
	}
	//!-cyclecheck
	//!+
	switch x.Kind() {
	case reflect.Bool:
		return false

	case reflect.String:
		return false

	// ...numeric cases omitted for brevity...

	//!-
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
		reflect.Int64:
		return false

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32,
		reflect.Uint64, reflect.Uintptr:
		return false

	case reflect.Float32, reflect.Float64:
		return false

	case reflect.Complex64, reflect.Complex128:
		return false
	//!+
	case reflect.Chan, reflect.UnsafePointer, reflect.Func:
		return false

	case reflect.Ptr, reflect.Interface:
		return cyclecheck(x.Elem(), seen)

	case reflect.Array, reflect.Slice:
		for i := 0; i < x.Len(); i++ {
			if cyclecheck(x.Index(i), seen) {
				return true
			}
		}
		return false

	// ...struct and map cases omitted for brevity...
	//!-
	case reflect.Struct:
		for i, n := 0, x.NumField(); i < n; i++ {
			if cyclecheck(x.Field(i), seen) {
				return true
			}
		}
		return false

	case reflect.Map:
		for _, k := range x.MapKeys() {
			if cyclecheck(x.MapIndex(k), seen) {
				return true
			}
		}
		return false
		//!+
	}
	panic("unreachable")
}

//!-

func Cyclecheck(x interface{}) bool {
	seen := make(map[comparison]bool)
	return cyclecheck(reflect.ValueOf(x), seen)
}

type comparison struct {
	x unsafe.Pointer
	t reflect.Type
}
