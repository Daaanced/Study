package main

import "fmt"

func main() {
	mainSlice := make([]int, 100)
	for i := 0; i < len(mainSlice); i++ {
		mainSlice[i] = i + 1
	}

	fmt.Println(mainSlice)

	mainSlice = append(mainSlice[:10], mainSlice[len(mainSlice)-10:]...)
	fmt.Println(mainSlice, len(mainSlice), cap(mainSlice))
	for i := 0; i < len(mainSlice)/2; i++ {
		mainSlice[i], mainSlice[len(mainSlice)-1-i] = mainSlice[len(mainSlice)-1-i], mainSlice[i]
	}
	fmt.Println(mainSlice, len(mainSlice), cap(mainSlice))
}
