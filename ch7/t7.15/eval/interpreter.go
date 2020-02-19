package eval

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

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
		expr2, err := Parse(exprStr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			exprStr = ""
			continue
		}
		expr = expr2
		break
	}
	env := Env(map[Var]float64{})
	for _, variable := range(expr.Variables()) {
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
