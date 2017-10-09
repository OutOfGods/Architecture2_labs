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

func t_async() {
	var treePtr *treecontainer.Tree = &treecontainer.Tree{}
	treePtr.NewTreeFromFile("./test/generated_test.txt")

	ch := make(chan int, 48828125)

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

	c := 0
	for len(ch) != 0 {
		<-ch
		c++
	}
	// fmt.Println(c)
}

func t_sync() {
	var treePtr *treecontainer.Tree = &treecontainer.Tree{}
	treePtr.NewTreeFromFile("./test/generated_test.txt")

	var lst []int

	treePtr.TraversePre(func(n *treecontainer.Node) {
		i, _ := strconv.Atoi(n.Data.(string))
		lst = append(lst, i/1000)
	}, false)

	// fmt.Println(len(lst))
}

func main() {
	// t_sync()
	t_async()
}
