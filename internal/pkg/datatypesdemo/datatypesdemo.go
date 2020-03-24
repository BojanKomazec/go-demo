package datatypesdemo

import (
	"fmt"
	"log"
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

// This function returns values of different types, depending on some runtime condition.
func returnsEmptyInterface(returnString bool) interface{} {
	if returnString {
		return "test"
	}

	return 1
}

func testCallingFunctionWhichReturnsEmptyInterface() {
	s := returnsEmptyInterface(true)
	fmt.Println("returnsEmptyInterface() returned string ", s)
	n := returnsEmptyInterface(false)
	fmt.Println("returnsEmptyInterface() returned integer ", n)
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

// Function argument cannot be of type [...]T (array of any size).
// It has to have size specified: [N]T
func arrayOf3StringsContains(arr [3]string, s string) bool {
	for _, e := range arr {
		if e == s {
			return true
		}
	}

	return false
}

// We can use function which has slice as an argument.
func stringSliceContains(slice []string, s string) bool {
	for _, e := range slice {
		if e == s {
			return true
		}
	}

	return false
}

func searchElementInArrayDemo() {
	fmt.Println("searchElementInArrayDemo()")

	s := "four"
	fmt.Println("s =", s)

	// type of arr is [3]string
	arr := [...]string{"one", "two", "three"}
	fmt.Println("arr = ", arr)

	found := arrayOf3StringsContains(arr, s)

	if found {
		fmt.Println("String found in array")
	} else {
		fmt.Println("String was not found in array")
	}

	arr2 := [...]string{"one", "two", "three", "four"}
	fmt.Println("arr2 = ", arr2)

	// To pass an array of any size we need to convert it to a slice.
	found = stringSliceContains(arr2[:], s)

	if found {
		fmt.Println("String found in array")
	} else {
		fmt.Println("String was not found in array")
	}
}

type person struct {
	Name string
}

// This function will not make desired change actually.
// 'range' creates a (temporary) copy of the current element and all changes are performed on it.
func modifyPersons(persons []person) []person {

	for _, p := range persons {
		if p.Name == "Bob" {
			p.Name = "Bojan"
		}
	}

	log.Println("~modifyPersons(): persons =", persons)
	return persons
}

func modifyPersons2(persons []person) []person {

	for i := range persons {
		if persons[i].Name == "Bob" {
			persons[i].Name = "Bojan"
		}
	}

	log.Println("~modifyPersons2(): persons =", persons)
	return persons
}

func createSliceCopy(persons []person) []person {
	copyPersons := make([]person, len(persons))

	for i := range persons {
		copyPersons[i] = persons[i]
	}

	return copyPersons
}

// WARNING:  append will allocate excess memory -- unless this array is later filled to capacity by some further processing
// https://stackoverflow.com/questions/30182538/why-cant-i-duplicate-a-slice-with-copy
func createSliceCopy2(persons []person) []person {
	return append([]person{}, persons...)
}

func createSliceCopy3(persons []person) []person {
	personsCopy := make([]person, len(persons))
	copy(personsCopy, persons)
	return personsCopy
}

func demoPassingSliceToFunction() {
	persons := []person{
		{
			"Alice",
		},
		{
			"Bob",
		},
	}

	persons = modifyPersons(persons)
	log.Println("persons (after ineffective modification try) =", persons)

	copyPersons := createSliceCopy(persons)
	log.Println("copyPersons (should be the same as persons) =", copyPersons)
	copyPersons = modifyPersons2(copyPersons)
	log.Println("persons (after modification of copyPersons) =", persons)
	log.Println("copyPersons =", copyPersons)

	persons = modifyPersons2(persons)
	log.Println("persons (after real modification) =", persons)
	log.Println("copyPersons (after real modification of persons) =", copyPersons)
}

func arrayDemo() {
	demoArrayDeclaration()
	searchElementInArrayDemo()
	demoPassingSliceToFunction()
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
	// Declaration by using var keyword.
	// declared but not initialized slice has value nil ("nil-slice")
	var slice0 []string
	if slice0 == nil {
		fmt.Println("slice0 is nil: ", slice0)
	}
	fmt.Printf("len(slice0) = %d, cap(slice0) = %d, slice0 = %#v\n", len(slice0), cap(slice0), slice0)
	// Output:
	// slice0 is nil:  []
	// len(slice0) = 0, cap(slice0) = 0, slice0 = []string(nil)

	// It is possible to append elements to nil-slice
	slice0 = append(slice0, "firstInslice0")
	fmt.Printf("slice0 = %#v\n", slice0)
	// slice0 = []string{"firstInslice0"}

	// Declaration by using "short declaration operator" (:=) and make function
	slice01 := make([]string, 0)
	if slice01 == nil {
		fmt.Println("slice01 is nil: ", slice01)
	} else {
		fmt.Println("slice01 is NOT nil: ", slice01)
	}
	fmt.Printf("len(slice01) = %d, cap(slice01) = %d, slice01 = %#v\n", len(slice01), cap(slice01), slice01)
	// Output:
	// slice01 is NOT nil:  []
	// len(slice01) = 0, cap(slice01) = 0, slice01 = []string{}

	slice01 = append(slice01, "firstInslice01")
	fmt.Printf("slice01 = %#v\n", slice01)
	// slice0 = []string{"firstInslice01"}

	// Declaration by using "short declaration operator" (:=) and empty slice literal
	// slice literal => underlying array is allocated (even if it's a zero-length slice)
	slice02 := []string{}
	if slice02 == nil {
		fmt.Println("slice02 is nil: ", slice02)
	} else {
		fmt.Println("slice02 is NOT nil: ", slice02)
	}
	fmt.Printf("len(slice02) = %d, cap(slice02) = %d, slice02 = %#v\n", len(slice02), cap(slice02), slice02)
	// Output:
	// slice02 is NOT nil:  []
	// len(slice02) = 0, cap(slice02) = 0, slice02 = []string{}

	slice02 = append(slice02, "firstInslice02")
	fmt.Printf("slice02 = %#v\n", slice02)
	// slice02 = []string{"firstInslice02"}

	var slice1 []string
	slice1 = []string{""}
	fmt.Println("slice1 =", slice1)           // value printed: []
	fmt.Println("len(slice1) =", len(slice1)) // 1

	var slice2 []string
	// if we don't call make here or assigne []string{} before, no memory will be allocated to slice and indexing oparation will fail:
	// panic: runtime error: index out of range
	// slice2[0] = ""

	// we need to allocate memory before assignments per element:
	slice2 = make([]string, 1)
	slice2[0] = ""
	fmt.Println("slice2 =", slice2)           // value printed: []
	fmt.Println("len(slice2) =", len(slice2)) // 1
	if slice2[0] == "" {
		fmt.Println("slice2[0] is empty string:", slice2[0])
	}

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

// how to create an empty list
func nilSliceDemo() {
	// nil slice; both length and capacity are 0.
	var list []int
	fmt.Printf("list= %v, len= %d, cap=%d\n", list, len(list), cap(list))

	list = append(list, 1)
	fmt.Printf("list= %v, len= %d, cap=%d\n", list, len(list), cap(list))
	list = append(list, 2)
	fmt.Printf("list= %v, len= %d, cap=%d\n", list, len(list), cap(list))
	list = append(list, 3)
	fmt.Printf("list= %v, len= %d, cap=%d\n", list, len(list), cap(list))
	list = append(list, 4)
	fmt.Printf("list= %v, len= %d, cap=%d\n", list, len(list), cap(list))
}

func emptyInterfaceDemo() {

	sliceEmptyInterace := []interface{}{"one", "two", "three"}
	fmt.Println("sliceEmptyInterace =", sliceEmptyInterace)

	// Create string slice from empty interface slice
	sliceString := make([]string, len(sliceEmptyInterace))
	for i := range sliceEmptyInterace {
		sliceString[i] = sliceEmptyInterace[i].(string)
	}

	fmt.Println("sliceString =", sliceString)
}

// ShowDemo func
func ShowDemo() {
	fmt.Printf("\n\ndatatypesdemo.ShowDemo()\n\n")
	// arrayDemo()
	// demoTypeAssertion()
	// emptyInterfaceDemo()
	// nilSliceDemo()
	sliceDemo()
	// testCallingFunctionWhichReturnsEmptyInterface()
	fmt.Printf("\n\n~datatypesdemo.ShowDemo()\n\n")
}
