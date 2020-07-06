package errorspackagedemo

import (
	// If we try to import standard library "errors" we'll get an error:
	// 		errors redeclared in this block [compiler] or
	// 		errors redeclared as imported package name
	"fmt"
	"log"

	"github.com/pkg/errors"
)

func f1() error {
	return errors.New("Error in f1()")
}

func f2() error {
	return fmt.Errorf("Error in f2()")
}

func fa() error {
	err1 := f1()

	// This shows it is safe simply to replace default "errors" library with "github.com/pkg/errors" package.
	fmt.Printf("(fmt) err1: %+v\n", err1)
	fmt.Printf("(fmt) err1 (string): %s\n", err1.Error())

	log.Printf("(log) err1: %+v\n", err1)
	log.Printf("(log) err1 (string): %s\n", err1.Error())

	// Printf with %+v prints stack trace for errors created via errors.New()

	return errors.Wrap(err1, "Error in fa()")
}

func fb() error {
	err2 := f2()

	// This shows it is safe simply to replace default "errors" library with "github.com/pkg/errors" package.
	fmt.Printf("(fmt) err2: %+v\n", err2)
	fmt.Printf("(fmt) err2 (string): %s\n", err2.Error())

	log.Printf("(log) err2: %+v\n", err2)
	log.Printf("(log) err2 (string): %s\n", err2.Error())

	// Output shows that Printf with %+v DOES NOT print stack trace for errors created via fmt.Errorf():
	//
	// (fmt) err2: Error in f2()
	// (fmt) err2 (string): Error in f2()
	// 2020/06/23 18:52:58 (log) err2: Error in f2()
	// 2020/06/23 18:52:58 (log) err2 (string): Error in f2()
	//
	// Solution: instead of fmt.Errorf() use errors.Errorf()
	//
	// When our code catches error from some package, we don't know how is that error object created (and if
	// that package used "github.com/pkg/errors" package at all). All we can do it catch that error and wrap it
	// in our error created with errors.New() which will then capture at least the line in which our code calls
	// function from the 3rd party package.
	// This concept is explained here: https://dave.cheney.net/2016/06/12/stack-traces-and-the-errors-package

	return errors.Wrap(err2, "Error in fb()")
}

func printError() {
	erra := fa()
	fmt.Printf("erra: %+v\n", erra)

	errb := fb()
	fmt.Printf("errb: %+v\n", errb)
}

// ShowDemo func
func ShowDemo() {
	fmt.Printf("\n\nerrorspackagedemo.ShowDemo()\n\n")
	printError()
	fmt.Printf("\n\n~errorspackagedemo.ShowDemo()\n\n")
}
