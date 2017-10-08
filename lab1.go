package main

import (
	"./treecontainer"
	"fmt"

	"sync"
)

func dosearch() {
	var treePtr *treecontainer.Tree = &treecontainer.Tree{}
	treePtr.NewTreeFromFile("./test/test_input_1.txt")
	treePtr.PrintTree()
	treePtr.TraverseBF(func(n *treecontainer.Node) { fmt.Println(n.Data) }, false)
	treePtr.SearchPre(
		func(node *treecontainer.Node) *treecontainer.Node {
			if node.Data == "13323" {
				return node
			}
			return nil
		},
		false)
}

func main() {
	var treePtr *treecontainer.Tree = &treecontainer.Tree{}
	treePtr.NewTreeFromFile("./test/gen_test.txt")

	var wg sync.WaitGroup
	treePtr.TraversePreAsync(func(n *treecontainer.Node) {
		fmt.Println(n.Data)
	}, false, &wg)

	// fmt.Println(len(ch))

	// for len(ch) > 0 {
	// 	fmt.Println(<-ch)
	// }

	wg.Wait()
}
