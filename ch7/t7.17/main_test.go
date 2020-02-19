package main

import (
	"testing"
	"os"
)

func TestXmlSelect1(tt *testing.T) {
	file, err := os.Open("./test.xml")
	if err != nil {
		tt.Fatal(err)
		tt.Fail()
		return
	}
	defer file.Close()
	query := []string{"div", "div", "h2"}
	XmlSelect(file, query)

	// Output:
}

func TestXmlSelect2(tt *testing.T) {
	file, err := os.Open("./test.xml")
	if err != nil {
		tt.Fatal(err)
		tt.Fail()
		return
	}
	defer file.Close()
	query := []string{"div", "div@[class=div1]", "a@[href=#dt-compat"}
	XmlSelect(file, query)

	// Output:
	// TBD
}

func TestParseSelector(tt *testing.T) {
	selector := ParseSelector("h2@[id=determinism]")
	tt.Logf("%v\n", selector)

	// Output:
	// {h2 map[id:determinism]}
}
