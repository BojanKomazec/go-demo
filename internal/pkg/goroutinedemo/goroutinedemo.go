package goroutinedemo

import (
	"fmt"
	"time"
)

func declareFunctionInsideFunction() {
	// anonymous function
	add := func(a, b int) int {
		return a + b
	}
	fmt.Println("add(1, 2) =", add(1, 2))

	// immediately invoked anonymous function
	fmt.Println("3 + 4 =", func(a, b int) int {
		return a + b
	}(3, 4))
}

// https://github.com/golang/go/wiki/CommonMistakes#using-goroutines-on-loop-iterator-variables
//
// Incorrect implementation output:
// goroutine #5
// goroutine #5
// goroutine #5
// goroutine #5
// goroutine #5
//
// Correct implementation (possible) output:
// goroutine #1
// goroutine #0
// goroutine #2
// goroutine #3
// goroutine #4
func demoClosuresCapturingLoopVariable() {
	// Incorrect implementation:
	for i := 0; i < 5; i++ {
		go func() {
			// warning "loop variable i captured by func literal (go-vet)" is reported on i
			// This is potentially unwanted behaviour as we want each goroutine to capture
			// a new value of loop variable but this might not be the case as Go closures
			// capture variables by reference. Loop might finish before goroutines start being
			// executed.
			fmt.Printf("goroutine #%d\n", i)
		}()
	}

	// Correct implementation:
	for i := 0; i < 5; i++ {
		go func(i int) {
			fmt.Printf("goroutine #%d\n", i)
		}(i)
	}
}

func demoLimitCurrentNumberOfGoroutines() {
	const maxNumOfParallelGoroutines = 4
	limiter := make(chan struct{}, maxNumOfParallelGoroutines)
	for i := 0; i < 15; i++ {
		limiter <- struct{}{}
		go func(i int) {
			fmt.Printf("Starting goroutine #%d\n", i)
			time.Sleep(5 * time.Second)
			fmt.Printf("Terminating goroutine #%d\n", i)
			<-limiter
		}(i)
	}
}

// ShowDemo func
func ShowDemo() {
	fmt.Printf("\n\ngoroutinedemo.ShowDemo()\n\n")
	declareFunctionInsideFunction()
	demoClosuresCapturingLoopVariable()
	demoLimitCurrentNumberOfGoroutines()
	fmt.Printf("\n\n~goroutinedemo.ShowDemo()\n\n")
}
