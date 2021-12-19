package main

import "fmt"

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func (d Weekday) String() string {
	switch d {
	case Sunday:
		return "Sunday"
	case Monday:
		return "Monday"
	case Tuesday:
		return "Tuesday"
	case Wednesday:
		return "Wednesday"
	case Thursday:
		return "Thursday"
	case Friday:
		return "Friday"
	case Saturday:
		return "Saturday"
	}

	return "Sunday" //default 0 -> "Sunday"
}

func main() {
	var d Weekday
	fmt.Println(d)
	fmt.Println(Weekday(1))
}
