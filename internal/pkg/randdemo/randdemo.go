package randdemo

import (
	"fmt"
	"math/rand"
	"time"
)

// GetRandom func returns random integer in range [0, n)
func GetRandom(n int) int {
	seed := time.Now().UnixNano()
	rand.Seed(seed)
	return rand.Intn(n)
}

// GenerateRandomString generates string containing random characters and of
// length specified via argument.
// ASCII letters (upper- and lowercase), numbers and special characters are in
// range [33, 126].
// To get random in range [33, 126] we need to pass 126-33 = 93 to GetRandom()
// and add 33 to its output. If GetRandom returns
// 0 => 0 + 33 = 33
// 93 => 93 +33 = 126
func GenerateRandomString(stringLength int) string {
	bytes := make([]byte, stringLength)
	for i := range bytes {
		bytes[i] = byte(GetRandom(93))
	}
	return string(bytes)
}

// GenerateRandomStrings generates a list which contains random strings.
// List lenght and length of each string are specified in function arguments.
func GenerateRandomStrings(listLength int, stringLength int) []string {
	list := make([]string, listLength)
	for i := range list {
		list[i] = GenerateRandomString(stringLength)
	}
	return list
}

// on each run of the app, the sequence of random numbers generated here would be the same
// (unless different seed is set; by deafult the same seed, 1, is used)
func demoGenerateRandomsWithDefaultSeed() {
	fmt.Println("demoGenerateRandomsWithDefaultSeed()")
	for range [10]int{} {
		fmt.Println(rand.Intn(100))
	}
	fmt.Println("~demoGenerateRandomsWithDefaultSeed()")
}

// on each run of the app this sequence will be different as we're setting a different seed each time
func demoGenerateRandomsWithCustomSeed() {
	fmt.Println("demoGenerateRandomsWithCustomSeed()")
	seed := time.Now().UnixNano()
	rand.Seed(seed)
	for range [10]int{} {
		fmt.Println(rand.Intn(100))
	}
	fmt.Println("~demoGenerateRandomsWithCustomSeed()")
}

func generateRandomByteArray(length int) ([]byte, error) {
	file := make([]byte, length)
	rand.Read(file)
	return file, nil
}

// ShowDemo func
func ShowDemo() {
	fmt.Println()
	fmt.Println("randdemo.ShowDemo()")
	demoGenerateRandomsWithDefaultSeed()
	demoGenerateRandomsWithCustomSeed()
	fmt.Println("~randdemo.ShowDemo()")
}
