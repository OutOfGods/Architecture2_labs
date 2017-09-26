package container

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	//Key interface{}
	Data     interface{}
	Children []*Node
	Parent   *Node
}

type Tree struct {
	Root        *Node
	CurrNodePtr *Node
}

func (tree *Tree) NewTreeFromFile(file_name string) {
	fmt.Println("NewTreeFromFile function started")
	fmt.Printf("Filename: %s", file_name)
	file, err := os.Open(file_name)
	if err != nil {
		fmt.Println("Error: file not found!")
		return
	}

	prev_level := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		read_line := scanner.Text()
		if read_line == "" {
			break
		}

		node_text := strings.Split(read_line, " ")
		curr_level, data := len(node_text[0]), node_text[1]

		if tree.Root == nil {
			//fmt.Println("Root initialization")
			tree.Root = &Node{Data: data, Parent: nil}
			tree.CurrNodePtr = tree.Root
		} else {
			//tree.PrintTree()
			if curr_level >= prev_level {
				if curr_level == prev_level {
					tree.CurrNodePtr = tree.CurrNodePtr.Parent
				}
				tree.CurrNodePtr.Children = append(tree.CurrNodePtr.Children, &Node{Data: data, Parent: tree.CurrNodePtr})
				tree.CurrNodePtr = tree.CurrNodePtr.Children[len(tree.CurrNodePtr.Children)-1]
			} else {
				for i := 0; i < prev_level-curr_level+1; i++ {
					tree.CurrNodePtr = tree.CurrNodePtr.Parent
				}
				tree.CurrNodePtr.Children = append(tree.CurrNodePtr.Children, &Node{Data: data, Parent: tree.CurrNodePtr})
				tree.CurrNodePtr = tree.CurrNodePtr.Children[len(tree.CurrNodePtr.Children)-1]
			}
			prev_level = curr_level
			//fmt.Println("*********************************************")
		}
	}

	file.Close()
}

// iterate children left to right and right to left

func (node *Node) IterateChildren(fn func(*Node), r2l bool) { // right to left
	if r2l {
		for i := len(node.Children) - 1; i >= 0; i-- {
			if node.Children[i] != nil {
				fn(node.Children[i])
			}
		}
	} else {
		for i := 0; i < len(node.Children); i++ {
			if node.Children[i] != nil {
				fn(node.Children[i])
			}
		}
	}
}

// depth first pre-order

func (node *Node) TraversePre(fn func(*Node), r2l bool) {
	fn(node)
	node.IterateChildren(func(n *Node) { n.TraversePre(fn, r2l) }, r2l)
}

func (tree *Tree) TraversePre(fn func(*Node), r2l bool) {
	if tree.Root != nil {
		tree.Root.TraversePre(fn, r2l)
	}
}

// depth first post-order

func (node *Node) TraversePost(fn func(*Node), r2l bool) {
	node.IterateChildren(func(n *Node) { n.TraversePost(fn, r2l) }, r2l)
	fn(node)
}

func (tree *Tree) TraversePost(fn func(*Node), r2l bool) {
	if tree.Root != nil {
		tree.Root.TraversePost(fn, r2l)
	}
}

// breadth first

func (node *Node) TraverseBF(fn func(*Node), r2l bool) {
	var queue []*Node
	queue = append(queue, node)

	// TODO
}

// dump to dot <-- TODO maybe

func (node *Node) PrintTree(level int) {
	for i := 0; i < level; i++ {
		fmt.Print("-")
	}
	fmt.Println(node.Data)

	for i := 0; i < len(node.Children); i++ {
		if node.Children[i] != nil {
			node.Children[i].PrintTree(level + 1)
		}
	}
}

func (tree *Tree) PrintTree() {
	if tree.Root != nil {
		tree.Root.PrintTree(0)
	}
}
