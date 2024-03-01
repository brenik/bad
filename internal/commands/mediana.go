package commands

import (
	"github/brenik/bad/internal/storage"
	"sort"
)

func CalcMediana() {
	sort.Ints(storage.IntLines)

	ind := (storage.Count / 2)

	if storage.Count%2 == 0 {
		m1 := storage.IntLines[ind-1]

		m2 := storage.IntLines[ind]

		storage.Med = (m1 + m2) / 2
	} else {
		storage.Med = storage.IntLines[ind]
	}
}
