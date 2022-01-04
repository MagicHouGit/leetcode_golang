package main

import "fmt"

func main() {

	// day0 := 1
	// month0 := 3
	// year0 := 1975
	// day1 := 31
	// month1 := 8
	// year1 := 2019
	// day2 := 1
	// month2 := 1
	// year2 := 1973
	day3 := 18
	month3 := 7
	year3 := 1999
	day4 := 15
	month4 := 8
	year4 := 1993

	day5 := 29
	month5 := 2
	year5 := 2000

	fmt.Println(dayOfTheWeek(day5, month5, year5))

	// fmt.Println(dayOfTheWeek(day0, month0, year0))
	// fmt.Println(dayOfTheWeek(day1, month1, year1))
	// fmt.Println(dayOfTheWeek(day2, month2, year2))
	fmt.Println(dayOfTheWeek(day3, month3, year3))
	fmt.Println(dayOfTheWeek(day4, month4, year4))

}

var week = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
var monthDays = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30}

func dayOfTheWeek(day int, month int, year int) string {

	total := 0
	total += day
	for _, v := range monthDays[:month-1] {
		total += v
	}
	if month > 2 && (year%4 == 0 && year%100 != 0 || year%400 == 0) {
		total++
	}
	for i := 1971; i < year; i++ {
		if i%4 == 0 && i%100 != 0 || i%400 == 0 {
			total++
		}
		total += 365
	}
	return week[(total+3)%7]
}
