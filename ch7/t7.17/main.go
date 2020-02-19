package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	XmlSelect(os.Stdin, os.Args[1:])
}

func XmlSelect(reader io.Reader, query []string) {
	dec := xml.NewDecoder(reader)
	var stack []xml.StartElement // stack of element names
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok) // push
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			if containsAll(stack, query) {
				fmt.Printf("%s: %s\n", Join(stack, " "), tok)
			}
		}
	}
}

func Join(x []xml.StartElement, sep string) string {
	text := ""
	for _, element := range(x) {
		if len(text) > 0 {
			text += sep
		}
		text += element.Name.Local
	}
	return text
}

// containsAll reports whether x contains the elements of y, in order.
func containsAll(x []xml.StartElement, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if Equal(x[0], y[0]) {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}

func Equal(element xml.StartElement, query string) bool {
	selector := ParseSelector(query)
	if element.Name.Local != selector.Name {
		return false
	}
	for key, value := range(selector.Attributes) {
		ok := false
		for _, attr := range(element.Attr) {
			if attr.Name.Local == key && attr.Value == value {
				ok = true
				break
			}
		}
		if !ok {
			return false
		}
	}
	return true
}

func ParseSelector(query string) Selector {
	ss := strings.Split(query, "@")
	selector := Selector{Attributes:map[string]string{}}
	selector.Name = ss[0]
	if len(ss) == 2 && len(ss[1]) > 0 {
		attr := ss[1]
		attr = strings.TrimPrefix(attr, "[")
		attr = strings.TrimSuffix(attr, "]")
		for _, conditionTemp := range(strings.Split(attr, ",")) {
			condition := strings.Split(conditionTemp, "=")
			selector.Attributes[condition[0]] = condition[1]
		}
	}
	return selector
}

type Selector struct {
	Name string
	Attributes map[string]string
}
