package main

import (
	"fmt"
	"time"
)

func main() {
	var a int
	fmt.Scanf("%d", &a)
	date := time.Date(2000, 12, 1, 0, 0, 0, 0, time.UTC)
	switch {
	case (a <= 1964 && a >= 1946):
		fmt.Println("Привет, бумер!")
	case (a <= 1980 && a >= 1965):
		fmt.Println("Привет, представитель X!")
	case (a <= 1996 && a >= 1981):
		fmt.Println("Привет, миллениал!")
	case (a <= 2012 && a >= 1997):
		fmt.Println("Привет, зумер!")
	case (a >= 2013):
		fmt.Println("Привет, альфа!")
	default:
		fmt.Printf("ШО ТО ТЫ ЗАПИДЕЛСЯ СО СВОИМИ %d ГОДКАМИ\n", a)
	}
	fmt.Println(date)
}
