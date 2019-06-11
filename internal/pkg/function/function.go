package function

import (
	"fmt"
)

// VariadicFunction - variadic function example
func VariadicFunction(args ...interface{}) {
	for index, value := range args {
		fmt.Println("index =", index, "; value =", value)
	}
}
