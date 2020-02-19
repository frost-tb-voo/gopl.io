package main

import (
	"testing"
	"os"
)

func TestCreateNode(tt *testing.T) {
	file, err := os.Open("./test.xml")
	if err != nil {
		tt.Fatal(err)
		tt.Fail()
		return
	}
	defer file.Close()
	expr, err := CreateNode(file)
	if err != nil {
		tt.Fatal(err)
		tt.Fail()
		return
	}
	tt.Logf("%v\n", expr)

	// Output:
}

