package types

import (
	"fmt"
)

// IotaDemo ; in the current implementation iota values start from 1 but this is not guaranteed and can change in future
// This function also shows how to declare enums in Go.
func IotaDemo() {
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

func EnumDemo() {
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
