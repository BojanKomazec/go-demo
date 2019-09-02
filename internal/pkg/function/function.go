package function

import (
	"errors"
	"fmt"
)

// VariadicFunction - variadic function example
func variadicFunction(args ...interface{}) {
	fmt.Printf("\nvariadicFunction()\n")

	for index, value := range args {
		fmt.Println("index =", index, "; value =", value)
	}
}

type myFunc func()

func foo(fn myFunc) (string, error) {
	fn()
	return "value returned from foo()", errors.New("error returned from foo()")
}

func testVarScope() {
	fmt.Printf("\ntestVarScope()\n")
	var (
		s   string
		err error
	)

	s = "test1"
	err = errors.New("This is an eror message")

	s, err = foo(func() {
		println("s =", s, "err =", err.Error())
	})

	println("s =", s, "err =", err.Error())
}

// ShowDemo func
func ShowDemo() {
	fmt.Printf("\n\nfunction.ShowDemo()\n\n")
	variadicFunction(1, 'a', true, "bcdef")
	testVarScope()
	fmt.Printf("\n\n~function.ShowDemo()\n\n")
}
