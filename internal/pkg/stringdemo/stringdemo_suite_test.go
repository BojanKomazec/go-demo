package stringdemo

import (
	"testing"

	"github.com/BojanKomazec/go-demo/internal/pkg/randdemo"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

func TestStringdemo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Stringdemo Suite")
}

func BenchmarkTrimPrefixSuffixDemo1(b *testing.B) {
	input := randdemo.GenerateRandomString(98)
	input = "{" + input + "}"
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		trimPrefixSuffixCharacter(input, "{", "}")
	}
}

func BenchmarkTrimFirsLastRune(b *testing.B) {
	input := randdemo.GenerateRandomString(98)
	input = "{" + input + "}"
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		trimFirsLastRune(input)
	}
}

var _ = Describe("StringDemo", func() {
	DescribeTable("Map()",
		func(in []string, t Transform, expected []string) {
			Expect(Map(in, t)).To(Equal(expected))
		},
		Entry(
			"maps empty array to empty array",
			[]string{},
			func(s string) string { return "modified_" + s },
			[]string{},
		),
		Entry(
			"maps each string in an array to a modified one as per transform function",
			[]string{"a", "bb", "ccc"},
			func(s string) string { return "modified_" + s },
			[]string{"modified_a", "modified_bb", "modified_ccc"},
		),
	)
})
