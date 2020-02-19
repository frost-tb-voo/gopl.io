// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 198.

// Package eval provides an expression evaluator.
package eval

import (
	"fmt"
	"math"
)

//!+Min1

func (v Var) Min(env Env) float64 {
	return env[v]
}

func (l literal) Min(_ Env) float64 {
	return float64(l)
}

//!-Min1

//!+Min2

func (u unary) Min(env Env) float64 {
	return u.x.Min(env)
}

func (b binary) Min(env Env) float64 {
	return math.Min(b.x.Min(env), b.y.Min(env))
}

func (c call) Min(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Min(c.args[0].Min(env), c.args[1].Min(env))
	case "sin":
		return c.args[0].Min(env)
	case "sqrt":
		return c.args[0].Min(env)
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}

//!-Min2
