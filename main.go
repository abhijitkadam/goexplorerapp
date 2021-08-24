package main

import (
	"fmt"
	"sync"
	"time"
)

//https://www.ardanlabs.com/blog/2014/02/the-nature-of-channels-in-go.html
func main() {
	//ChannelBasics()
	//ChannelWorkers()
	SelectDemo()
}

//Channels are unbuffered & buffered
//By default send and receive operations on channel are blocking
//If you have to close a channel it will be the sender or generator that will be closing the channel.
//Sending receiving on a closed channel is bad
//Sending receving on a nil channel is a blocking operation. var ch chan int (would make() is nil channel).
//or you can make a channel nil by: ch = nil
func ChannelBasics() {
	//unbuffered channel
	ch := make(chan int64)
	//buffered channel
	//ch := make(chan int64, 3)

	//Sender sending 3 values on a channel
	go func() {
		defer close(ch)
		ch <- 5
		ch <- 10
		ch <- 15
	}()

	// out1 := <-ch
	// fmt.Println(out1)
	// out2 := <-ch
	// fmt.Println(out2)
	// out3 := <-ch
	// fmt.Println(out3)

	//Receiver reading values from the channel
	time.Sleep(time.Second * 5)
	for out := range ch {
		fmt.Println(out)
	}

}

//Channels are syncronised
//Multiple sender can send into a channel
//Multiple receiver can receive from a channel
//Sending receiving on a close channel is bad
func ChannelWorkers() {

	ch := make(chan int64)

	//Work generator
	go func() {
		defer close(ch)
		ch <- 5
		ch <- 10
		ch <- 15
		ch <- 15
		ch <- 17
		ch <- 18
		ch <- 19
		ch <- 20
	}()

	//Lets create some workers to get the work items from channel ch.
	//The concurrent workers (go routines) will receive data (work item num) from ch
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			for out := range ch {
				fmt.Println("Worker number: ", num, " working on item no : ", out)
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("Done working on the go work items")

}
