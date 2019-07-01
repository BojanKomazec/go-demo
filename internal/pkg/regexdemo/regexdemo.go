package regexdemo

import (
	"fmt"
	"regexp"
)

// ShowDemo func
func ShowDemo() {
	fmt.Printf("\n\nregexdemo.ShowDemo()\n\n")
	s1 := "\"{\"id\":1, \"name\":\"one\", \"address\": {\"number\": \"7\", \"street\":\"Moor Street\"}}\",\"{\"id\":2, \"name\":\"two\", \"address\": {\"number\": \"223\", \"street\":\"Hills Road\"}\"}"
	regExp := "\"{.*?}\""
	regExpObj, err := regexp.Compile(regExp)
	if err != nil {
		fmt.Println("ERROR occurred:", err)
	}
	fmt.Println(regExpObj.FindString(s1))
	fmt.Println(regExpObj.FindAllString(s1, -1))

	fmt.Printf("\n\n~regexdemo.ShowDemo()\n\n")
}
