package types

import (
	"fmt"
)

// IotaDemo ; in the current implementation iota values start from 1 but this is not guaranteed and can change in future
// This function also shows how to declare enums in Go.
func iotaDemo() {
	fmt.Println("\nIotaDemo()")

	type Grades int
	const (
		A Grades = iota
		B
		C
		D
		E
		F
	)

	fmt.Println("A =", A) // output: A = 0
	fmt.Println("C =", C) // output: C = 2

	var grade Grades
	grade = 123
	fmt.Println("grade =", grade) // Output: grade = 123 // We don't want this to happen but seems that's unavoidable.

	type Days int
	const (
		Monday Days = iota + 1
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		Sunday
	)

	fmt.Println("Monday =", Monday) // 1
	fmt.Println("Sunday =", Sunday) // 7

	// printing stringified enum version - use %v
	fmt.Printf("Wed = %v\n", Wednesday) // 7

}

func enumDemo() {
	fmt.Println("\nEnumDemo()")

	type Days string
	const (
		Monday    Days = "Monday"
		Tuesday        = "Tuesday"
		Wednesday      = "Wednesday"
		Thursday       = "Thursday"
		Friday         = "Friday"
		Saturday       = "Saturday"
		Sunday         = "Sunday"
	)

	var day Days
	fmt.Println("day =", day)

	day = Monday
	fmt.Println("day =", day)

	// error: Monday1 is undefined
	// day = Monday1

	day = "Monday1"
	fmt.Println("day =", day) // Output: day = Monday1 // We don't want this to be possible

	var strDay string

	// cast (type conversion) of enumerator to string
	strDay = string(Monday)
	fmt.Println("strDay =", strDay)
	strDay = Tuesday
	fmt.Println("strDay =", strDay)

	// var day3 Days
	// day3 = Wednesday
	// cannot use day3 (type Days) as type string in assignment
	// strDay = day3
	// fmt.Println("strDay =", strDay)

}

// The expression T(x) converts the value x to the type T.
// string(x) evaluates to a string containing the UTF-8 encoding of the value of x (NOT to the string representation of the integer x.)
// This makes sense as strings in Go are UTF-8 encoded.
// So don't use string(x) for conversion of integer to string!
// https://golang.org/doc/go1.15#vet
func int2stringConversionDemo() {
	fmt.Printf("string(1) is %s\n", string(1))       // string(1) is
	fmt.Printf("string(9786) is %s\n", string(9786)) // string(9786) is ☺

	// this is a proper way to convert an integer to a string
	fmt.Printf(`fmt.Sprintf("%%d", 9786) is %s`+"\n", fmt.Sprintf("%d", 9786)) // fmt.Sprintf("%d", 9786) is 9786
}

func ShowDemo() {
	enumDemo()
	iotaDemo()
	int2stringConversionDemo()
}
