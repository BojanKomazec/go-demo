package randdemo

import (
	"fmt"
	"math/rand"
	"time"
)

// on each run of the app, the sequence of random numbers generated here would be the same
// (unless different seed is set; by deafult the same seed, 1, is used)
func demoGenerateRandomsWithDefaultSeed() {
	fmt.Println("demoGenerateRandomsWithDefaultSeed()")
	for range [10]int{} {
		fmt.Println(rand.Intn(100))
	}
	fmt.Println("~demoGenerateRandomsWithDefaultSeed()")
}

// on each run of the app this sequence will be different as we're setting a different seed each time
func demoGenerateRandomsWithCustomSeed() {
	fmt.Println("demoGenerateRandomsWithCustomSeed()")
	seed := time.Now().UnixNano()
	rand.Seed(seed)
	for range [10]int{} {
		fmt.Println(rand.Intn(100))
	}
	fmt.Println("~demoGenerateRandomsWithCustomSeed()")
}

// ShowDemo func
func ShowDemo() {
	fmt.Println()
	fmt.Println("randdemo.ShowDemo()")
	demoGenerateRandomsWithDefaultSeed()
	demoGenerateRandomsWithCustomSeed()
	fmt.Println("~randdemo.ShowDemo()")
}
