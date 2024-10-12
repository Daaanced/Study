package main

import (
	"encoding/json"
	"fmt"
	"time"
)

//import "fmt"

type Person struct {
	Name        string    `json:"Имя"`
	Email       string    `json:"Почта"`
	DateOfBirth time.Time `json:"-"`
}

func main() {
	he := Person{Name: "Алекс", Email: "alex@yandex.ru"}
	hej, _ := json.MarshalIndent(he, "", "  ")
	fmt.Println(string(hej))
}
