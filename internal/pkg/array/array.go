package array

import (
	"fmt"
)

// DemoDeclaration func
func DemoDeclaration() {
	arr1 := [3]string{"AA", "BB", "CC"}
	fmt.Println("arr1 =", arr1)
	fmt.Println("len(arr1) =", len(arr1))

	// https://programming.guide/go/three-dots-ellipsis.html:
	// In an array literal, the ... notation specifies a length equal to the number of elements in the literal.
	arr2 := [...]string{"DD", "EE", "FF", "GG"}
	fmt.Println("arr2 =", arr2)
	fmt.Println("len(arr2) =", len(arr2))
}
