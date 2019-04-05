package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(daysBetween(parseDate("2012-01-01"), parseDate("2012-01-07"))) // 6
	fmt.Println(daysBetween(parseDate("2016-01-01"), parseDate("2017-01-01"))) // 366 because leap year
	fmt.Println(daysBetween(parseDate("2017-01-01"), parseDate("2018-01-01"))) // 365
	fmt.Println(daysBetween(parseDate("2016-01-01"), parseDate("2016-01-01"))) // 0
}

func daysBetween(a, b time.Time) int {
	return int(b.Sub(a).Hours() / 24)
}

func parseDate(s string) time.Time {
	d, _ := time.Parse("2006-01-02", s)
	return d
}
