package sortdemo

import "fmt"

type record struct {
	propertyA []string
	propertyB []string
}

// Task: sort slice of records in such way that:
// propertyA length is ascending
// those records with propertyA lenght 0 are the last
// propertyB length is ascending
// those records with propertyB lenght 0 are the last

// {"1001", "1002"}
// {"AB", "CD"}

// {"1001", "1002", "1003"},
// {"AB", "EF"},

// {"1001", "1002"},
// {"AB", "EF"},

// {"1001", "1002"},
// {"CD"},

// {"1001"},
// {"CD"},

// {"1001"},
// {"CD"},

// {"1001", "1002"},
// {"CD"},

// {"1001", "1002"}
// {"AB", "CD"}

// {"1001", "1002"},
// {"AB", "EF"},

// {"1001", "1002", "1003"},
// {"AB", "EF"},

func test_sorting() {
	records := []record{
		record{
			[]string{"1001", "1002"},
			[]string{"AB", "CD"},
		},
		record{
			[]string{"1001", "1002", "1003"},
			[]string{"AB", "EF"},
		},
		record{
			[]string{"1001", "1002"},
			[]string{"AB", "EF"},
		},
		record{
			[]string{"1001", "1002"},
			[]string{"CD"},
		},
		record{
			[]string{"1001"},
			[]string{"CD"},
		},
	}

	fmt.Println("records =", records)
}
