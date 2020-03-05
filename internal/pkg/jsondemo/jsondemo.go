// Use https://jsonlint.com/ for verifying whether JSON is valid or not.
// Use https://nosmileface.dev/jsondiff/ to compare two JSONs.

package jsondemo

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/nsf/jsondiff"
	"github.com/yudai/gojsondiff"
	// import "github.com/yazgazan/jaydiff" is a program, not an importable package
	// "github.com/yazgazan/jaydiff"
)

// Struct fields have to be exported (start with capital letter) so json
// package can access them.
// Case of characters in struct field name does not need to match the case
// in JSON.
// If we want to marshal this struct to JSON where property names start with small letter,
// we'd need to use `json:"age"` etc...
type person struct {
	Age  int
	Name string

	// go-lint: don't use underscores in Go names; struct field Home_Address should be HomeAddress
	// Home_Address string
	// If we name struct field as "HomeAddress" but JSON key is named "home_address", this field
	// will remain empty after unmarshalling unless we explicitly declare the name of the matching
	// key in JSON via `json:"key_name">`
	HomeAddress string `json:"home_address"`
}

// personRegistry is JSON array of person elements
type personRegistry struct {
	collection []person
}

type attachment struct {
	ID          int
	URL         string
	Slug        string
	Title       string
	Description string
	Caption     string
	Parent      int
	MimeType    string `json:"mime_type"`
}

type address struct {
	street string
	city   string
	number int
}

type myMixedTypesStruct struct {
	fieldBool   bool
	fieldInt    int
	fieldString string
	fieldArray  []int
}

// To serialize myMixedTypesStruct struct not into an object but an array, we need to define an array data structure.
// We are using empty interface as all types in the array are different.
type mixedTypesArray [4]interface{}

type rootStruct struct {
	MixedTypesArray mixedTypesArray
}

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

var jsonPersons = `[
	{
		"age": 28,
		"name": "Ada",
		"home_address": "London"
	},
	{
		"age": 32,
		"name": "Nevile",
		"home_address": "Colchester"
	}
]`

var jsonPersons2 = `[
	{
		"age": 28,
		"name": "Ada",
		"home_address": "London"
	}
]`

func demoStructToJSON() {
	fmt.Printf("\njsondemo.demoStructToJSON()\n")
	person := &person{Age: 40, Name: "Bojan", HomeAddress: "Kent, UK"}
	b, err := json.Marshal(person)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
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

func unmarshalJSONWithSingleObjectDemo() {
	jsonStr := "{\"{\"age\":40,\"name\":\"Bojan\", \"home_address\":\"London, UK\"}\"}"
	jsonStr = strings.TrimSuffix(strings.TrimPrefix(jsonStr, "{\""), "\"}")
	fmt.Println("jsonStr: ", jsonStr)
	// output: jsonStr:  {"age":40,"name":"Bojan", "home_address":"London, UK"}

	var person person
	err := json.Unmarshal([]byte(jsonStr), &person)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Printf("person = %+v\n", person)
	// output: person = {Age:40 Name:Bojan HomeAddress:London, UK}

	jsonAttachment := "{\"id\":6715,\"url\":\"https://dev.avastbrowser.com/wp-content/uploads/2019/06/CCleanerBrowserInstaller-75.0.71.41.exe\",\"slug\":\"ccleanerbrowserinstaller-75-0-71-41\",\"title\":\"CCleanerBrowserInstaller-75.0.71.41\",\"description\":\"\",\"caption\":\"\",\"parent\":6714,\"mime_type\":\"application/x-msdownload\"}"

	var attachment attachment
	err = json.Unmarshal([]byte(jsonAttachment), &attachment)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Printf("attachment = %+v\n", attachment)
}

func unmarshalJSONWithRootArrayDemo() {
	// json := ""{\"id\":6643,\"url\":\"https://dev.avastbrowser.com/wp-content/uploads/2019/06/AVGBrowserInstaller-75.0.817.82.exe\",\"slug\":\"avgbrowserinstaller-75-0-817-82\",\"title\":\"AVGBrowserInstaller-75.0.817.82\",\"description\":\"\",\"caption\":\"\",\"parent\":6642,\"mime_type\":\"application/x-msdownload\"}","{\"id\":6644,\"url\":\"https://dev.avastbrowser.com/wp-content/uploads/2019/06/AVGBrowserInstallerIncremental-74.0.773.110-75.0.817.82.exe\",\"slug\":\"avgbrowserinstallerincremental-74-0-773-110-75-0-817-82\",\"title\":\"AVGBrowserInstallerIncremental-74.0.773.110-75.0.817.82\",\"description\":\"\",\"caption\":\"\",\"parent\":6642,\"mime_type\":\"application/x-msdownload\"}","{\"id\":6645,\"url\":\"https://dev.avastbrowser.com/wp-content/uploads/2019/06/AVGBrowserInstallerIncremental-74.0.783.133-75.0.817.82.exe\",\"slug\":\"avgbrowserinstallerincremental-74-0-783-133-75-0-817-82\",\"title\":\"AVGBrowserInstallerIncremental-74.0.783.133-75.0.817.82\",\"description\":\"\",\"caption\":\"\",\"parent\":6642,\"mime_type\":\"application/x-msdownload\"}","{\"id\":6646,\"url\":\"https://dev.avastbrowser.com/wp-content/uploads/2019/06/AVGBrowserInstallerIncremental-74.0.791.133-75.0.817.82.exe\",\"slug\":\"avgbrowserinstallerincremental-74-0-791-133-75-0-817-82\",\"title\":\"AVGBrowserInstallerIncremental-74.0.791.133-75.0.817.82\",\"description\":\"\",\"caption\":\"\",\"parent\":6642,\"mime_type\":\"application/x-msdownload\"}""
}

func jsonArrayDemo() {
	jsonPersons := `
[
	{
		"age": 28,
		"name": "Ada",
		"home_address": "London"
	},
	{
		"age": 32,
		"name": "Nevile",
		"home_address": "Colchester"
	}
]`
	var r personRegistry
	err := json.Unmarshal([]byte(jsonPersons), &r)
	if err != nil {
		fmt.Println("err =", err)
	}
}

func jsonDiffDemo() {
	// opts := jsondiff.DefaultConsoleOptions()
	opts := jsondiff.Options{
		Added: jsondiff.Tag{Begin: "\033[0;32m", End: "\033[0m"},
		// Prefix:  "prefix",
		Removed: jsondiff.Tag{Begin: "\033[0;31m", End: "\033[0m"},
		// Removed: jsondiff.Tag{Begin: "", End: ""},
		Changed: jsondiff.Tag{Begin: "\033[0;33m", End: "\033[0m"},
		Indent:  "    ",
	}
	diff, desc := jsondiff.Compare([]byte(jsonPersons), []byte(jsonPersons2), &opts)
	fmt.Printf("diff = %s\ndesc = %s\n", diff.String(), desc)
}

// CLI version also gives an error (@todo: report it to the package author):
//    $ jd ./internal/pkg/jsondemo/test_data/left.json  ./internal/pkg/jsondemo/test_data/right.json
//    Failed to unmarshal file: json: cannot unmarshal array into Go value of type map[string]interface {}
func gojsondiffDemo() {
	differ := gojsondiff.New()
	diff, err := differ.Compare([]byte(jsonPersons), []byte(jsonPersons2))
	if err != nil {
		fmt.Println("err =", err)
		return
	}
	fmt.Println("diff.Deltas() =", diff.Deltas())
}

func jaydiffDemo() {
}

// https://stackoverflow.com/questions/28015753/serialize-a-mixed-type-json-array-in-go
func serializeArrayOfMixedTypes() {
	fmt.Printf("\njsondemo.serializeArrayOfMixedTypes()\n")
	mts := myMixedTypesStruct{
		true,
		123,
		"text",
		[]int{
			1000,
			1001,
			1002,
		},
	}

	dataToSerialize := rootStruct{
		mixedTypesArray{
			mts.fieldBool,
			mts.fieldInt,
			mts.fieldString,
			mts.fieldArray,
		},
	}

	b, err := json.Marshal(dataToSerialize)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

	fmt.Printf("\n~jsondemo.serializeArrayOfMixedTypes()\n")
}

func serializeMyStruct5() {
	fmt.Printf("\njsondemo.serializeMyStruct5()\n")

	// we want to represent this in JSON as an array e.g.
	// "filter" : ""
	// "filter" : "test"
	// "filter" : [
	//		"test1",
	//      "test2",
	// ]
	type filter struct {
		Component1 string
		// if array is empty, we want it to be empty string in JSON
		// if array has a single element, we want it to be a string in JSON (value of that element)
		// if array has multiple elements, we want it to be an array of strings
		Component2 []string
	}

	type myStruct5 struct {
		ID  string
		Fil filter
	}

	s := myStruct5{
		ID: "1",
		Fil: filter{
			"comp1str",
			[]string{
				"comp2str1",
				"comp2str2",
			},
		},
	}

	fmt.Println("Serializing #1:")
	b, err := json.MarshalIndent(s, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}

// we want to represent this in JSON as an array e.g.
// "filter" : [
//		"test",
//		""
// ]
//
// "filter" : [
//		"test",
//		"test1"
// ]
//
// "filter" : [
//		"com1str",
//		[
//          "com2str1",
//          "com2str2",
//      ]
// ]
type filter struct {
	Component1 string
	// if array is empty, we want it to be empty string in JSON
	// if array has a single element, we want it to be a string in JSON (value of that element)
	// if array has multiple elements, we want it to be an array of strings
	Component2 []string
}

func (f filter) MarshalJSON() ([]byte, error) {
	fmt.Println("MarshalJSON()")

	var v interface{}
	if f.Component2 == nil {
		v = []string{
			f.Component1,
			"",
		}
	} else if len(f.Component2) == 1 {
		v = []string{
			f.Component1,
			f.Component2[0],
		}
	} else {
		v = []interface{}{
			f.Component1,
			f.Component2,
		}
	}

	return json.Marshal(v)
}

type myStruct5 struct {
	ID  string
	Fil filter
}

func serializeMyStruct6() {
	fmt.Printf("\njsondemo.serializeMyStruct6()\n")

	s1 := myStruct5{
		ID: "1",
		Fil: filter{
			"comp1str",
			nil,
		},
	}

	fmt.Println("Serializing #1:")
	b, err := json.MarshalIndent(s1, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

	s2 := myStruct5{
		ID: "1",
		Fil: filter{
			"comp1str",
			[]string{
				"comp2str",
			},
		},
	}

	fmt.Println("Serializing #2:")
	b, err = json.MarshalIndent(s2, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

	s3 := myStruct5{
		ID: "1",
		Fil: filter{
			"comp1str",
			[]string{
				"comp2str1",
				"comp2str2",
			},
		},
	}

	fmt.Println("Serializing #3:")
	b, err = json.MarshalIndent(s3, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}

// ShowDemo func
func ShowDemo() {
	fmt.Printf("\n\njsondemo.ShowDemo()\n\n")
	demoStructToJSON()
	demo1()
	jsonArrayDemo()
	unmarshalJSONWithSingleObjectDemo()
	jsonDiffDemo()
	gojsondiffDemo()
	jaydiffDemo()
	serializeArrayOfMixedTypes()
	serializeMyStruct5()
	serializeMyStruct6()
	fmt.Printf("\n\n~jsondemo.ShowDemo()\n")
}
