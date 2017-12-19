package treecontainer

import (
	"../queue"
	"../stack"
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
//	"sync"
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

// non-recursive
func (node *Node) TraversePreNR(fn func(*Node), r2l bool) {
// 	s := new(stack.Stack)
	s := stack.New()
	s.Push(node)

	for !s.IsEmpty() {
		curr := s.Pop().(*Node)

		fn(curr)

		curr.IterateChildren(func(n *Node) { s.Push(n) }, !r2l)
	}
}

// depth first pre-order search

func (node *Node) SearchPre(test func(*Node) *Node, r2l bool) *Node {
	if test(node) != nil {
		return node
	}
	return node.SearchChildren(func(n *Node) *Node { return n.SearchPre(test, r2l) }, r2l)
}

// non-recursive
func (node *Node) SearchPreNR(test func(*Node) *Node, r2l bool) *Node {
//	s := new(stack.Stack)
	s := stack.New()
	s.Push(node)

	for !s.IsEmpty() {
		curr := s.Pop().(*Node)

		if test(curr) != nil {
			return curr
		}
	
		curr.IterateChildren(func(n *Node) { s.Push(n) }, !r2l)
	}
	return nil
}

func (node *Node) SearchPreNRAsyncAux(test func(*Node) *Node, r2l bool, done *bool) *Node {
	s := stack.New()
	s.Push(node)

	for !s.IsEmpty() {
		curr := s.Pop().(*Node)

		if test(curr) != nil {
			// *done = true
			return curr
		}
		
		if *done == true {
			return nil
		}
	
		curr.IterateChildren(func(n *Node) { s.Push(n) }, !r2l)
	}
	return nil
}

func (node *Node) SearchPreNRAsync(test func(*Node) *Node, r2l bool) *Node {
	if test(node) != nil {
		return node
	}

	// var wg sync.WaitGroup

	var children []*Node

	node.IterateChildren(func(n *Node) { children = append(children, n) }, r2l)

	// done := make(chan bool)

	// var doneChans []chan bool

	// for i := 0; i < len(children); i++ {
	// 	doneChans[i] = make(chan bool)
	// }

	result := make(chan *Node)
	done := false
	doneN := 0
	
	for i := 0; i < len(children); i++ {
		// wg.Add(1)
		go func(n *Node) {
			// defer func() {
			// 	if r := recover(); r != nil {
			// 		return
			// 	}
			// }()
			// defer wg.Done();
			defer func () { doneN++ }()
			res := n.SearchPreNRAsyncAux(test, r2l, &done)
			if res != nil {
				// fmt.Println("write to result")
				done = true
				result <- res
				close(result)
			} else {
				result <- nil
			}
		}(children[i])
	}
	
	// wg.Wait()
	for retval := range result {
		if retval != nil {
			return retval
		}
		if doneN == len(children) {
			close(result)
			return nil
		}
	}
	return nil
}

// depth first post-order traversal

func (node *Node) TraversePost(fn func(*Node), r2l bool) {
	node.IterateChildren(func(n *Node) { n.TraversePost(fn, r2l) }, r2l)
	fn(node)
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

// tree aliases

func (tree *Tree) PrintTree() {
	if tree.Root != nil {
		tree.Root.PrintTree(0)
	}
}

// reflection lol (it actually works)
func (tree *Tree) traverseTreeAlias(mName string, fn func(*Node), r2l bool) {
	in := make([]reflect.Value, 2)
	in[0] = reflect.ValueOf(fn)
	in[1] = reflect.ValueOf(r2l)
	if tree.Root != nil {
		reflect.ValueOf(tree.Root).MethodByName(mName).Call(in)
	}
}

// reflection lol (it actually works)
func (tree *Tree) searchTreeAlias(mName string, test func(*Node) *Node, r2l bool) *Node {
	in := make([]reflect.Value, 2)
	in[0] = reflect.ValueOf(test)
	in[1] = reflect.ValueOf(r2l)
	if tree.Root != nil {
		return reflect.ValueOf(tree.Root).MethodByName(mName).Call(in)[0].Interface().(*Node)
	}
	return nil
}

func (tree *Tree) TraversePre(fn func(*Node), r2l bool) {
	// tree.traverseTreeAlias("TraversePre", fn, r2l)
	if tree.Root != nil {
		tree.Root.TraversePre(fn, r2l)
	}
}

func (tree *Tree) TraversePreNR(fn func(*Node), r2l bool) {
	if tree.Root != nil {
		tree.Root.TraversePreNR(fn, r2l)
	}
}

func (tree *Tree) SearchPre(test func(*Node) *Node, r2l bool) *Node {
	// return tree.searchTreeAlias("SearchPre", test, r2l)
	if tree.Root != nil {
		return tree.Root.SearchPre(test, r2l)
	}
	return nil
}

func (tree *Tree) SearchPreNR(test func(*Node) *Node, r2l bool) *Node {
	if tree.Root != nil {
		return tree.Root.SearchPreNR(test, r2l)
	}
	return nil
}

func (tree *Tree) SearchPreNRAsync(test func(*Node) *Node, r2l bool) *Node {
	if tree.Root != nil {
		return tree.Root.SearchPreNRAsync(test, r2l)
	}
	return nil
}

func (tree *Tree) TraversePost(fn func(*Node), r2l bool) {
	if tree.Root != nil {
		tree.Root.TraversePost(fn, r2l)
	}
}

func (tree *Tree) SearchPost(test func(*Node) *Node, r2l bool) *Node {
	if tree.Root != nil {
		return tree.Root.SearchPost(test, r2l)
	}
	return nil
}

func (tree *Tree) TraverseBF(fn func(*Node), r2l bool) {
	if tree.Root != nil {
		tree.Root.TraverseBF(fn, r2l)
	}
}

func (tree *Tree) SearchBF(fn func(*Node) *Node, r2l bool) *Node {
	if tree.Root != nil {
		return tree.Root.SearchBF(fn, r2l)
	}
	return nil
}
