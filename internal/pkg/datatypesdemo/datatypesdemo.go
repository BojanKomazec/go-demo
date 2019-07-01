package datatypesdemo

import (
	"fmt"
)

// https://golang.org/ref/spec#Type_assertions
func demoTypeAssertion() {
	var a interface{} = "this is a string"
	s := a.(string)
	fmt.Println("s =", s)
	// output: s = this is a string

	// If the type assertion is false, a run-time panic occurs:
	// i := a.(int)
	// panic: interface conversion: interface {} is string, not int
	// fmt.Println("i =", i)
}

func demoDeclaration() {
	arr1 := [3]string{"AA", "BB", "CC"}
	fmt.Println("arr1 =", arr1)
	fmt.Println("len(arr1) =", len(arr1))

	// https://programming.guide/go/three-dots-ellipsis.html:
	// In an array literal, the ... notation specifies a length equal to the number of elements in the literal.
	arr2 := [...]string{"DD", "EE", "FF", "GG"}
	fmt.Println("arr2 =", arr2)
	fmt.Println("len(arr2) =", len(arr2))
}

func arrayDemo() {
	demoDeclaration()
}

// makeRange creates an increasing sequence of integers
// Arguments:
//    min - first number in the sequence
//    max - last  number in the sequence
// Returns:
//    a slice of integers
func makeRange(min, max int) []int {
	slice := make([]int, max-min+1)
	for i := range slice {
		slice[i] = min + i
	}
	return slice
}

func sliceDemo() {
	integersSequence := makeRange(3, 11)
	for i, v := range integersSequence {
		fmt.Printf("integersSequence[%d] = %d\n", i, v)
	}
}

// ShowDemo func
func ShowDemo() {
	demoTypeAssertion()
	sliceDemo()
}
