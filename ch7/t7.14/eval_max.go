// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 198.

// Package eval provides an expression evaluator.
package eval

import (
	"fmt"
	"math"
)

//!+env

type Env map[Var]float64

//!-env

//!+Max1

func (v Var) Max(env Env) float64 {
	return env[v]
}

func (l literal) Max(_ Env) float64 {
	return float64(l)
}

//!-Max1

//!+Max2

func (u unary) Max(env Env) float64 {
	return u.x.Max(env)
}

func (b binary) Max(env Env) float64 {
	return math.Max(b.x.Max(env), b.y.Max(env))
}

func (c call) Max(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Max(c.args[0].Max(env), c.args[1].Max(env))
	case "sin":
		return c.args[0].Max(env)
	case "sqrt":
		return c.args[0].Max(env)
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}

//!-Max2
