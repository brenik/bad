package commands

import (
	"fmt"
	"github/brenik/bad/internal/service"
	"github/brenik/bad/internal/storage"
	"golang.org/x/exp/maps"
	"strconv"
)

func Compare() {
	var incT = make(map[int]int)
	var dicT = make(map[int]int)

	last, err := strconv.Atoi(storage.Lines[0])
	if err != nil {
		fmt.Printf("%d of type %T", last, last)
	}
	for i, line := range storage.Lines {
		n, err := strconv.Atoi(line)
		if err != nil {
			fmt.Printf("%d of type %T", n, n)
		}
		if n > storage.Max {
			storage.Max = n
		}
		if n < storage.Min {
			storage.Min = n
		}
		storage.Sum += n

		if n > last {
			incT[i] = n
		} else {
			if len(incT) > len(storage.Inc) {
				maps.Clear(storage.Inc)
				storage.Inc = service.ShallowCopyMap(incT).(map[int]int)
			}
			maps.Clear(incT)
			incT[i] = n
		}

		if n < last {
			dicT[i] = n
		} else {
			if len(dicT) > len(storage.Dic) {
				maps.Clear(storage.Dic)
				storage.Dic = service.ShallowCopyMap(dicT).(map[int]int)
			}
			maps.Clear(dicT)
			dicT[i] = n
		}

		last = n
	}
	if len(incT) > len(storage.Inc) {
		maps.Clear(storage.Inc)
		storage.Inc = service.ShallowCopyMap(incT).(map[int]int)
	}
	maps.Clear(incT)
	if len(dicT) > len(storage.Dic) {
		maps.Clear(storage.Dic)
		storage.Dic = service.ShallowCopyMap(dicT).(map[int]int)
	}
	maps.Clear(dicT)

}
