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
