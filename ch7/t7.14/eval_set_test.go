package eval

import (
	"fmt"
	"testing"
	"math"
)

func TestSet(t *testing.T) {
	// 	{"sqrt(A / pi)", Env{"A": 87616, "pi": math.Pi}, "167"},
	// 	{"pow(x, 3) + pow(y, 3)", Env{"x": 12, "y": 1}, "1729"},
	// 	{"pow(x, 3) + pow(y, 3)", Env{"x": 9, "y": 10}, "1729"},
	// 	{"5 / 9 * (F - 32)", Env{"F": -40}, "-40"},
	// 	{"5 / 9 * (F - 32)", Env{"F": 32}, "0"},
	// 	{"5 / 9 * (F - 32)", Env{"F": 212}, "100"},
	// 	//!-Eval
	// 	// additional tests that don't appear in the book
	// 	{"-1 + -x", Env{"x": 1}, "-2"},
	// 	{"-1 - x", Env{"x": 1}, "-2"},
	// 	//!+Eval
	{
		x := Var("A")
		y := Var("pi")
		bb := binary{op:'/',x:x,y:y}
		args := make([]Expr,0)
		args = append(args, bb)
		expr := call{fn:"sqrt",args:args}
		fmt.Printf("%v\n", expr)

		env := Env{"A": 87616, "pi": math.Pi}
		fmt.Printf("%v\n", expr.Min(env))
		fmt.Printf("%v\n", expr.Max(env))
	}
	{
		argsLeft := make([]Expr,0)
		argsLeft = append(argsLeft, Var("x"))
		argsLeft = append(argsLeft, literal(3))
		left := call{fn:"pow",args:argsLeft}

		argsRight := make([]Expr,0)
		argsRight = append(argsRight, Var("y"))
		argsRight = append(argsRight, literal(3))
		right := call{fn:"pow",args:argsRight}

		expr := binary{op:'+', x:left, y:right}
		fmt.Printf("%v\n", expr)

		env := Env{"x": 9, "y": 10}
		fmt.Printf("%v\n", expr.Min(env))
		fmt.Printf("%v\n", expr.Max(env))
	}
}