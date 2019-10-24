package kingpindemo

import (
	"fmt"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	// Short('v') means that -v is allowed (apart from --verbose)
	verbose = kingpin.Flag("verbose", "Verbose mode.").Short('v').Bool()
	// name=xxxx is allowed
	name = kingpin.Arg("name", "Name of user.").Required().String()
	// --color=red
	color = kingpin.Flag("color", "Color").Enum("red", "green", "blue")
)

// ShowDemo func
// $ go run cmd/main.go -v Bojan
// $ go run cmd/main.go --verbose Bojan
// $ go run cmd/main.go --verbose name=Bojan
// $ go run cmd/main.go --verbose Bojan --color=red
// $ go run cmd/main.go -v Bojan --color=orange => error: enum value must be one of red,green,blue, got 'orange', try --help
func ShowDemo() {
	fmt.Printf("\n\nkingpindemo.ShowDemo()\n\n")
	kingpin.Parse()
	fmt.Printf("%v, %s, %s\n", *verbose, *name, *color)
	fmt.Printf("\n\n~kingpindemo.ShowDemo()\n\n")
}
