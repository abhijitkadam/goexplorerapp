package main

type Node struct {
	Value int64
	Next  *Node
}

func NewNode(v int64) *Node {
	return &Node{Value: v}
}

func OpsOnNode() {
	node := NewNode(5)
	node.Next = NewNode(10)
}

func AddElement(rootNode *Node, value int64) {
	//Traverse the node and add the value at the end
}

func TraverseAndPrintNodeValues(rootNode *Node) {
	//Traverse and print each value
	//temp pointing to root. Then you traverse till Next pointer is not nil
}
