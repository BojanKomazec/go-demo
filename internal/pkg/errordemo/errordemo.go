package errordemo

import (
	"errors"
	"fmt"
)

func errorDemo() {
	var err error
	println("err =", err)

	// panic: runtime error: invalid memory address or nil pointer dereference
	// [signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x4d208b]
	// println("err.Error() =", err.Error())

	err = errors.New("This is an error message")
	println("err.Error() =", err.Error())
}

// ShowDemo func
func ShowDemo() {
	fmt.Printf("\n\nerrordemo.ShowDemo()\n\n")
	errorDemo()
	fmt.Printf("\n\n~errordemo.ShowDemo()\n\n")
}
