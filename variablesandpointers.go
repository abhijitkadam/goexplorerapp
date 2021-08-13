package main

import "fmt"

func VariablesAndPointer() {
	var name string = "Abhijit"
	fmt.Println(name)

	role := "Go Developer"

	fmt.Println(role)

	namePtr := &name
	fmt.Printf("%p %v %T \n", namePtr, *namePtr, namePtr)

	*namePtr = "Somnath"

	fmt.Printf("%p %v \n", namePtr, *namePtr)

	namePtr = &role
	fmt.Printf("%p %v \n", namePtr, *namePtr)

	//For revision on basic or elementary types like bool, int, int32, int64, float explore:
	//https://tour.golang.org/

	//All the values in go are intitialized with default values.
	//like 0 for ints and false for boolean and empty string "" for string
	// for pointer type default is nil
	//accesing value of nil pointer will result in crash. So null check should be done.
	var aPtr *string
	fmt.Println(*aPtr)
}
