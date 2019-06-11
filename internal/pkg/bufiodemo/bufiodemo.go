package bufiodemo

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadStringLine() (string, error) {
	stdinReader := bufio.NewReader(os.Stdin)

	if inputString, err := stdinReader.ReadString('\n'); err != nil {
		fmt.Println("Invalid input. Error: ", err)
		return "", err
	} else {
		// if exists, remove trailing Line Feed (added on Windows and on Unix/OSX)
		if strings.HasSuffix(inputString, "\n") {
			inputString = strings.TrimSuffix(inputString, "\n")
		}

		// if exist, remove trailing Carriage Return (added on Windows)
		if strings.HasSuffix(inputString, "\r") {
			inputString = strings.TrimSuffix(inputString, "\r")
		}

		inputString = strings.TrimSpace(inputString)

		fmt.Println("inputString = ", inputString)
		return inputString, nil
	}
}

func printElementsInfo(arr []string) {
	fmt.Println("Number of elements:", len(arr))
	if len(arr) > 0 {
		fmt.Println("Elements: ")
		for i, v := range arr {
			fmt.Println("arr[", i, "] =", v, "; len(v)=", len(v))
		}
	} else {
		fmt.Println("No elements found!")
	}
}

// strings.Split(str, " "):
// If it comes across N space characters in a row, it would split them into N-1 chunks containing empty string.
func ToIntegers(str string) ([]int, error) {
	parts := strings.Split(str, " ")
	printElementsInfo(parts)
	slice := make([]int, 0, len(parts))
	for _, v := range parts {
		if len(v) == 0 {
			continue
		}
		if n, err := strconv.Atoi(v); err != nil {
			fmt.Println("Error:", err)
			return nil, err
		} else {
			slice = append(slice, n)
		}
	}
	return slice, nil
}

func ReadIntegersLine() ([]int, error) {
	if inputString, err := ReadStringLine(); err != nil {
		fmt.Println("Error: ", err)
		return nil, err
	} else {
		fmt.Println("inputString = ", inputString)
		return ToIntegers(inputString)
	}
}

func printElements(arr []int) {
	fmt.Println("Number of elements:", len(arr))
	if len(arr) > 0 {
		fmt.Println("Elements: ")
		for _, v := range arr {
			fmt.Printf("%d ", v)
		}
	} else {
		fmt.Println("No elements found!")
	}
}

func ReadIntegersLineDemo() {
	if intSlice, err := ReadIntegersLine(); err != nil {
		fmt.Println("Error:", err)
	} else {
		printElements(intSlice)
	}
}
