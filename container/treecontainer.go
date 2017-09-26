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

func (tree *Tree) printTree(root *Node, level int) {
	for i := 0; i < level; i++ {
		fmt.Print("-")
	}
	fmt.Println(root.Data)
	for i := 0; i < len(root.Children); i++ {
		if root.Children[i] != nil {
			tree.printTree(root.Children[i], level+1)
		}
	}
}

func (tree *Tree) PrintTree() {
	if tree.Root != nil {
		tree.printTree(tree.Root, 0)
	}
}
