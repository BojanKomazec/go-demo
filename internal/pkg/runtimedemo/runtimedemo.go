package runtimedemo

import (
	"fmt"
	"runtime"
	"time"
)

func cpuTest() {
	logicalCPUMaxCount := runtime.NumCPU()
	fmt.Println("Number of logical CPUs on the local machine:", logicalCPUMaxCount)
}

// Sinnce v1.5 Go makes the default value of GOMAXPROCS the same as the number of CPUs on the machine
// If running on machine with INTEL® CORE™ i5-7300HQ PROCESSOR the logicalCPUMaxCount will be 4 as it
// has 4 cores and no support for hyper-threading.
func setMaxCPUNumTest() {
	logicalCPUMaxCount := runtime.NumCPU()
	prevMaxNumSimultaneousCPUs := runtime.GOMAXPROCS(logicalCPUMaxCount)
	fmt.Println("Previous number of logical CPUs that can run simultaneously:", prevMaxNumSimultaneousCPUs) // e.g. 4
	currMaxNumSimultaneousCPUs := runtime.GOMAXPROCS(0)
	fmt.Println("Current number of logical CPUs that can run simultaneously:", currMaxNumSimultaneousCPUs) // 4 (again)
}

// GoRoutineCountBackgroundMonitor func
func GoRoutineCountBackgroundMonitor() {
	for {
		fmt.Printf("%s Current number of goroutines: %d\n", time.Now().String(), runtime.NumGoroutine())
		time.Sleep(time.Second * 1)
	}
}

// ShowDemo func
func ShowDemo() {
	cpuTest()
	setMaxCPUNumTest()
}
