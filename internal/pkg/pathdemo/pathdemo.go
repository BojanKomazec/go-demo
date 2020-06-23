package pathdemo

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"
)

func filepathBaseDemo() {
	path := "./some/path/some_file.txt"
	base := filepath.Base(path) // some_file.txt
	fmt.Println("base =", base)
}

// ExtractFileNameWithoutExtension function
func ExtractFileNameWithoutExtension(path string) string {
	fileName := filepath.Base(path)
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

// ShowDemo func
func ShowDemo() {
	log.Printf("\n\npathdemo.ShowDemo()\n\n")
	filepathBaseDemo()
	log.Printf("\n\n~pathdemo.ShowDemo()\n\n")
}
