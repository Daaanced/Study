package main

import (
	"fmt"

	"myproject/contacts"
)

const (
	one = (iota * 2) + 1
	three
	five
	seven
	nine
	eleven
)

func main() {
	fmt.Println(one, three, five, seven, nine, eleven)
	contacts.SetSupport("Служба поддержки")
	//fmt.Println(contacts.GetContact())
	fmt.Println("Email:", contacts.Email)
}
