package service

import (
	"fmt"
	"github/brenik/bad/internal/storage"
	"sort"
)

func PrintResult() {
	fmt.Println("Max = ", storage.Max)
	fmt.Println("Min = ", storage.Min)
	fmt.Println("Mediana = ", storage.Med)
	fmt.Println("Average = ", storage.Avg)
	printMap("Increase sequence: ", storage.Inc)
	fmt.Println("")
	printMap("Decrease sequence: ", storage.Dic)
	fmt.Println("")
}

func printMap(label string, m map[int]int) {
	fmt.Printf(label)
	keys := make([]int, 0, len(m))

	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	i := 0

	for _, k := range keys {
		if i == 0 {
			fmt.Printf("%d", m[k])
		} else {
			fmt.Printf(", %d", m[k])
		}
		i++
	}

}
