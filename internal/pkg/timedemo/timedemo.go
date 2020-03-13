package timedemo

import "fmt"
import "time"

func unixTimeDemo() {

	var defatultTimeValue time.Time
	// 0001-01-01 00:00:00 +0000 UTC
	fmt.Println("defatultTimeValue =", defatultTimeValue)

	timeNow := time.Now()
	// e.g. 2020-03-11 10:03:11.344109696 +0000 GMT m=+0.000992278
	fmt.Println("timeNow =", timeNow)
	// e.g. 2020-03-11 10:03:11.344109696 +0000 UTC
	fmt.Println("timeNow.UTC() =", timeNow.UTC())

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
		fmt.Printf("Date %s in unix time = %d\n", strTime, tn2)
	}

	t3 := time.Date(2016, time.August, 15, 23, 7, 55, 0, time.UTC)
	tn3 := t3.Unix()
	fmt.Printf("Date %s in unix time = %d\n", t3, tn3)

	t4 := time.Time{}
	tn4 := t4.Unix()
	fmt.Printf("Date %s in unix time = %d\n", t4, tn4)
	fmt.Println(t4.String())
	fmt.Println(t4.Format("2006-01-02 15:04:05"))

	fmt.Println("\nUnix 0 time:")

	unixZeroTime := time.Unix(0, 0)
	// 1970-01-01 01:00:00 +0100 BST
	fmt.Println("unixZeroTime = ", unixZeroTime)

	unixZeroTimeUTC := time.Unix(0, 0).UTC()
	// 1970-01-01 00:00:00 +0000 UTC
	fmt.Println("unixZeroTimeUTC = ", unixZeroTimeUTC)
}

// ToTime converts stringified timestamp into time.Time object
func ToTime(timestamp string) (time.Time, error) {
	t, err := time.Parse("2006-01-02 15:04:05", timestamp)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

func toTimeDemo() {
	s := "2020-03-10 15:38:05"
	t, err := ToTime(s)
	if err != nil {
		fmt.Println("err = ", err.Error())
		return
	}
	if t != time.Date(2020, time.Month(3), 10, 15, 38, 5, 0, time.UTC) {
		fmt.Println("Timestamps are not matching")
		return
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
