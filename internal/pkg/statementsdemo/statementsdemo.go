package statementsdemo

import "fmt"

import "errors"

func breakOuterLoop1() {
	fmt.Printf("\n\nruntimedemo.breakOuterLoop1()\n\n")
	for i := range make([]int, 10) {
		fmt.Println(i)
		switch i {
		case 5:
			break
		}
	}
}

func breakOuterLoop2() {
	fmt.Printf("\n\nruntimedemo.breakOuterLoop2()\n\n")
	var err error
	for i := range make([]int, 10) {
		fmt.Println(i)
		switch i {
		case 5:
			err = errors.New("I don't like 5s!")
			break
		}
		if err != nil {
			fmt.Println("Error occurred: ", err)
			break
		}
	}
}

func breakOuterLoop3() {
	fmt.Printf("\n\nruntimedemo.breakOuterLoop3()\n\n")
loop:
	for i := range make([]int, 10) {
		fmt.Println(i)
		switch i {
		case 5:
			fmt.Println("I don't like 5s!")
			break loop
		}
	}
}

// ShowDemo func
func ShowDemo() {
	fmt.Printf("\n\nstatementsdemo.ShowDemo()\n\n")
	breakOuterLoop1()
	breakOuterLoop2()
	breakOuterLoop3()
	fmt.Printf("\n\n~statementsdemo.ShowDemo()\n\n")
}
