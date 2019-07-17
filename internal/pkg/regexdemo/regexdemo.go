package regexdemo

import (
	"fmt"
	"regexp"
)

func demoFindingAllStringsThatMatchRegex() {
	fmt.Printf("\nregexdemo.demoFindingAllStringsThatMatchRegex()\n")
	s1 := "\"{\"id\":1, \"name\":\"one\", \"address\": {\"number\": \"7\", \"street\":\"Moor Street\"}}\",\"{\"id\":2, \"name\":\"two\", \"address\": {\"number\": \"223\", \"street\":\"Hills Road\"}\"}"
	regExp := "\"{.*?}\""
	regExpObj, err := regexp.Compile(regExp)
	if err != nil {
		fmt.Println("ERROR occurred:", err)
	}
	fmt.Println(regExpObj.FindString(s1))
	fmt.Println(regExpObj.FindAllString(s1, -1))
	fmt.Printf("\n~regexdemo.demoFindingAllStringsThatMatchRegex()\n")
}

func checkIsMatching(str, regex string) {
	matched, _ := regexp.MatchString(regex, str)
	fmt.Println("matched =", matched)
}

func demoTestThatStringMatchesRegex() {
	fmt.Printf("\nregexdemo.demoTestThatStringMatchesRegex()\n")
	strArr1 := [...]string{
		"MyApp1.2.3.4.exe",
		"MyApp678.34.6.78675.exe",
	}

	regex := "MyApp[0-9].[0-9].[0-9].[0-9].[exe|pkg]"
	for _, str := range strArr1 {
		checkIsMatching(str, regex)
	}

	regex = "MyApp[0-9]+.[0-9]+.[0-9]+.[0-9]+.[exe|pkg]"
	for _, str := range strArr1 {
		checkIsMatching(str, regex)
	}

	fmt.Printf("\n~regexdemo.demoTestThatStringMatchesRegex()\n")
}

// ShowDemo func
func ShowDemo() {
	fmt.Printf("\n\nregexdemo.ShowDemo()\n\n")
	demoFindingAllStringsThatMatchRegex()
	demoTestThatStringMatchesRegex()
	fmt.Printf("\n\n~regexdemo.ShowDemo()\n\n")
}
