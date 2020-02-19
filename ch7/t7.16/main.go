package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"./eval"
)

//!+main

func main() {
	http.HandleFunc("/calc", calc)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//!-main

func calc(w http.ResponseWriter, req *http.Request) {
	expr, err := eval.Parse(req.URL.Query().Get("expr"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "%v\n", err)
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}
	err = expr.Check(map[eval.Var]bool{})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "%v\n", err)
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}
	if len(expr.Variables()) > 0 {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "Invalid expr: %v\n", strings.Join(expr.Variables(), ", "))
		fmt.Fprintf(os.Stderr, "Invalid expr: %v\n", strings.Join(expr.Variables(), ", "))
		return
	}
	fmt.Fprintf(w, "%v\n", expr.Eval(eval.Env(map[eval.Var]float64{})))
}
