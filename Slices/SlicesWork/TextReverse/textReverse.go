package main

import (
	"fmt"
)

func main() {
	str := "Hello, World!"
	runes := []rune(str) // Преобразуем строку в слайс рун
	for i := 0; i < len(runes)/2; i++ {
		runes[i], runes[len(runes)-1-i] = runes[len(runes)-1-i], runes[i]
	}
	reversedStr := string(runes) // Преобразуем обратно в строку
	fmt.Println(reversedStr)
}
