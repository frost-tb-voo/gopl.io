package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {
	callees := map[string]bool{}
	callers := getCallers(*callee, callees)
	fmt.Printf("%v\n", callers)
}

var (
	callee *string
)

type Package struct {
	Path    string
	Name    string
	Imports []string
	Deps    []string
}

func init() {
	callee = flag.String("callee", "fmt", "callee package name")
	flag.Parse()
}

func getCallers(callee string, callees map[string]bool) (callers []string) {
	cmd := exec.Command("go", "list", "-json", callee)
	cmd.Stdin = strings.NewReader("")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
		return
	}
	data := out.String()
	// fmt.Printf("%v\n", data)

	var pkg Package
	json.Unmarshal([]byte(data), &pkg)
	// fmt.Printf("%v\n", pkg)
	callees[callee] = true
	for _, subPkg := range pkg.Deps {
		if callees[subPkg] {
			continue
		}
		// fmt.Printf("%v\n", subPkg)
		callers = append(callers, subPkg)
		subCallers := getCallers(subPkg, callees)
		callers = append(callers, subCallers...)
	}
	return
}
