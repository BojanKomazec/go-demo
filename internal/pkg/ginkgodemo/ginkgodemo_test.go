package ginkgodemo

import (
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
})
