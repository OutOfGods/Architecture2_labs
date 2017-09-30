package treecontainer

import (
	"bufio"
	"fmt"
	"../queue"
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
			tree.Root = &Node{Data: data, Parent: nil}
			tree.CurrNodePtr = tree.Root
		} else {
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

func (node *Node) SearchChildren(test func(*Node) *Node, r2l bool) *Node {
	if r2l {
		for i := len(node.Children) - 1; i >= 0; i-- {
			res := test(node.Children[i])
			if res != nil {
				return res
			}
		}
	} else {
		for i := 0; i < len(node.Children); i++ {
			res := test(node.Children[i])
			if res != nil {
				return res
			}
		}
	}
	return nil
}

// depth first pre-order traversal

func (node *Node) TraversePre(fn func(*Node), r2l bool) {
	fn(node)
	node.IterateChildren(func(n *Node) { n.TraversePre(fn, r2l) }, r2l)
}

func (tree *Tree) TraversePre(fn func(*Node), r2l bool) {
	if tree.Root != nil {
		tree.Root.TraversePre(fn, r2l)
	}
}

// depth first pre-order search

func (node *Node) SearchPre(test func(*Node) *Node, r2l bool) *Node {
	if test(node) != nil {
		return node
	}
	return node.SearchChildren(func(n *Node) *Node { return n.SearchPre(test, r2l) }, r2l)
}

func (tree *Tree) SearchPre(test func(*Node) *Node, r2l bool) *Node {
	if tree.Root != nil {
		return tree.Root.SearchPre(test, r2l)
	}
	return nil
}

// depth first post-order traversal

func (node *Node) TraversePost(fn func(*Node), r2l bool) {
	node.IterateChildren(func(n *Node) { n.TraversePost(fn, r2l) }, r2l)
	fn(node)
}

func (tree *Tree) TraversePost(fn func(*Node), r2l bool) {
	if tree.Root != nil {
		tree.Root.TraversePost(fn, r2l)
	}
}

// depth first post-order search

func (node *Node) SearchPost(test func(*Node) *Node, r2l bool) *Node {
	res := node.SearchChildren(func(n *Node) *Node { return n.SearchPost(test, r2l) }, r2l)
	if res != nil {
		return res
	} else if test(node) != nil {
		return node
	}
	return nil
}

func (tree *Tree) SearchPost(test func(*Node) *Node, r2l bool) *Node {
	if tree.Root != nil {
		return tree.Root.SearchPost(test, r2l)
	}
	return nil
}

// breadth first traversal

func (node *Node) TraverseBF(fn func(*Node), r2l bool) {
	q := queue.NewQueue(5)
	q.Push(node)

	for !q.IsEmpty() {
		curr := q.Pop().(*Node)

		fn(curr)

		curr.IterateChildren(func(n *Node) { q.Push(n) }, r2l)
	}
}

func (tree *Tree) TraverseBF(fn func(*Node), r2l bool) {
	if tree.Root != nil {
		tree.Root.TraverseBF(fn, r2l)
	}
}

// breadth first search

func (node *Node) SearchBF(test func(*Node) *Node, r2l bool) *Node {
	q := queue.NewQueue(5)
	q.Push(node)

	for !q.IsEmpty() {
		curr := q.Pop().(*Node)

		res := test(curr)
		if res != nil {
			return res
		}

		curr.IterateChildren(func(n *Node) { q.Push(n) }, r2l)
	}
	return nil
}

func (tree *Tree) SearchBF(fn func(*Node) *Node, r2l bool) *Node {
	if tree.Root != nil {
		return tree.Root.SearchBF(fn, r2l)
	}
	return nil
}

// dump to dot <-- TODO maybe

// print tree

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
