package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/russross/blackfriday"
)

func main() {
	bs, err := ioutil.ReadFile("README.md")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%q\n", string(bs))
	node := blackfriday.New().Parse(bs)
	fmt.Printf("%v\n", node)
	fmt.Printf("%v\n", node.FirstChild)
	fmt.Printf("%v\n", node.FirstChild.FirstChild)
	fmt.Printf("%v\n", node.FirstChild.FirstChild.FirstChild)
	fmt.Printf("---\n")
	depth := 0
	node.Walk(func(node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
		switch node.Type {
		case blackfriday.Text, blackfriday.Code: // leaf nodes don't get entered, really.
			fmt.Printf("%s%v\n", strings.Repeat("\t", depth+1), node)
		default:
			if entering {
				depth++
				fmt.Printf("%s%v {\n", strings.Repeat("\t", depth), node)
				fmt.Printf("%s<body:%v>\n", strings.Repeat("\t", depth+1), node.Literal)
			} else {
				fmt.Printf("%s} end %v\n", strings.Repeat("\t", depth), node.Type)
				depth--
			}
		}
		return blackfriday.GoToNext
	})
	fmt.Printf("---\n")
	node.Walk(func(node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
		if !entering {
			return blackfriday.GoToNext
		}
		switch node.Type {
		case blackfriday.Heading:
			fmt.Printf("%v -- %v\n", node.Literal, node.FirstChild)
		}
		return blackfriday.GoToNext
	})
}
