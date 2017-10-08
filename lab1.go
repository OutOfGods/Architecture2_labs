package main

import (
	"./treecontainer"
	"fmt"

	"strconv"
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
	treePtr.NewTreeFromFile("./test/generated_test.txt")

	ch := make(chan int, 1000)
	var wg sync.WaitGroup
	treePtr.TraversePreAsync(func(n *treecontainer.Node) {
		// fmt.Println(n.Data)
		if len(ch) != cap(ch) {
			i, _ := strconv.Atoi(n.Data.(string))
			ch <- i / 1000
		} else {
			panic("no space left in buffer")
		}

	}, false, &wg)

	wg.Wait()

	for len(ch) != 0 {
		fmt.Println(<-ch)
	}

	// fmt.Println(len(ch))

	// for len(ch) > 0 {
	// 	fmt.Println(<-ch)
	// }
}
