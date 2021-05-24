package concurrency

import (
	"fmt"
	"sync"
)

// Sends values 0 through 9 into "in" channel and makes sure that they all come out in the correct order from the
// "out" channel.
//
// Example and implementation are from: https://medium.com/capital-one-tech/building-an-unbounded-channel-in-go-789e175cd2cd.
func ExampleInOut() {
	in, out := InOut()
	lastVal := -1
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		for v := range out {
			vi := v.(int)
			fmt.Println("Reading:", vi)
			if lastVal + 1 != vi {
				panic("sequence is out of order")
			}
			lastVal = vi
		}
		wg.Done()
		fmt.Println("Finished reading!")
	}()

	for i := 0; i < 10; i++ {
		fmt.Println("Writing:", i)
		in <- i
	}

	close(in)
	fmt.Println("Finished writing!")
	wg.Wait()

	if lastVal != 9 {
		panic("last value isn't 99")
	}
	// Unordered output:
	// Writing: 0
	// Writing: 1
	// Writing: 2
	// Writing: 3
	// Writing: 4
	// Writing: 5
	// Writing: 6
	// Writing: 7
	// Writing: 8
	// Writing: 9
	// Finished writing!
	// Reading: 0
	// Reading: 1
	// Reading: 2
	// Reading: 3
	// Reading: 4
	// Reading: 5
	// Reading: 6
	// Reading: 7
	// Reading: 8
	// Reading: 9
	// Finished reading!
}
