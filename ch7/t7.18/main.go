package main

import ("encoding/xml"
"io"
"os"
"fmt")

type Node interface{} // CharData or *Element
type CharData string
type Element struct {
	Type xml.Name
	Attr []xml.Attr
	Children []Node
}

func main() {
	CreateNode(os.Stdin)
}

func CreateNode(reader io.Reader) (Node, error) {
	dec := xml.NewDecoder(reader)
	var stack []Node // stack of element names
	root := Element{}
	stack = append(stack, root)
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "CreateNode: %v\n", err)
			return nil, err
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			newNode := Element{Type:tok.Name, Attr:tok.Attr, Children:[]Node{}}
			currentNode := stack[len(stack)-1]
			switch currentNode := currentNode.(type) {
			case Element:
				fmt.Fprintf(os.Stderr, "CreateNode: append %v to %v\n", newNode, currentNode)
				currentNode.Children = append(currentNode.Children, newNode)
			default:
				fmt.Fprintf(os.Stderr, "CreateNode: invalid %T %v\n", tok, tok)
				return nil, fmt.Errorf("CreateNode: invalid %T %v", tok, tok)
			}
			stack = append(stack, newNode) // push
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			if len(stack) > 0 {
				currentNode := stack[len(stack)-1]
				switch currentNode := currentNode.(type) {
				case Element:
					currentNode.Children = append(currentNode.Children, CharData(string([]byte(tok))))
				}
			} else {
				fmt.Fprintf(os.Stderr, "CreateNode: invalid %T %v '%v'\n", tok, tok, string([]byte(tok)))
				return nil, fmt.Errorf("CreateNode: invalid %T %v '%v'", tok, tok, string([]byte(tok)))
			}
		}
	}
	return root, nil
}


