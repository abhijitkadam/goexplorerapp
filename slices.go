package main

import "fmt"

func GoArrayAndSlices() {
	nums := []int32{5, 10, 4, 33, 56}
	fmt.Println(len(nums))
	fmt.Println(nums)

	//nums[low : high]
	//including low and *exluding high
	numsRange := nums[0:3]
	fmt.Println(numsRange)

	numsRange2 := nums[1:4]
	fmt.Println(numsRange2)
}

func GoSlices() {

	//slice is managing the underlying array
	//below is slice with length 0 and capacity 100.
	//So we can only safely access elements equal to length of slice.
	//Capacity reserves contagious memory locations for future additions (append) to underlying array via slice.
	//length should not be be grater than capacity for the obivious reason
	numSlice := make([]int32, 0, 10)

	fmt.Println(numSlice)

	//Append will all elements to slice (i.e. the underlying managed array).
	//While appending what if the length of equal to capacity already?
	numSlice = append(numSlice, 2, 5, 6, 7)

	fmt.Println(len(numSlice))

	//If there is no space i.e the capacity has exhausted then what append will do?
	//Append will allocate higher capacity and then add the extra elements
	//We no contagious memory blocks are avaialable then it will move the entire array from begining to new memory location.
	numSlice = append(numSlice, 2, 5, 6, 7, 4, 3, 2, 12, 6, 89)
	fmt.Println(len(numSlice))
}

func GoSlicesRange() {
	numSlice := make([]int32, 0, 10)
	numSlice = append(numSlice, 2, 5, 6, 7, 4, 3, 2, 12, 6, 89)

	//value semantics
	//v is a copy of object (in case it is slice of some struct)
	for i, v := range numSlice {
		fmt.Println(i, " : ", v)
	}

	fmt.Println("---------------")

	//If we want to prevent a copy or use other way then below 2 approaches are for reference sematics
	for i := 0; i < len(numSlice); i++ {
		fmt.Println(i, " : ", numSlice[i])
	}

	//you can mutate also if you use reference
	for i, _ := range numSlice {
		numSlice[i] = 9
	}

	fmt.Println(numSlice)

	fmt.Println("---------------")

	//concat slices
	another := make([]int32, 0, 100)
	another = append(another, 7)
	another = append(another, numSlice...)
	fmt.Println(another)

}
