package fmtdemo

import (
	"fmt"
)

// ScanDemo function
func ScanDemo() {
	slc := make([]int, 0, 100)
	if n, err := fmt.Scanln(&slc); err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("n =", n)
	}
}

func ReadIntegersFromLine() {
	// slc := make([]int, 0, 10)
	var strInput string
	if n, err := fmt.Scanln(&strInput); err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("n =", n)
		fmt.Println("strInput =", strInput)
	}
}
