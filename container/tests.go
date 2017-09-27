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

func TestSearchingPrinting(file_name string) {
	var t *Tree = &Tree{}
	t.NewTreeFromFile(file_name)

	fmt.Println()
	fmt.Println("------------------ tree ------------------")

	t.PrintTree()

	fmt.Println("------------------ search pre l2r ------------------")

	n := t.SearchPre(func(n *Node) *Node {
		fmt.Println(n.Data)
		if n.Data == "11332" {
			return n
		}
		return nil
	}, false)

	if n == nil {
		fmt.Print("NOT FOUND  ")
		fmt.Println(n)
	} else {
		fmt.Print("FOUND  ")
		fmt.Println(n.Data)
	}

	fmt.Println("------------------ search post l2r ------------------")

	n = t.SearchPost(func(n *Node) *Node {
		fmt.Println(n.Data)
		if n.Data == "11332" {
			return n
		}
		return nil
	}, false)

	if n == nil {
		fmt.Print("NOT FOUND  ")
		fmt.Println(n)
	} else {
		fmt.Print("FOUND  ")
		fmt.Println(n.Data)
	}
}

// TODO make actual tests
