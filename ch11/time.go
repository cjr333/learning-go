package main

import (
	"fmt"
	"time"
)

func main() {
	t, err := time.Parse("PM 2006-02-01 03:04:05 -0700", "PM 2016-13-03 05:00:00 +0000")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(t.Format("January 2, 2006 at 3:04:05PM MST"))

	t, err = time.Parse("January 2, 2006 at 3:04:05PM MST", "March 13, 2016 at 5:00:00PM +0000")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(t.Format("2006-02-01 03:04:05PM -0700"))
}
