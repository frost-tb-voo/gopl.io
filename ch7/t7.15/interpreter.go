package eval

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
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

func Interpreter() {
	var exprStr string
	if len(os.Args) > 1 {
		exprStr = strings.Join(os.Args[1:], " ")
	}
	var expr Expr
	input := bufio.NewScanner(os.Stdin)
	for ;; {
		if exprStr == "" {
			fmt.Printf("Input new expression:\n")
			input.Scan()
			exprStr = input.Text()
		}
		expr, err := Parse(exprStr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			exprStr = ""
			continue
		}
		break
	}
	var env Env
	for variable := range(expr.Variables()) {
		for ;; {
			var valueStr string
			if valueStr == "" {
				fmt.Printf("Input value of %v:\n", variable)
				input.Scan()
				valueStr = input.Text()
			}
			value, err := strconv.ParseFloat(valueStr, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				valueStr = ""
				continue
			}
			env[Var(variable)] = value
			break
		}
	}
	fmt.Printf("Thanks, %v = %v\n", expr, expr.Eval(env))
}
