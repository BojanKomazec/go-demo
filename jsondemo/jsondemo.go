package jsondemo

import (
	"encoding/json"
	"fmt"
)

// type person struct {
// 	name string
// 	age int
// 	siblings []person
// }

// type address struct {
// 	street string
// 	city string
// 	number int
// }

// type register {
// 	value11 bool
// 	value12 int
// 	value13 string
// 	value14 []
// }

type myStruct1 struct {
	field1 int
}

type myStruct2 struct {
	field1 string
	field2 myStruct1
}

type myStruct3 struct {
	field1 bool
}

// struct fields with JSON tags have to be exported (their names should start with capital letter)
type myStruct4 struct {
	Value1 string    `json:"example"`
	Value2 myStruct1 `json:"example2,omitempty"`
	Value3 myStruct3 `json:"example4,omitempty"`
}

func demo1() {
	json1 := `{
		"example": "",
		"example2": {
			"example3": 0
		}
	}`
	fmt.Println("json1 =", json1)
	res1 := myStruct4{}
	json.Unmarshal([]byte(json1), &res1)
	fmt.Println(res1)
	fmt.Println("Value1 =", res1.Value1)
	fmt.Println("Value2 =", res1.Value2)
	fmt.Println("Value3 =", res1.Value3)

	json2 := `{
		"example": "",
		"example4": {
			"example5": false
		}
	}`
	fmt.Println("json2 =", json2)
	res2 := myStruct4{}
	json.Unmarshal([]byte(json2), &res2)
	fmt.Println(res2)
	fmt.Println("Value1 =", res2.Value1)
	fmt.Println("Value2 =", res2.Value2)
	fmt.Println("Value3 =", res2.Value3)
}

func ShowDemo() {
	demo1()
}
