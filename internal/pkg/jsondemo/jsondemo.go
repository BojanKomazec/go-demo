package jsondemo

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Struct fields have to be exported (start with capital letter) so json
// package can access them.
// Case of characters in struct field name does not need to match the case
// in JSON.
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

func demoStructToJSON() {
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

// ShowDemo func
func ShowDemo() {
	fmt.Printf("\n\njsondemo.ShowDemo()\n\n")
	demoStructToJSON()
	demo1()
	unmarshalJSONWithSingleObjectDemo()
	fmt.Printf("\n\n~jsondemo.ShowDemo()\n")
}
