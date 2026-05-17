package main

import (
	"fmt"
	"sort"
)

type People struct {
	names string
	age   int
}

func main() {
	info := []People{
		{"Playmaker", 20},
		{"Joshua", 18},
		{"Bryan", 22},
		{"Chris", 45},
	}
	fmt.Println("sorting age by ascending order")
	//sorting age by ascending order
	sort.Slice(info, func(i, j int) bool {
		return info[i].age < info[j].age
	})
	fmt.Println(info)
	fmt.Println()

	fmt.Println("sorting age by descending order")
	//sorting  age by descending order
	sort.Slice(info, func(i, j int) bool {
		return info[i].age > info[j].age
	})
	fmt.Println()

	fmt.Println("sorting names by ascending order")
	//sorting  names in ascending order
	sort.Slice(info, func(i, j int) bool {
		return info[i].names < info[j].names
	})
	fmt.Println(info)
	fmt.Println()

	fmt.Println("sorting name by descending order")
	sort.Slice(info, func(i, j int) bool {
		return info[i].names > info[j].names
	})

	fmt.Println(info)
}
