package main

import "fmt"

func main() {
	OpsOnNode()

	rootNode := NewNode(5)
	AddElement(rootNode, 10)
	AddElement(rootNode, 20)
	AddElement(rootNode, 25)

	//Should print 5 -> 10 -> 20 -> 25
	TraverseAndPrintNodeValues(rootNode)

}

func GoLangForLoopForms() {

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	//While loop like behavior
	count := 0
	for count < 10 {
		count += 1
		fmt.Println(count)
	}

	count2 := 0
	//While loop like behavior
	for {
		count2 += 1
		if count2 > 10 {
			break
		}
		fmt.Println(count2)
	}

}
