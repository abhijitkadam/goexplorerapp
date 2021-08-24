package main

import (
	"fmt"
	"sync"
)

func PipelineDemo() {
	//PipelineDemoSingle()
	PipelineDemoWithMergeMultiple()
}

func PipelineDemoSingle() {

	gCh := Generator(2, 4, 6, 8)
	dCh := Doubler(gCh)

	for i := range dCh {
		fmt.Println(i)
	}

}

func PipelineDemoWithMerge() {

	gCh := Generator(2, 4, 6, 8)
	d1Ch := Doubler(gCh)
	d2Ch := Doubler(gCh)

	dCh := merge_two(d1Ch, d2Ch)

	for i := range dCh {
		fmt.Println(i)
	}

}

func PipelineDemoWithMergeMultiple() {

	gCh := Generator(2, 4, 6, 8)
	d1Ch := Doubler(gCh)
	d2Ch := Doubler(gCh)
	d3Ch := Doubler(gCh)

	dCh := merge(d1Ch, d2Ch, d3Ch)

	for i := range dCh {
		fmt.Println(i)
	}

}

func Generator(vals ...int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for _, v := range vals {
			ch <- v
		}
	}()

	return ch
}

func Doubler(inch <-chan int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)

		for v := range inch {
			ch <- v * 2
		}

	}()

	return ch
}

func merge_two(left, right <-chan int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)

		for left != nil || right != nil {
			select {
			case v, ok := <-left:
				if ok {
					ch <- v
				} else {
					left = nil
				}

			case v, ok := <-right:
				if ok {
					ch <- v
				} else {
					right = nil
				}
			}
		}

	}()

	return ch
}

func merge(chans ...<-chan int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)

		var wg sync.WaitGroup
		//wg.Add(len(chans))
		for _, c := range chans {
			wg.Add(1)
			go func(inCh <-chan int) {
				defer wg.Done()
				for v := range inCh {
					ch <- v
				}
			}(c)
		}

		wg.Wait()
	}()

	return ch
}
