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

//go:generate stringer -type=Weekday
func main() {
	var d Weekday
	fmt.Println(d)
	fmt.Println(Weekday(1))
}
