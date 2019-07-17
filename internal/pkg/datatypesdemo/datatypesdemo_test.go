package datatypesdemo

import (
	"testing"

	"github.com/BojanKomazec/go-demo/internal/pkg/randdemo"
)

// BenchmarkCopyList100x16-4
//    1000000 - number of invokations of function under test (FUT) (benchmark tool decides on this)
//    1478 ns/op - average running time of the FUT
//    4080 B/op
//    8 allocs/op
func BenchmarkCopyList1_100x16(b *testing.B) {
	input := randdemo.GenerateRandomStrings(100, 16)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		copyList1(input)
	}
}

func BenchmarkCopyList2_100x16(b *testing.B) {
	input := randdemo.GenerateRandomStrings(100, 16)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		copyList2(input)
	}
}
