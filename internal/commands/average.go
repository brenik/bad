package commands

import "github/brenik/bad/internal/storage"

func CalcAverage() {
	storage.Avg = storage.Sum / len(storage.Lines)
}
