package commands

import (
	"fmt"
	"github/brenik/bad/internal/storage"
	"strconv"
)

func Compare() {

	var incaT []int
	var dicaT []int
	storage.Count = len(storage.Lines)
	storage.IntLines = make([]int, storage.Count)

	last, err := strconv.Atoi(storage.Lines[0])
	if err != nil {
		fmt.Printf("%d of type %T", last, last)
	}
	for i, line := range storage.Lines {
		n, err := strconv.Atoi(line)
		if err != nil {
			fmt.Printf("%d of type %T", n, n)
		}
		storage.IntLines[i] = n
		if n > storage.Max {
			storage.Max = n
		}
		if n < storage.Min {
			storage.Min = n
		}
		storage.Sum += n

		if n > last {
			incaT = append(incaT, n)
		} else {
			if len(incaT) > len(storage.Inca) {
				storage.Inca = nil
				storage.Inca = incaT
			}
			incaT = nil
			incaT = append(incaT, n)
		}

		if n < last {
			dicaT = append(dicaT, n)
		} else {
			if len(dicaT) > len(storage.Dica) {
				storage.Dica = nil
				storage.Dica = dicaT
			}
			dicaT = nil
			dicaT = append(dicaT, n)
		}

		last = n
	}

	if len(incaT) > len(storage.Inca) {
		storage.Inca = nil
		storage.Inca = incaT
		incaT = nil
	}
	if len(dicaT) > len(storage.Dica) {
		storage.Dica = nil
		storage.Dica = dicaT
		dicaT = nil
	}

}
