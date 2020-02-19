package eval

import (
	"fmt"
	"strconv"
)

//!+String1

func (v Var) String() string {
	return string(v)
}

func (l literal) String() string {
	return strconv.FormatFloat(float64(l), 'f', -1, 64)
}

//!-String1

//!+String2

func (u unary) String() string {
	switch u.op {
	case '+':
		return "+ (" + u.x.String() + ")"
	case '-':
		return "- (" + u.x.String() + ")"
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

func (b binary) String() string {
	switch b.op {
	case '+':
		return "(" + b.x.String() + ") + (" + b.y.String() + ")"
	case '-':
		return "(" + b.x.String() + ") - (" + b.y.String() + ")"
	case '*':
		return "(" + b.x.String() + ") * (" + b.y.String() + ")"
	case '/':
		return "(" + b.x.String() + ") / (" + b.y.String() + ")"
	}
	panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
}

func (c call) String() string {
	switch c.fn {
	case "pow":
		return "pow((" + c.args[0].String() + "), (" + c.args[1].String() + "))"
	case "sin":
		return "sin((" + c.args[0].String() + "))"
	case "sqrt":
		return "sqrt((" + c.args[0].String() + "))"
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}

//!-String2
