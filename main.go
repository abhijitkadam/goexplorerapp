package main

import (
	"fmt"
	"sync"
)

func main() {
	//GoKeyWordDemo()
	ConcurrenyWithWaitGroupExampleWithDefer()
}

func ConcurrenyWithWaitGroupExample() {
	var wg sync.WaitGroup

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(count int) {
			fmt.Println("Running concurrently : ", count)
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println("Completed!")
}

func ConcurrenyWithWaitGroupExampleWithDefer() {
	var wg sync.WaitGroup

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(count int) {
			//With defer we made sure that wg.Done() would always be called when this func returns
			//So we are defering execution of wg.Done untill the end i.e. when function returns
			defer wg.Done()
			fmt.Println("Running concurrently : ", count)
		}(i)
	}

	wg.Wait()
	fmt.Println("Completed!")
}
