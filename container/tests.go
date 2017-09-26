package container

import "fmt"

func TestNewTreeFromFilePrinting(file_name string) {
	var treePtr *Tree = &Tree{}
	treePtr.NewTreeFromFile(file_name)
	fmt.Println()
	treePtr.PrintTree()
}

func TestTraversalsPrinting(file_name string) {
	var t *Tree = &Tree{}
	t.NewTreeFromFile(file_name)

	fmt.Println()
	fmt.Println("------------------ tree ------------------")

	t.PrintTree()

	fmt.Println("------------------ pre l2r ------------------")

	t.TraversePre(func(n *Node) { fmt.Println(n.Data) }, false)

	fmt.Println("------------------ pre r2l ------------------")

	t.TraversePre(func(n *Node) { fmt.Println(n.Data) }, true)

	fmt.Println("------------------ post l2r ------------------")

	t.TraversePost(func(n *Node) { fmt.Println(n.Data) }, false)

	fmt.Println("------------------ post r2l ------------------")

	t.TraversePost(func(n *Node) { fmt.Println(n.Data) }, true)
}

// TODO make actual tests
