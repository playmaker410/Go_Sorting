package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func Flexiblesort[T any](items []T, less func(a, b T) bool) {
	sort.Slice(items, func(i, j int) bool {
		return less(items[i], items[j])
	})
}

func main() {
	// Ascending
	No := []int{5, 2, 9, 1}
	Flexiblesort(No, func(i, j int) bool {
		return i < j
	})
	fmt.Println(No)

	// Descending

	Flexiblesort(No, func(i, j int) bool {
		return i > j
	})
	fmt.Println(No)
	//custom
	People := []Person{
		{"Adaku", 20},
		{"Helem", 20},
		{"Chris", 36},
		{"Godwin", 29},
	}

	Flexiblesort(People, func(i, j Person) bool {
		if i.Age != j.Age { // if ages are differents
			return i.Age < j.Age // put the smaller one first
		}
		return i.Name < j.Name // if age are same age sort Alphabetically by name.
	})
	fmt.Println(People)

}
