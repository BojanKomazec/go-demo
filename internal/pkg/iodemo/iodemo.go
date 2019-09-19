package iodemo

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func ioutilReadAllDemo() {
	file, err := os.Open("./internal/pkg/iodemo/test_data/txtfile.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	fmt.Print(b)
}

// ShowDemo func
func ShowDemo() {
	fmt.Printf("\n\niodemo.ShowDemo()\n\n")
	ioutilReadAllDemo()
	fmt.Printf("\n\n~iodemo.ShowDemo()\n\n")
}
