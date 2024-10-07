package main

import "fmt"

type MyMap map[string]int

func main() {
	var catalog MyMap
	var sum = 0
	order := []string{"хлеб", "буженина", "сыр", "огурцы"}
	catalog = make(MyMap, 12)
	catalog["хлеб"] = 50
	catalog["молоко"] = 100
	catalog["масло"] = 200
	catalog["соль"] = 20
	catalog["огурцы"] = 200
	catalog["сыр"] = 600
	catalog["ветчина"] = 700
	catalog["буженина"] = 900
	catalog["помидоры"] = 250
	catalog["рыба"] = 300
	catalog["хамон"] = 1500
	fmt.Println(catalog)
	fmt.Println("Деликатесы:")
	for k, v := range catalog {
		if v > 500 {
			fmt.Println(k)
		}
	}
	for i := 0; i < len(order); i++ {
		sum += catalog[order[i]]
	}
	fmt.Print(sum)
}
