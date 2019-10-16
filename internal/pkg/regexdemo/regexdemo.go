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

func demoExtractVersionFromString() {
	// goal: extract "1.22.333.444" from this string
	s1 := "LeadingText-1.22.333.444-TrailingText"
	fmt.Println("s1 =", s1)

	regex, err := regexp.Compile("\\d+(\\.\\d+)+")

	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println(regex.FindAllString(s1, -1))
		fmt.Println(regex.FindString(s1))
	}
}

func demoExtractAnyCharactersBetweenAndIncludingTwoNumbers() {
	fmt.Printf("\nregexdemo.demoExtractAnyCharactersBetweenAndIncludingTwoNumbers()\n")

	// goal: extract "1.22.333.444" or "1.22" or "12233444" from string
	strings := []string{
		"LeadingText-1-TrailingText",
		"LeadingText-12-TrailingText",
		"LeadingText-1.2-TrailingText",
		"LeadingText-11.2-TrailingText",
		"LeadingText-1.22-TrailingText",
		"LeadingText-11.22-TrailingText",
		"LeadingText-1234-TrailingText",
		"LeadingText-1234-TrailingText-5678", // it should extract only first match - 1234
	}

	regex, err := regexp.Compile("\\d(.*\\d)*")

	if err != nil {
		fmt.Println("error: ", err)
	} else {
		for _, s := range strings {
			fmt.Printf("In string %s found pattern %s\n", s, regex.FindString(s))
		}
	}

	regex2, err := regexp.Compile("\\d([\\.\\d]*\\d)*")

	if err != nil {
		fmt.Println("error: ", err)
	} else {
		for _, s := range strings {
			fmt.Printf("In string %s found pattern %s\n", s, regex2.FindString(s))
		}
	}

	fmt.Printf("\n~regexdemo.demoExtractAnyCharactersBetweenAndIncludingTwoNumbers()\n")
}

func findStringIndexDemo() {
	strs := []string{
		"aa-bbb.cccc-1.0.157-ddddd-1.ext",
	}
	regex := regexp.MustCompile("\\d(.*\\d)*")

	for _, s := range strs {
		loc := regex.FindStringIndex(s)
		if loc == nil {
			fmt.Printf("Failed to extract version number from string %s\n", s)
		} else {
			fmt.Printf("Input string: %s. Matched string: %s\n", s, s[loc[0]:loc[1]])
		}
	}

	// expected was 1.0.157 but we got 1.0.157-ddddd-1
	// we need to limit allowed characters to . and digits:
	regex2 := regexp.MustCompile("\\d([\\.\\d]*\\d)*")
	for _, s := range strs {
		loc := regex2.FindStringIndex(s)
		if loc == nil {
			fmt.Printf("Failed to extract version number from string %s\n", s)
		} else {
			fmt.Printf("Input string: %s. Matched string: %s\n", s, s[loc[0]:loc[1]])
		}
	}
}

// ShowDemo func
func ShowDemo() {
	fmt.Printf("\n\nregexdemo.ShowDemo()\n\n")
	demoExtractVersionFromString()
	demoFindingAllStringsThatMatchRegex()
	demoTestThatStringMatchesRegex()
	demoExtractAnyCharactersBetweenAndIncludingTwoNumbers()
	findStringIndexDemo()
	fmt.Printf("\n\n~regexdemo.ShowDemo()\n\n")
}
