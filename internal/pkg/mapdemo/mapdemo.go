package mapdemo

import (
	"fmt"
)

type Test struct {
	Value1 string
	Value2 int
}

func populateMap(testMap map[string]*Test) {
	//
	// Adding element members one by one
	//
	testMap["a"] = &Test{} // an empty struct has to be added first
	testMap["a"].Value1 = "A"
	testMap["a"].Value2 = 1

	//
	// Adding entire element in one go
	//
	testMap["b"] = &Test{
		Value1: "B",
		Value2: 2,
	}
}

func demoCreateAndPopulateMap() {
	testMap := make(map[string]*Test)
	populateMap(testMap)
	fmt.Println("testMap = ", testMap)
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
	demoCreateAndPopulateMap()
	demoEmptyMap()
}
