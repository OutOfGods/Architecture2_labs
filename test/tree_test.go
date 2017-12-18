package test

import (
	"../treecontainer"
	"fmt"
	"testing"
)

func SearchForTest(n *treecontainer.Node) *treecontainer.Node {
	if n.Data == "11332" {
		return n
	}
	return nil
}

func TraverseForTest(n *treecontainer.Node) {
	if n.Data == "11332" {
	}
}

func BenchmarkNewTree(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var treePtr *treecontainer.Tree = &treecontainer.Tree{}
		treePtr.NewTreeFromFile("./test_input_1.txt")
	}
}

func BenchmarkTraversePrel2r(b *testing.B) {
	var treePtr *treecontainer.Tree = &treecontainer.Tree{}
	treePtr.NewTreeFromFile("./generated_test1.txt")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		treePtr.TraversePre(TraverseForTest, false)
	}
}

func BenchmarkTraversePrel2rNR(b *testing.B) {
	var treePtr *treecontainer.Tree = &treecontainer.Tree{}
	treePtr.NewTreeFromFile("./generated_test1.txt")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		treePtr.TraversePreNR(TraverseForTest, false)
	}
}

func BenchmarkSearchPrel2r(b *testing.B) {
	var treePtr *treecontainer.Tree = &treecontainer.Tree{}
	treePtr.NewTreeFromFile("./generated_test1.txt")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		treePtr.SearchPre(SearchForTest, false)
	}
}

func BenchmarkSearchPrel2rNR(b *testing.B) {
	var treePtr *treecontainer.Tree = &treecontainer.Tree{}
	treePtr.NewTreeFromFile("./generated_test1.txt")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		treePtr.SearchPreNR(SearchForTest, false)
	}
}

func BenchmarkSearchPrel2rNRAsync(b *testing.B) {
	var treePtr *treecontainer.Tree = &treecontainer.Tree{}
	treePtr.NewTreeFromFile("./generated_test1.txt")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		treePtr.SearchPreNRAsync(SearchForTest, false)
	}
}

func BenchmarkSearchPostl2rNR(b *testing.B) {
	var treePtr *treecontainer.Tree = &treecontainer.Tree{}
	treePtr.NewTreeFromFile("./test_input_1.txt")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		treePtr.SearchBF(SearchForTest, false)
	}
}

func BenchmarkSearchPostl2r(b *testing.B) {
	var treePtr *treecontainer.Tree = &treecontainer.Tree{}
	treePtr.NewTreeFromFile("./test_input_1.txt")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		treePtr.SearchPost(SearchForTest, false)
	}
}

func ExampleTraversePrel2r() {
	var treePtr *treecontainer.Tree = &treecontainer.Tree{}
	treePtr.NewTreeFromFile("./test_input_1.txt")
	treePtr.TraversePre(func(n *treecontainer.Node) { fmt.Println(n.Data) }, false)
	// Output:
	//0.0
	//1.1
	//2.1
	//3.1
	//3.2
	//2.2
	//2.3
	//3.3
	//3.4
	//1.2
	//2.4
	//3.5
	//2.5
	//3.6
	//3.7
	//1.3
	//2.6
	//3.8
	//1.4
	//2.7
	//3.8
	//2.7
	//3.9
	//3.10
	//3.11
	//2.8
	//3.12
	//3.13
}

func ExampleTraversePrer2l() {
	var treePtr *treecontainer.Tree = &treecontainer.Tree{}
	treePtr.NewTreeFromFile("./test_input_1.txt")
	treePtr.TraversePre(func(n *treecontainer.Node) { fmt.Println(n.Data) }, true)
	// Output:
	//0.0
	//1.4
	//2.8
	//3.13
	//3.12
	//2.7
	//3.11
	//3.10
	//3.9
	//2.7
	//3.8
	//1.3
	//2.6
	//3.8
	//1.2
	//2.5
	//3.7
	//3.6
	//2.4
	//3.5
	//1.1
	//2.3
	//3.4
	//3.3
	//2.2
	//2.1
	//3.2
	//3.1
}

func ExampleTraversePostl2r() {
	var treePtr *treecontainer.Tree = &treecontainer.Tree{}
	treePtr.NewTreeFromFile("./test_input_2.txt")
	treePtr.TraversePost(func(n *treecontainer.Node) { fmt.Println(n.Data) }, false)
	// Output:
	//Dima
	//Stepan
	//Kirill
	//Seryi
	//Nadya
	//Vika
	//Anton
	//Ivan
	//Natasha
	//Sveta
	//Anzhela
	//Vadim
	//Konstantin
	//Katya
	//Timofey
	//Bohdan
	//Neshta
	//Aleksandr
	//Anatolii
	//Svyatoslav
	//Kazimir
	//John
	//Vladimir
	//Maksim
	//Vasya
	//Mishel
	//Danil
	//Andrey
}

func ExampleTraversePostr2l() {
	var treePtr *treecontainer.Tree = &treecontainer.Tree{}
	treePtr.NewTreeFromFile("./test_input_2.txt")
	treePtr.TraversePost(func(n *treecontainer.Node) { fmt.Println(n.Data) }, true)
	// Output:
	//Vasya
	//Maksim
	//Mishel
	//John
	//Kazimir
	//Svyatoslav
	//Vladimir
	//Aleksandr
	//Anatolii
	//Danil
	//Timofey
	//Bohdan
	//Neshta
	//Vadim
	//Anzhela
	//Konstantin
	//Natasha
	//Sveta
	//Katya
	//Vika
	//Nadya
	//Anton
	//Seryi
	//Stepan
	//Dima
	//Kirill
	//Ivan
	//Andrey
}

func ExampleTraverseBFl2r() {
	var treePtr *treecontainer.Tree = &treecontainer.Tree{}
	treePtr.NewTreeFromFile("./test_input_3.txt")
	treePtr.TraverseBF(func(n *treecontainer.Node) { fmt.Println(n.Data) }, false)
	// Output:
	//1
	//11
	//12
	//111
	//112
	//113
	//114
	//1111
	//1112
	//1113
	//1114
	//1131
	//1132
	//1133
	//1134
	//1135
	//11331
	//11332
	//11341
	//11342
	//11343
}

func ExampleTraverseBFr2l() {
	var treePtr *treecontainer.Tree = &treecontainer.Tree{}
	treePtr.NewTreeFromFile("./test_input_3.txt")
	treePtr.TraverseBF(func(n *treecontainer.Node) { fmt.Println(n.Data) }, true)
	// Output:
	//1
	//12
	//11
	//114
	//113
	//112
	//111
	//1135
	//1134
	//1133
	//1132
	//1131
	//1114
	//1113
	//1112
	//1111
	//11343
	//11342
	//11341
	//11332
	//11331
}

func ExampleNewTree() {
	var t *treecontainer.Tree = &treecontainer.Tree{}
	t.NewTreeFromFile("./test_input_2.txt")
	t.PrintTree()
	// Output:
	//Andrey
	//-Ivan
	//--Kirill
	//---Dima
	//---Stepan
	//--Seryi
	//--Anton
	//---Nadya
	//---Vika
	//-Katya
	//--Sveta
	//---Natasha
	//--Konstantin
	//---Anzhela
	//---Vadim
	//-Neshta
	//--Bohdan
	//---Timofey
	//-Danil
	//--Anatolii
	//---Aleksandr
	//--Vladimir
	//---Svyatoslav
	//---Kazimir
	//---John
	//--Mishel
	//---Maksim
	//---Vasya
}
