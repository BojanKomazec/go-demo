package ginkgodemo

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

func TestConfig(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ginkgo demo Suite")
}

type Colour string

const (
	colourRed   Colour = "red"
	colourGreen Colour = "green"
)

func foo(colour Colour) Colour {
	return colour
}

func add(n1, n2 int) int {
	return n1 + n2
}

func return3() int {
	fmt.Println("return3()")
	return 3
}

func return4() int {
	fmt.Println("return4()")
	return 4
}

var _ = Describe("In package demo", func() {
	DescribeTable("function foo() returns the input argument",
		func(colour Colour) {
			Expect(foo(colour)).To(Equal(colour))
		},
		Entry(
			"colour Red",
			colourRed,
		),
		Entry(
			"colour Green",
			colourGreen,
		),
	)

	Describe("Describe #1", func() {
		BeforeEach(func() {
			fmt.Println("BeforeEach #1")
		})

		Describe("Describe #2", func() {

			BeforeEach(func() {
				fmt.Println("BeforeEach #2")
			})

			When("When #1", func() {
				It("It #1", func() {
					Expect(add(1, 3)).To(BeEquivalentTo(return4()))
					// This is the order of execution:
					// 		In package demo Describe #1 Describe #2 when When #1
					// 		It #1
					// 		BeforeEach #1
					// 		BeforeEach #2
					// 		return4()
				})
			})

			When("When #2", func() {
				DescribeTable("function add() returns the input argument",
					func(n1, n2, expectedRes int) {
						fmt.Println("Executing DescribeTable function...")
						Expect(add(n1, n2)).To(Equal(expectedRes))
					},
					Entry(
						"1 + 2 = 3",
						1,
						2,
						return3(), // this will be executed BEFORE running test suite!
					),
					Entry(
						"4 + 5 = 9",
						4,
						5,
						9,
					),
				)
			})
		})

	})
})
