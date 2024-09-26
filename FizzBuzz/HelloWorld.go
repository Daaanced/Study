package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Print("FizzBuzz\n")
			continue
		}
		if i%3 == 0 {
			fmt.Print("Fizz\n")
			continue
		}
		if i%5 == 0 {
			fmt.Print("Buzz\n")
			continue
		}
		fmt.Printf("%d\n", i)
	}
}
