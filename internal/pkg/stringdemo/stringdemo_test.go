package stringdemo

import (
	"testing"

	"github.com/BojanKomazec/go-demo/internal/pkg/randdemo"
)

func BenchmarkTrimPrefixSuffixDemo1(b *testing.B) {
	input := randdemo.GenerateRandomString(98)
	input = "{" + input + "}"
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		trimPrefixSuffixDemo1(input)
	}
}

func BenchmarkTrimPrefixSuffixDemo2(b *testing.B) {
	input := randdemo.GenerateRandomString(98)
	input = "{" + input + "}"
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		trimPrefixSuffixDemo2(input)
	}
}
