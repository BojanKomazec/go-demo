package mapdemo

import (
	"fmt"
)

type test struct {
	Value1 string
	Value2 int
}

func populateMap(testMap map[string]*test) {
	//
	// Adding element members one by one
	//
	testMap["a"] = &test{} // an empty struct has to be added first
	testMap["a"].Value1 = "A"
	testMap["a"].Value2 = 1

	//
	// Adding entire element in one go
	//
	testMap["b"] = &test{
		Value1: "B",
		Value2: 2,
	}
}

// type definition (NOTE - this is NOT type alias!)
type mapString2Int map[string]int

func demoCreateAndPopulateMap() {
	testMap := make(map[string]*test)
	populateMap(testMap)
	fmt.Println("testMap = ", testMap)

	// len() works on maps
	fmt.Printf("len(testMap) = %d\n", len(testMap))

	// Map type is a reference type so if map var is only declared
	// its value is nil.
	var testMap2 mapString2Int
	if testMap2 == nil {
		fmt.Println("testMap2 (upon declaration) is nil ")
	}
	testMap2 = make(mapString2Int)
	if testMap2 != nil {
		fmt.Println("testMap2 (upon using make for allocation) is not nil")
	}

	// map can be declared and initialized at the same time by using "map literal"
	testMap3 := mapString2Int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	fmt.Println("testMap3 =", testMap3)
}

// Maps in Go have undefined ordering of keys: key order is intentionally randomized between runs to prevent dependency on any perceived order.
// On each run the order of keys or values will be different!
// Also, this is not a mutable iteration: deleting a key will require restarting the iteration.
func demoIterateThroughMap() {
	testMap1 := mapString2Int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	// to iterate through key-value pairs use for-range
	for key, value := range testMap1 {
		fmt.Println("Key:", key, "Value:", value)
	}

	// to iterate through keys use for-range
	for key := range testMap1 {
		fmt.Println("Key:", key)
	}

	// to iterate through values use for-range
	for _, value := range testMap1 {
		fmt.Println("Value:", value)
	}
}

func demoEmptyMap() {
	m := make(map[string]string)
	m["a"] = "A"
	m["b"] = "B"
	fmt.Println("m = ", m)

	// before Go v1.11
	m = make(map[string]string)
	m["c"] = "C"
	m["d"] = "D"
	fmt.Println("m = ", m)

	// after Go v1.11 (faster)
	// compiler will clear old map but will not allocate new space
	for key := range m {
		delete(m, key)
	}
	fmt.Println("m = ", m)
}

// ShowDemo func
func ShowDemo() {
	fmt.Printf("\n\nmapdemo.ShowDemo()\n\n")
	demoCreateAndPopulateMap()
	demoEmptyMap()
	demoIterateThroughMap()
	fmt.Printf("\n\n~mapdemo.ShowDemo()\n\n")
}
