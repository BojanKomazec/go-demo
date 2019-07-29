package osdemo

import (
	"errors"
	"path/filepath"
	"testing"

	"github.com/BojanKomazec/go-demo/internal/pkg/randdemo"
)

// Benchmark writing to a file a random string of 1024 characters.
// cwd is the directory of the package under test. In this case it is:
// /home/bojan/dev/go/src/github.com/BojanKomazec/go-demo/internal/pkg/osdemo
func BenchmarkWriteToFile1024(b *testing.B) {
	randomStr := randdemo.GenerateRandomString(1024)
	filePath, err := filepath.Abs("../../../data-vol/demo/os/dir1/benchmark.txt")
	if err != nil {
		panic(err)
	}
	dirPath := filepath.Dir(filePath)
	if err := CreateDirIfNotExist(dirPath); err != nil {
		panic(err)
	}
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		for j := 0; j < 100000; j++ {
			if l, err := WriteToFile(filePath, randomStr); err != nil {
				panic(err)
			} else {
				if l != 1024 {
					panic(errors.New("Number of characters written is not 1024"))
				}
			}
		}
	}
}

func BenchmarkWriteToFileBuffered1024(b *testing.B) {
	randomStr := randdemo.GenerateRandomString(1024)
	filePath, err := filepath.Abs("../../../data-vol/demo/os/dir1/benchmark_buffered.txt")
	if err != nil {
		panic(err)
	}
	dirPath := filepath.Dir(filePath)
	if err := CreateDirIfNotExist(dirPath); err != nil {
		panic(err)
	}
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		for j := 0; j < 100000; j++ {
			if l, err := WriteToFileAtPathBuffered(filePath, randomStr); err != nil {
				panic(err)
			} else {
				if l != 1024 {
					panic(errors.New("Number of characters written is not 1024"))
				}
			}
		}
	}
}
