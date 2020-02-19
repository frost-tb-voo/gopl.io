package eval

import (
	"fmt"
)

func (v Var) Variables() []string {
	return []string{string(v)}
}

func (l literal) Variables() []string {
	return []string{}
}

func (u unary) Variables() []string {
	return u.x.Variables()
}

func (b binary) Variables() []string {
	return append(b.x.Variables(), b.y.Variables()...)
}

func (c call) Variables() []string {
	switch c.fn {
	case "pow":
		return append(c.args[0].Variables(), c.args[1].Variables()...)
	case "sin":
		return c.args[0].Variables()
	case "sqrt":
		return c.args[0].Variables()
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}

