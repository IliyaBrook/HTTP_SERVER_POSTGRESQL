package main

import "fmt"

func main() {
	result := logDays(1)
	fmt.Println(result)
}

func logDays(dayName int) string {
	switch dayName {
	case 1:
		return "יום ראשון הוא יום מעולה"
	case 2:
		return "יום שני"
	case 3:
		return "יום שלישי"
	case 4:
		return "יום רביעי"
	case 5:
		return "יום חבמישי"
	default:
		return "כנראה יום חופש"
	}
}
