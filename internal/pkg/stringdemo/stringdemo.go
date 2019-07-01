package stringdemo

import (
	"fmt"
	"strings"
)

func breakingLongStringsDemo() {
	// s1 is a "raw string literal" (character sequence between back quotes)
	s1 := `This is
		to show that
		backtick string segments
		that appear on a new line in the code
		will be printed on a new line and
		that new line character\n
		is not parsed`
	fmt.Println("s1 =", s1)

	// raw string literal is the string composed of the uninterpreted (implicitly
	// UTF-8-encoded) characters between the quotes; in particular, backslashes
	// have no special meaning.
	s2 := `This is
to show that
backticks don't parse escape escape characters: \"I am quoted!\"`
	fmt.Println("s2 =", s2)

	// string concatenation
	// escapes are interpreted in normal strings
	// Segments coded in a new line are not printed in a new line.
	s3 := "This is " +
		"to show that " +
		"doublequotes DO parse escape escape characters: \"I am quoted!\" " +
		"and that segments coded in a new line " +
		"will NOT be printed on a new line unless " +
		"a new line character\n" +
		"is hardcoded into a string."
	fmt.Println("s3 =", s3)

	// It is still possible to use string formatters like %s, %d etc...:
	query := fmt.Sprintf(`
SELECT
	attachments, version_numbers[1] as version
FROM
	public.myapp_version
WHERE
	attachments IS NOT NULL AND
	attachments[1] IS NOT NULL AND
	version_name[1] = '%s' AND
	version_theme[1] = '%s'
ORDER BY
	version;`, "1.0.0.0", "trial")
	fmt.Println("query =", query)
}

func stringComparisonDemo() {
	s1 := "foo"
	s2 := "Foo"

	if s1 == s2 {
		fmt.Println("Strings are equal")
	} else {
		fmt.Println("Strings are NOT equal")
	}

	// getting the string length
	fmt.Printf("len(s1) = %d, len(s2) = %d\n", len(s1), len(s2))
}

func trimDemo() {
	s1 := "{\"{\"name\":\"Bojan\", \"age\": 40}\"}"
	s2 := strings.TrimPrefix(s1, "{")
	s3 := strings.TrimSuffix(s2, "}")

	fmt.Println("s3 =", s3)
	// output: s3 = "{"name":"Bojan", "age": 40}"

	s4 := strings.TrimSuffix(strings.TrimPrefix(s1, "{"), "}")
	fmt.Println("s4 =", s4)
	// output: s4 = "{"name":"Bojan", "age": 40}"
}

func splitDemo() {
	fmt.Println("splitDemo()")
	s1 := "abcdef"

	// if string does not contain separator, entire original string is returned
	segments := strings.Split(s1, "-")
	fmt.Println("segments =", segments)
	// output: segments = [abcdef]
	fmt.Println("~splitDemo()")
}

// ShowDemo func
func ShowDemo() {
	breakingLongStringsDemo()
	stringComparisonDemo()
	trimDemo()
	splitDemo()
}