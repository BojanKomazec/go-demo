package timedemo

import "fmt"
import "time"

func unixTimeDemo() {
	t1 := time.Now().Unix()
	fmt.Println("time.Now().Unix() =", t1)

	strTime := "2019-10-30 22:23:28"

	layout := "2006-01-02 15:04:05"
	var tn2 int64
	t2, err := time.Parse(layout, strTime)
	if err != nil {
		fmt.Println(err)
	} else {
		tn2 = t2.Unix()
		fmt.Printf("Date %s in unix time = %d", strTime, tn2)
	}
}

func comparisonDemo() {
	strTime1 := "2019-10-30 22:23:28"

	layout := "2006-01-02 15:04:05"
	var tn1 int64
	t1, err := time.Parse(layout, strTime1)
	if err != nil {
		fmt.Println(err)
	} else {
		tn1 = t1.Unix()
		fmt.Printf("Date %s in unix time = %d\n", strTime1, tn1)
	}

	strTime2 := "2020-02-21 11:12:29"

	var tn2 int64
	t2, err := time.Parse(layout, strTime2)
	if err != nil {
		fmt.Println(err)
	} else {
		tn2 = t2.Unix()
		fmt.Printf("Date %s in unix time = %d\n", strTime2, tn2)
	}

	if t1.Before(t2) {
		fmt.Printf("Date %s (%d) is before %s (%d)\n", strTime1, tn1, strTime2, tn2)
	} else {
		fmt.Printf("Date %s (%d) is after or same as %s (%d)\n", strTime1, tn1, strTime2, tn2)
	}

	strTime3 := "2020-02-21 11:12:28"

	var tn3 int64
	t3, err := time.Parse(layout, strTime3)
	if err != nil {
		fmt.Println(err)
	} else {
		tn3 = t3.Unix()
		fmt.Printf("Date %s in unix time = %d\n", strTime3, tn3)
	}

	if t2.After(t3) {
		fmt.Printf("Date %s (%d) is after %s (%d)\n", strTime2, tn2, strTime3, tn3)
	} else {
		fmt.Printf("Date %s (%d) is before or same as %s (%d)\n", strTime2, tn2, strTime3, tn3)
	}

	// We can directly compare Unix times (in64) as that's total count of seconds form 01/01/1970
	if tn2 > tn3 {
		fmt.Printf("Date %s (%d) is after %s (%d)\n", strTime2, tn2, strTime3, tn3)
	} else if tn2 < tn3 {
		fmt.Printf("Date %s (%d) is before %s (%d)\n", strTime2, tn2, strTime3, tn3)
	} else {
		fmt.Printf("Date %s (%d) is same timestamp as %s (%d)\n", strTime2, tn2, strTime3, tn3)
	}

	if t2.After(t2) {
		fmt.Printf("Date %s (%d) is after %s (%d)\n", strTime2, tn2, strTime2, tn2)
	} else {
		fmt.Printf("Date %s (%d) is before or same as %s (%d)\n", strTime2, tn2, strTime2, tn2)
	}
}

// ShowDemo func
func ShowDemo() {
	fmt.Printf("\n\ntimedemo.ShowDemo()\n\n")
	unixTimeDemo()
	comparisonDemo()
	fmt.Printf("\n\n~timedemo.ShowDemo()\n\n")
}
