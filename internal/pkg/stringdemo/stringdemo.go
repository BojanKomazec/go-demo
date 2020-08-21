package stringdemo

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Transformer returns modified input string
type Transformer func(s string) string

// Map iterates through string slice, modifies each string and returns a slice with modified strings.
func Map(in []string, t Transformer) []string {
	// out := make([]string, len(in)) // creates a slice which already contains 3 elements (empty strings)
	out := make([]string, 0, len(in)) // creates a slice which contains 0 but has capacity to hold 3 elements (3 strings)
	for _, s := range in {
		out = append(out, t(s))
	}
	return out
}

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

func trimPrefixSuffixCharacter(in, prefixChar, suffixChar string) string {
	return strings.TrimPrefix(strings.TrimSuffix(in, suffixChar), prefixChar)
}

// "{abcdefghijklmnopqrstuvwyz}"
func trimPrefixSuffixDemo1(in string) {
	s := in
	s = trimPrefixSuffixCharacter(s, "}", "{")
	fmt.Println("s =", s)
}

func trimFirsLastRune(in string) string {
	return trimFirstRune(trimLastRune(in))
}

// "{abcdefghijklmnopqrstuvwyz}"
func trimPrefixSuffixDemo2(in string) {
	s := in
	s = trimFirsLastRune(s)
	fmt.Println("s =", s)
}

func getRunesCount(s string) int {
	n := 0
	for range s {
		n++
	}
	return n
}

func getRunesCount2(s string) int {
	// since Go v1.11 (simpler and faster; no copies are being made)
	return len([]rune(s))
}

func runeDemo() {
	s := "中国 are some Chinese characters"
	fmt.Println("len(s) =", len(s)) // 34
	l := 0

	// range on strings iterates over Unicode code points (runes).
	// The first value is the starting byte index of the rune
	// and the second the rune itself.
	for i, r := range s {
		fmt.Printf("rune index = %d, rune = %v\n", i, r)
		l++
	}
	fmt.Println("l =", l) // 30 (中 is one rune)
}

// taken from https://stackoverflow.com/a/48801414/404421
func trimFirstRune(s string) string {
	for i := range s {
		if i > 0 {
			// The value i is the index in s of the second
			// rune.  Slice to remove the first rune.
			return s[i:]
		}
	}
	// There are 0 or 1 runes in the string.
	return ""
}

func trimLastRune(s string) string {
	i := 0
	for i = range s {
	}
	return s[:i]
}

func splitDemo() {
	fmt.Println("splitDemo()")
	s1 := "abcdef"

	// strings.Split() returns a slice.
	// If string does not contain separator, entire original string is returned.
	segments := strings.Split(s1, "-")
	fmt.Println("segments =", segments)
	// output: segments = [abcdef]
	fmt.Println("~splitDemo()")
}

func joiningStringsDemo() {
	arr := [...]string{"This", "is", "now", "a", "long", "string"}
	// strings.Join has a slice as its first argument => we need to convert an array (all arrays in Go are fix-sized!)
	// to a slice (which describes a portion of an array and is variable-sized)
	log.Println(strings.Join(arr[:], " "))

	// we could have avoided the need for convertion array to slice by declaring variable as a slice at the very beginning:
	slice := []string{"This", "is", "now", "a", "long", "string"}
	log.Println(strings.Join(slice, " "))
}

func replaceDemo() {
	s1 := "This-is-some-1-2-3-4-text.3x3"
	fmt.Println("s1 (before removing substring) =", s1)

	// strings.Replace() can be used for removing a substring from a string
	s1 = strings.Replace(s1, "-text", "", -1)
	fmt.Println("s1 (after removing substring) =", s1)
}

func findIndexOfSubstringDemo() {
	strs := []string{
		"abcdefgh",
		"estrsfghijk",
	}

	substr := "def"

	for _, s := range strs {
		i := strings.Index(s, substr)
		if i == -1 {
			fmt.Printf("String %s does not contain substring %s.\n", s, substr)
		} else {
			fmt.Printf("String %s contains substring %s at starting index %d.\n", s, substr, i)
		}
	}
}

func extractSubstringDemo() {
	s := "abcdefghi"
	substrIndex := []int{2, 5}
	substr := s[substrIndex[0]:substrIndex[1]]
	fmt.Printf("Substring of %s at slice indexes [%d, %d] is %s\n", s, substrIndex[0], substrIndex[1], substr)

	// extract last character of string
	fmt.Printf("Last character of %s is %s\n", s, s[len(s)-1:])
}

func trimLaeadingAndTrailingCharactersFromSet() {
	s := "$^abcdefg,.*"
	cutset := ".,*^$"
	out := strings.Trim(s, cutset)
	fmt.Printf("Input string: %s. Cutset: %s. Output string: %s\n", s, cutset, out)
}

func findSubstringDemo() {
	s := "abcde/fgh\\i"
	c1 := "/"

	fmt.Printf("String %s contains %s: %t\n", s, c1, strings.Contains(s, c1))
}

func removeNonAlphanumeric(s string) string {
	fmt.Println("s = ", s)

	for i, r := range s {
		fmt.Printf("rune index = %d, rune = %v\n", i, r)
	}

	// First remove all Unicode characters.
	// HTML/XML format for Unicode code points: &#nnnn; (n is decimal number - from 0 to 9)
	// NOTE: use raw string (`...`) with regexp.MustCompile to avoid having to escape twice
	reg := regexp.MustCompile(`&#\\d{4};`)
	s = reg.ReplaceAllString(s, "")
	fmt.Println(s)

	// Now remove all non-alphanumeric.
	reg = regexp.MustCompile(`[^a-zA-Z0-9\\s]+`)
	s = reg.ReplaceAllString(s, "")
	return s
}

// whitespace = space, tab (\t), new line (line feed, LF, \n), carriage return (CR, \r) or form feed (\f)
func collapseMultipleWhiteSpacesIntoSingle(s string) string {
	// "+" insures greediness - it will capture multiple instances
	space := regexp.MustCompile(`\\s+`)
	return space.ReplaceAllString(s, " ")
}

func removeNonAlphanumericDemo() {

	// string containing Unicode Decimal Code and non-alphanumeric characters
	s1 := "Bojan&#8217;s test string - (1234)?"
	r1 := removeNonAlphanumeric(s1)
	fmt.Println("Result:", r1)
	r1 = collapseMultipleWhiteSpacesIntoSingle(r1)
	fmt.Println("Result:", r1)

	s2 := "Test &#8211; Test #234"
	r2 := removeNonAlphanumeric(s2)
	fmt.Println("Result:", r2)
	r2 = collapseMultipleWhiteSpacesIntoSingle(r2)
	fmt.Println("Result:", r2)
}

func typeConversionsDemo() {

	n1 := 123456
	// int to string
	s1 := strconv.Itoa(n1)

	// golang does not have 'assert' funtion like C++
	// This can be used to emulate it:
	if s1 != "123456" {
		panic(fmt.Sprintf("assertion failed (%s)", s1))
	}

	t := time.Now().Unix()

	// cannot use t (variable of type int64) as int value in argument to strconv.Itoa
	// s2 := strconv.Itoa(t);

	s2 := strconv.FormatInt(t, 10)
	fmt.Println("s2 (converted from int64) =", s2)
}

// strings.EqualFold
func equalFoldDemo() {
	s1 := "American dream"
	s2 := "american Dream"

	// We want to perform case-insensitive string comparison (in which case "American dream" is equal to "american Dream")

	// strings.ToLower creates a copy of the string
	// This approach is caught by linter with a hint of the improvement:
	//  SA6005: should use strings.EqualFold instead (staticcheck)
	if strings.ToLower(s1) == strings.ToLower(s2) {
		log.Printf("Case-insensitive string comparison: %s == %s\n", s1, s2)
	} else {
		log.Printf("Case-insensitive string comparison: %s != %s\n", s1, s2)
	}

	if strings.EqualFold(s1, s2) {
		log.Printf("Case-insensitive string comparison: %s == %s\n", s1, s2)
	} else {
		log.Printf("Case-insensitive string comparison: %s != %s\n", s1, s2)
	}
}

// ShowDemo func
func ShowDemo() {
	log.Printf("\n\nstringdemo.ShowDemo()\n\n")
	// breakingLongStringsDemo()
	equalFoldDemo()
	// extractSubstringDemo()
	// findIndexOfSubstringDemo()
	// findSubstringDemo()
	// joiningStringsDemo()
	// replaceDemo()
	// removeNonAlphanumericDemo()
	// runeDemo()
	// stringComparisonDemo()
	// trimDemo()
	// trimLaeadingAndTrailingCharactersFromSet()
	// typeConversionsDemo()
	// splitDemo()
	log.Printf("\n\n~stringdemo.ShowDemo()\n\n")
}
