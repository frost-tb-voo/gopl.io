// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 333.

// Package display provides a means to display structured data.
package display

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

//!+Display

func Display(name string, x interface{}) {
	fmt.Printf("Display %s (%T):\n", name, x)
	display(name, reflect.ValueOf(x))
}

//!-Display

// formatAtom formats a value without inspecting its internal structure.
// It is a copy of the the function in gopl.io/ch11/format.
func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	// ...floating-point and complex cases omitted for brevity...
	case reflect.Bool:
		if v.Bool() {
			return "true"
		}
		return "false"
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr,
		reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" +
			strconv.FormatUint(uint64(v.Pointer()), 16)
	default: // reflect.Array, reflect.Struct, reflect.Interface
		return v.Type().String() + " value"
	}
}

//!+display
func display(path string, v reflect.Value) {
	fmt.Println(strings.Join(sdisplay(path, v, 0, 10), "\n"))
}

func join(pathes []string) string {
	return "(" + strings.Join(pathes, ", ") + ")"
}

func sdisplay(path string, v reflect.Value, current, max int) []string {
	if current > max {
		return []string{}
	}
	switch v.Kind() {
	case reflect.Invalid:
		return []string{fmt.Sprintf("%s = invalid", path)}
	case reflect.Slice, reflect.Array:
		results := make([]string, 0)
		for i := 0; i < v.Len(); i++ {
			results = append(results, sdisplay(fmt.Sprintf("%s[%d]", path, i), v.Index(i), current+1, max)...)
		}
		return results
	case reflect.Struct:
		results := make([]string, 0)
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			results = append(results, sdisplay(fieldPath, v.Field(i), current+1, max)...)
		}
		return results
	case reflect.Map:
		results := make([]string, 0)
		for _, key := range v.MapKeys() {
			// results = append(results, sdisplay(fmt.Sprintf("%s[%s]", path,
			// 	formatAtom(key)), v.MapIndex(key))...)
			results = append(results, sdisplay(fmt.Sprintf("%s[%s]", path,
				sdisplay("key", key, current+1, max)), v.MapIndex(key), current+1, max)...)
		}
		return results
	case reflect.Ptr:
		if v.IsNil() {
			return []string{fmt.Sprintf("%s = nil", path)}
		}
		return sdisplay(fmt.Sprintf("(*%s)", path), v.Elem(), current+1, max)
	case reflect.Interface:
		if v.IsNil() {
			return []string{fmt.Sprintf("%s = nil", path)}
		}
		results := make([]string, 0)
		result := fmt.Sprintf("%s.type = %s", path, v.Elem().Type())
		results = append(results, result)
		results = append(results, sdisplay(path+".value", v.Elem(), current+1, max)...)
		return results
	default: // basic types, channels, funcs
	}
	return []string{fmt.Sprintf("%s = %s", path, formatAtom(v))}
}

//!-display
