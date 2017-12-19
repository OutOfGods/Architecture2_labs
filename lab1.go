package main

import (
	"./treecontainer"
	"fmt"
)

func searchTest(node *treecontainer.Node) *treecontainer.Node {
	if node.Data == "13323" {
		// fmt.Println(node.Data)
		return node
	}
	return nil
}

func testR(treePtr *treecontainer.Tree) {
	res := treePtr.SearchPre(searchTest, false)
	fmt.Println("R ", res.Data)
}

func testNR(treePtr *treecontainer.Tree) {
	res := treePtr.SearchPreNR(searchTest, false)
	fmt.Println("NR ", res.Data)
}

func testNRA(treePtr *treecontainer.Tree) {
	// result := make(chan *treecontainer.Node)
	res := treePtr.SearchPreNRAsync(searchTest, false)
	// res := <- result
	fmt.Println("NRA ", res.Data)
	// for res := range result {
	// 	fmt.Println("NRA ", res.Data)
	// }
}

func main() {
	var treePtr *treecontainer.Tree = &treecontainer.Tree{}
	treePtr.NewTreeFromFile("./test/generated_test1.txt")
	testR(treePtr)
	testNR(treePtr)
	testNRA(treePtr)
}
