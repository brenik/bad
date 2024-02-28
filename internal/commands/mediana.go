package commands

import (
	"fmt"
	"github/brenik/bad/internal/storage"
	"strconv"
)

func CalcMediana() {
	count := len(storage.Lines)
	ind := (count / 2)

	if count%2 == 0 {
		m1, err := strconv.Atoi(storage.Lines[ind-1])
		if err != nil {
			fmt.Printf("%d of type %T", m1, m1)
		}
		m2, err := strconv.Atoi(storage.Lines[ind])
		if err != nil {
			fmt.Printf("%d of type %T", m1, m1)
		}
		storage.Med = (m1 + m2) / 2
	} else {
		m, err := strconv.Atoi(storage.Lines[ind])
		if err != nil {
			fmt.Printf("%d of type %T", m, m)
		}
		storage.Med = m
	}
}
