package main

import "fmt"

func RemoveDuplicates(input []string) []string {
	result := []string{}
	a := make(map[string]bool)
	for _, k := range input {
		if _, ok := a[k]; !ok {
			a[k] = true
			result = append(result, k)
		}
	}

	fmt.Println(a)

	return result
}

func main() {
	input := []string{
		"cat",
		"dog",
		"bird",
		"dog",
		"parrot",
		"cat",
	}
	output := RemoveDuplicates(input)
	for i, v := range input {
		fmt.Println(i, v)
	}
	fmt.Println(output)
}
