package main

import (
	"fmt"
	"sort"
)

func main() {
	slices := []int{14, 11, 3, 20, 5}
	sort.Ints(slices)
	fmt.Println("Sorted interger:", slices)

	names := []string{"Ada", "Ella", "Obi", "Chris"}
	sort.Strings(names)
	fmt.Println("Sorted strings:", names)

	Decimal := []float64{2.45, 0.44, 0.50, 1.5}
	sort.Float64s(Decimal)
	fmt.Println("Sorted float64:", Decimal)
}
