package datatypesdemo

import (
	"fmt"
	"strings"
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

func demoArrayDeclaration() {
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
	demoArrayDeclaration()
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
	// slice declaration
	// (!) don't mix []T{...} (slice declaration) with [...]T{...} (array declaration)
	sliceOfStrings := []string{"AA", "BB", "CC"}
	fmt.Println("sliceOfStrings =", sliceOfStrings)
	fmt.Println("len(sliceOfStrings) =", len(sliceOfStrings))

	// append one or more elements to slice
	sliceOfStrings = append(sliceOfStrings, "appended1", "appended2", "appended3")

	// converting array to slice
	// len() is defined for both arrays and slices
	arrayOfStrings := [...]string{"DD", "EE", "FF"}
	fmt.Println("arrayOfStrings =", arrayOfStrings)
	fmt.Println("len(arrayOfStrings) =", len(arrayOfStrings))

	sliceOfStrings2 := arrayOfStrings[:]
	fmt.Println("sliceOfStrings2 =", sliceOfStrings2)
	fmt.Println("len(sliceOfStrings2) =", len(sliceOfStrings2))

	integersSequence := makeRange(3, 11)
	for i, v := range integersSequence {
		fmt.Printf("integersSequence[%d] = %d\n", i, v)
	}

	dirPath := "/home/user/dev/go/src/github.com/UserName/app-demo/datavol/demo/os/dir1"

	segments := strings.Split(dirPath, "/")
	fmt.Println("segments =", segments)

	removedLastSegment := segments[:len(segments)-1]
	fmt.Println("removedLastSegment =", removedLastSegment)

	removedLast2Segments := segments[:len(segments)-2]
	fmt.Println("removedLast2Segment =", removedLast2Segments)

	last3Segments := segments[len(segments)-3:]
	fmt.Println("last3Segments =", last3Segments)
}

// Code to accompany the following talk:
// dotGo 2019 - Daniel Martí - Optimizing Go code without a blindfold
// https://www.youtube.com/watch?v=jiXnzkAzy30
// 1st implementation
func copyList1(in []string) []string {
	var out []string
	for _, s := range in {
		out = append(out, s)
	}
	return out
}

func copyList2(in []string) []string {
	out := make([]string, len(in))
	for i, s := range in {
		out[i] = s
	}
	return out
}

// ShowDemo func
func ShowDemo() {
	fmt.Printf("\n\ndatatypesdemo.ShowDemo()\n\n")
	demoTypeAssertion()
	sliceDemo()
	fmt.Printf("\n\n~datatypesdemo.ShowDemo()\n\n")
}
