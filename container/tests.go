package container

import "fmt"

func TestNewTreeFromFile(file_name string) {
	var treePtr *Tree = &Tree{}
	treePtr.NewTreeFromFile(file_name)
	fmt.Println()
	treePtr.PrintTree()
}
