package main

import (
	"fmt"
	"lab1/container"
)

func main() {
	fmt.Println("main")
	container.TestNewTreeFromFile("tree_input_tests/test_input_3.txt")

	var t *container.Tree = &container.Tree{}
	t.NewTreeFromFile("tree_input_tests/test_input_3.txt")
	fmt.Println()
	t.TraverseDF(func(n *container.Node) { fmt.Println(n.Data) })

}
