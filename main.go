package main

import (
	"fmt"
	"time"
)

func main() {
	//GoKeyWordDemo()

	for i := 0; i < 10000; i++ {
		go func(count int) {
			fmt.Println("Running concurrently : ", count)
		}(i)
	}

	time.Sleep(time.Second * 2)
}
