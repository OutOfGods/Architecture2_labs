package container

func TestNewTreeFromFile(file_name string) {
	var treePtr *Tree = &Tree{}
	treePtr.NewTreeFromFile(file_name)
	treePtr.PrintTree()
}
