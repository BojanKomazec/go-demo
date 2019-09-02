package pathdemo

import (
	"fmt"
	"log"
	"path/filepath"
)

func filepathBaseDemo() {
	path := "./some/path/some_file.txt"
	base := filepath.Base(path) // some_file.txt
	fmt.Println("base =", base)
}

// ShowDemo func
func ShowDemo() {
	log.Printf("\n\npathdemo.ShowDemo()\n\n")
	filepathBaseDemo()
	log.Printf("\n\n~pathdemo.ShowDemo()\n\n")
}
