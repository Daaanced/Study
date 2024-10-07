package main

import "fmt"

type MyMap map[int]int

func main() {
	var a [100]int
	var mainMap MyMap
	var cs int
	for i := range a {
		a[i] = i + 1
	}

	mainMap = make(MyMap, 100)
	mainSlice := []int{}
	for k := range a {
		mainMap[k] += k + 1
	}

	fmt.Scanf("%d", &cs)
	//fmt.Println(mainMap, len(mainMap))

	for k, _ := range mainMap {
		var sum = 0
		for c := k; c != 0; {
			sum += c % 10
			c /= 10
		}
		if cs == sum {
			mainSlice = append(mainSlice, k)
		}
	}
	fmt.Println(mainSlice)
}
