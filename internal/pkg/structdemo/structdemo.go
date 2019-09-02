package structdemo

import (
	"fmt"

	"github.com/BojanKomazec/go-demo/internal/pkg/structdemo/employee"
)

type person struct {
	name string
	age  int
}

type personRegistry struct {
	collection []person
}

func emptyStructDemo() {
	r1 := personRegistry{}
	fmt.Println("r1 =", r1)

	r2 := personRegistry{
		collection: []person{},
	}
	fmt.Println("r2 =", r2)

	r3 := personRegistry{
		collection: make([]person, 0),
	}
	fmt.Println("r3 =", r3)
}

// ShowDemo function
func ShowDemo() {
	fmt.Printf("\n\nstructdemo.ShowDemo()\n\n")
	e, err := employee.New("Bojan", "Komazec", 15)
	if err != nil {
		println(err)
	}

	e.PrintSalary()

	// implicit assignment of unexported field 'name' in employee.Employee literal
	// e2 := employee.Employee{"John", "Smith", 12}

	emptyStructDemo()
	fmt.Printf("\n\n~structdemo.ShowDemo()\n\n")
}
