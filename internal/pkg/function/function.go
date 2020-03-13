package function

import (
	"errors"
	"fmt"
)

// Go does not have optional parameters nor does it support method overloading.

// Regarding optional parameters, we can use default values of the type in some cases.
// The caller has to use default values explicitly
func defaultParameterValueDemo(stringArg string, intArg int) {
	if stringArg == "" {
		// default value used
	}

	if intArg == 0 {
		// default value used
	}
}

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

func execute(f func()) {
	f()
}

type person struct {
	Name string
	Age  int
}

func printMembers(p person) {
	fmt.Printf("printMembers(): Person Name = %s, Age = %d\n", p.Name, p.Age)
}

func captureDemo() {
	p := person{
		Name: "Bojan",
		Age:  41,
	}

	execute(func() {
		fmt.Printf("Person Name = %s, Age = %d\n", p.Name, p.Age)

		p.Age = 42
		execute(func() {
			fmt.Printf("Person Name = %s, Age = %d\n", p.Name, p.Age)

			p.Age = 43
			printMembers(p)
		})

		p.Age = 44
	})
}

func internalFunctionDemo() {

	// It is not possible to declare internal function by stating its name:
	// error: expected expression
	// func foo(){}

	// ...but it is possible to assign anonymous function to a variable and call it via that variable:
	foo := func() {
		fmt.Println("foo()")
	}

	foo()
}

// ShowDemo func
func ShowDemo() {
	fmt.Printf("\n\nfunction.ShowDemo()\n\n")
	// variadicFunction(1, 'a', true, "bcdef")
	// testVarScope()
	// captureDemo()
	internalFunctionDemo()
	fmt.Printf("\n\n~function.ShowDemo()\n\n")
}
